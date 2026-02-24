package rl

import "sort"

const inputScriptPressedHoldSeconds = 0.1

type InputScriptEventType int32

const (
	InputScriptEventOnPressed InputScriptEventType = iota
	InputScriptEventOnDown
	InputScriptEventOnReleased
	InputScriptEventOnUp
)

type InputScriptCommand struct {
	Time  float64
	Key   UnifiedKeyType
	Event InputScriptEventType
}

type inputScriptKeyState struct {
	pressed    bool
	down       bool
	released   bool
	holdUntil  float64
	manualDown bool
}

type inputScriptRuntime struct {
	commands []InputScriptCommand
	keys     map[UnifiedKeyType]inputScriptKeyState

	commandIndex int
	startTime    float64
	playhead     float64
	loaded       bool
	playing      bool
}

var scriptedInput = inputScriptRuntime{
	keys: make(map[UnifiedKeyType]inputScriptKeyState),
}

func LoadInputScript(commands []InputScriptCommand) {
	scriptedInput.commands = append(scriptedInput.commands[:0], commands...)
	sort.Slice(scriptedInput.commands, func(i, j int) bool {
		return scriptedInput.commands[i].Time < scriptedInput.commands[j].Time
	})

	scriptedInput.commandIndex = 0
	scriptedInput.startTime = 0
	scriptedInput.playhead = 0
	scriptedInput.loaded = len(scriptedInput.commands) > 0
	scriptedInput.playing = false
	clear(scriptedInput.keys)
}

func PlayInputScript() {
	if !scriptedInput.loaded || scriptedInput.playing {
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
	if !scriptedInput.loaded {
		return
	}

	scriptedInput.commandIndex = 0
	scriptedInput.playhead = 0
	clear(scriptedInput.keys)

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

	for key, state := range scriptedInput.keys {
		state.pressed = false
		state.released = false
		scriptedInput.keys[key] = state
	}

	for scriptedInput.commandIndex < len(scriptedInput.commands) {
		cmd := scriptedInput.commands[scriptedInput.commandIndex]
		if cmd.Time > elapsed {
			break
		}
		scriptedInput.apply(cmd)
		scriptedInput.commandIndex++
	}

	for key, state := range scriptedInput.keys {
		nextDown := state.manualDown || elapsed < state.holdUntil
		if state.down && !nextDown {
			state.released = true
		}
		state.down = nextDown

		if !state.down && !state.pressed && !state.released {
			delete(scriptedInput.keys, key)
			continue
		}

		scriptedInput.keys[key] = state
	}
}

func scriptedInputEnabled() bool {
	return scriptedInput.playing
}

func scriptedInputKeyPressed(key KeyType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedKeyboardKey(key)).pressed
}

func scriptedInputKeyDown(key KeyType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedKeyboardKey(key)).down
}

func scriptedInputKeyReleased(key KeyType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedKeyboardKey(key)).released
}

func scriptedInputKeyUp(key KeyType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return !scriptedInput.query(unifiedKeyboardKey(key)).down
}

func scriptedInputKeyDownCount(realValue int) int {
	if !scriptedInputEnabled() {
		return realValue
	}

	count := 0
	for key, state := range scriptedInput.keys {
		if key.Device() == Keyboard && state.down {
			count++
		}
	}
	return count
}

func scriptedInputMousePressed(button MouseButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedMouseKey(button)).pressed
}

func scriptedInputMouseDown(button MouseButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedMouseKey(button)).down
}

func scriptedInputMouseReleased(button MouseButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedMouseKey(button)).released
}

func scriptedInputMouseUp(button MouseButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return !scriptedInput.query(unifiedMouseKey(button)).down
}

func scriptedInputGamepadPressed(gamepad int, button GamepadButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedGamepadKey(gamepad, button)).pressed
}

func scriptedInputGamepadDown(gamepad int, button GamepadButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedGamepadKey(gamepad, button)).down
}

func scriptedInputGamepadReleased(gamepad int, button GamepadButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return scriptedInput.query(unifiedGamepadKey(gamepad, button)).released
}

func scriptedInputGamepadUp(gamepad int, button GamepadButtonType, realValue bool) bool {
	if !scriptedInputEnabled() {
		return realValue
	}
	return !scriptedInput.query(unifiedGamepadKey(gamepad, button)).down
}

func (s *inputScriptRuntime) elapsed() float64 {
	elapsed := GetTime() - s.startTime
	if elapsed < 0 {
		return 0
	}
	return elapsed
}

func (s *inputScriptRuntime) apply(cmd InputScriptCommand) {
	state := s.keys[cmd.Key]

	switch cmd.Event {
	case InputScriptEventOnPressed:
		state.pressed = true
		holdUntil := cmd.Time + inputScriptPressedHoldSeconds
		if holdUntil > state.holdUntil {
			state.holdUntil = holdUntil
		}
	case InputScriptEventOnDown:
		state.manualDown = true
		if !state.down {
			state.pressed = true
		}
	case InputScriptEventOnReleased:
		if state.down || state.manualDown || state.holdUntil > cmd.Time {
			state.released = true
		}
		state.manualDown = false
		state.holdUntil = cmd.Time
	case InputScriptEventOnUp:
		state.manualDown = false
		state.holdUntil = cmd.Time
	}

	s.keys[cmd.Key] = state
}

func (s *inputScriptRuntime) query(key UnifiedKeyType) inputScriptKeyState {
	state, ok := s.keys[key]
	if !ok {
		return inputScriptKeyState{}
	}
	return state
}

func unifiedKeyboardKey(key KeyType) UnifiedKeyType {
	return UnifiedKeyType(KeyboardMask) | UnifiedKeyType(key)
}

func unifiedMouseKey(button MouseButtonType) UnifiedKeyType {
	return UnifiedKeyType(MouseMask) | UnifiedKeyType(button)
}

func unifiedGamepadKey(gamepad int, button GamepadButtonType) UnifiedKeyType {
	gamepadIndex := gamepad
	if gamepadIndex < 0 {
		gamepadIndex = 0
	}
	maxIndex := (1 << InputGamepadIndexMaskWidth) - 1
	if gamepadIndex > maxIndex {
		gamepadIndex = maxIndex
	}

	indexBits := UnifiedKeyType(gamepadIndex) << (InputDeviceMaskShift - InputGamepadIndexMaskWidth)
	return UnifiedKeyType(GamepadMask) | indexBits | UnifiedKeyType(button)
}
