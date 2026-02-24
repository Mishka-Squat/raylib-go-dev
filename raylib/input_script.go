package rl

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

const inputScriptPressedHoldSeconds = 0.1

type InputScriptEventType int32

const (
	InputScriptEventOnPressed InputScriptEventType = iota
	InputScriptEventOnDown
	InputScriptEventOnReleased
	InputScriptEventOnUp
	InputScriptEventMousePosition
	InputScriptEventMouseWheelMove
)

type InputScriptCommand struct {
	Time  float64
	Key   UnifiedKeyType
	Event InputScriptEventType
	X     float32
	Y     float32
}

type inputScriptPendingRelease struct {
	Time float64
	Key  UnifiedKeyType
}

type inputScriptRuntime struct {
	commands []InputScriptCommand
	pending  []inputScriptPendingRelease

	commandIndex int
	startTime    float64
	playhead     float64
	playing      bool
}

var scriptedInput inputScriptRuntime

func LoadInputScript(commands []InputScriptCommand) {
	scriptedInput.commands = append(scriptedInput.commands[:0], commands...)
	sort.SliceStable(scriptedInput.commands, func(i, j int) bool {
		return scriptedInput.commands[i].Time < scriptedInput.commands[j].Time
	})

	scriptedInput.pending = scriptedInput.pending[:0]
	scriptedInput.commandIndex = 0
	scriptedInput.startTime = 0
	scriptedInput.playhead = 0
	scriptedInput.playing = false
}

func LoadInputScriptFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read input script file: %w", err)
	}

	var commands []InputScriptCommand
	if err = json.Unmarshal(data, &commands); err == nil {
		LoadInputScript(commands)
		return nil
	}

	var payload struct {
		Commands []InputScriptCommand `json:"commands"`
	}
	if err2 := json.Unmarshal(data, &payload); err2 == nil {
		LoadInputScript(payload.Commands)
		return nil
	}

	return fmt.Errorf("parse input script file as []InputScriptCommand or {commands:[...]}: %w", err)
}

func AddInputScriptCommand(command InputScriptCommand) {
	scriptedInput.commands = append(scriptedInput.commands, command)
	sort.SliceStable(scriptedInput.commands, func(i, j int) bool {
		return scriptedInput.commands[i].Time < scriptedInput.commands[j].Time
	})

	if scriptedInput.commandIndex > len(scriptedInput.commands) {
		scriptedInput.commandIndex = len(scriptedInput.commands)
	}
}

func PlayInputScript() {
	if len(scriptedInput.commands) == 0 || scriptedInput.playing {
		return
	}

	scriptedInput.startTime = GetTime() - scriptedInput.playhead
	scriptedInput.playing = true
}

func PauseInputScript() {
	if !scriptedInput.playing {
		return
	}

	scriptedInput.playhead = scriptedInput.elapsed()
	scriptedInput.playing = false
}

func RewindInputScript() {
	if len(scriptedInput.commands) == 0 {
		return
	}

	scriptedInput.commandIndex = 0
	scriptedInput.playhead = 0
	scriptedInput.pending = scriptedInput.pending[:0]

	if scriptedInput.playing {
		scriptedInput.startTime = GetTime()
	}
}

func UpdateInputScript() {
	if !scriptedInput.playing {
		return
	}

	elapsed := scriptedInput.elapsed()
	scriptedInput.playhead = elapsed

	for {
		nextCommandTime := inputScriptInfinity
		if scriptedInput.commandIndex < len(scriptedInput.commands) {
			nextCommandTime = scriptedInput.commands[scriptedInput.commandIndex].Time
		}

		nextReleaseTime := inputScriptInfinity
		if len(scriptedInput.pending) > 0 {
			nextReleaseTime = scriptedInput.pending[0].Time
		}

		nextTime := nextCommandTime
		if nextReleaseTime < nextTime {
			nextTime = nextReleaseTime
		}
		if nextTime > elapsed {
			break
		}

		if nextReleaseTime <= nextCommandTime {
			release := scriptedInput.pending[0]
			scriptedInput.pending = scriptedInput.pending[1:]
			scriptedInput.generateUp(release.Key)
			continue
		}

		command := scriptedInput.commands[scriptedInput.commandIndex]
		scriptedInput.commandIndex++
		scriptedInput.apply(command)
	}
}

const inputScriptInfinity = 1e18

func (s *inputScriptRuntime) elapsed() float64 {
	elapsed := GetTime() - s.startTime
	if elapsed < 0 {
		return 0
	}
	return elapsed
}

func (s *inputScriptRuntime) apply(command InputScriptCommand) {
	switch command.Event {
	case InputScriptEventOnPressed:
		s.generateDown(command.Key)
		s.queueRelease(command.Time+inputScriptPressedHoldSeconds, command.Key)
	case InputScriptEventOnDown:
		s.generateDown(command.Key)
	case InputScriptEventOnReleased, InputScriptEventOnUp:
		s.generateUp(command.Key)
	case InputScriptEventMousePosition:
		DebugGenerateMousePosition(command.X, command.Y)
	case InputScriptEventMouseWheelMove:
		DebugGenerateMouseWheelMove(command.X, command.Y)
	}
}

func (s *inputScriptRuntime) queueRelease(time float64, key UnifiedKeyType) {
	s.pending = append(s.pending, inputScriptPendingRelease{Time: time, Key: key})
	sort.SliceStable(s.pending, func(i, j int) bool {
		return s.pending[i].Time < s.pending[j].Time
	})
}

func (s *inputScriptRuntime) generateDown(key UnifiedKeyType) {
	switch key.Device() {
	case Keyboard:
		DebugGenerateKeyDown(key.Keyboard())
	case Mouse:
		DebugGenerateMouseDown(key.Mouse())
	}
}

func (s *inputScriptRuntime) generateUp(key UnifiedKeyType) {
	switch key.Device() {
	case Keyboard:
		DebugGenerateKeyUp(key.Keyboard())
	case Mouse:
		DebugGenerateMouseUp(key.Mouse())
	}
}
