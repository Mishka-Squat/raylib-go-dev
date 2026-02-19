/*
Package raylib - Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming.

raylib is highly inspired by Borland BGI graphics lib and by XNA framework.
raylib could be useful for prototyping, tools development, graphic applications, embedded systems and education.

NOTE for ADVENTURERS: raylib is a programming library to learn videogames programming; no fancy interface, no visual helpers, no auto-debugging... just coding in the most pure spartan-programmers way.
*/
package rl

import (
	"io"
	"runtime"
	"unsafe"

	"github.com/igadmg/gamemath/rect2"
	"github.com/igadmg/gamemath/vector2"
	"github.com/igadmg/gamemath/vector3"
	"github.com/igadmg/gamemath/vector4"
	"github.com/igadmg/goex/image/colorex"
	"github.com/igadmg/goex/mathex"
	"golang.org/x/exp/constraints"
)

func init() {
	// Make sure the main goroutine is bound to the main thread.
	runtime.LockOSThread()
}

// Wave type, defines audio wave data
type Wave struct {
	// Number of samples
	FrameCount uint32
	// Frequency (samples per second)
	SampleRate uint32
	// Bit depth (bits per sample): 8, 16, 32 (24 not supported)
	SampleSize uint32
	// Number of channels (1-mono, 2-stereo)
	Channels uint32
	// Buffer data pointer
	Data unsafe.Pointer
}

// NewWave - Returns new Wave
func NewWave(sampleCount, sampleRate, sampleSize, channels uint32, data []byte) Wave {
	d := unsafe.Pointer(&data[0])

	return Wave{sampleCount, sampleRate, sampleSize, channels, d}
}

// Checks if wave data is valid (data loaded and parameters)
func (w Wave) IsValid() bool {
	return w.Data != nil && // Validate wave data available
		w.FrameCount > 0 && // Validate frame count
		w.SampleRate > 0 && // Validate sample rate is supported
		w.SampleSize > 0 && // Validate sample size is supported
		w.Channels > 0 // Validate number of channels supported
}

func (w *Wave) Unload() {
	UnloadWave(w)
}

// AudioCallback function.
type AudioCallback func(data []float32, frames int)

// Sound source type
type Sound struct {
	Stream     AudioStream
	FrameCount uint32
	_          [4]byte
}

// Checks if a sound is valid (data loaded and buffers initialized)
func (s Sound) IsValid() bool {
	return s.FrameCount > 0 && // Validate frame count
		s.Stream.IsValid() // Validate stream buffer
}

func (s *Sound) Unload() {
	UnloadSound(s)
}

// Music type (file streaming from memory)
// NOTE: Anything longer than ~10 seconds should be streamed
type Music struct {
	Stream     AudioStream
	FrameCount uint32
	Looping    bool
	CtxType    int32
	CtxData    unsafe.Pointer
}

// Checks if a music stream is valid (context and buffers initialized)
func (m Music) IsValid() bool {
	return m.CtxData != nil && // Validate context loaded
		m.FrameCount > 0 && // Validate audio frame count
		m.Stream.IsValid() // Validate audio stream
}

func (m *Music) Unload() {
	UnloadMusicStream(m)
}

// AudioStream type
// NOTE: Useful to create custom audio streams not bound to a specific file
type AudioStream struct {
	// Buffer
	Buffer *AudioBuffer
	// Processor
	Processor *AudioProcessor
	// Frequency (samples per second)
	SampleRate uint32
	// Bit depth (bits per sample): 8, 16, 32 (24 not supported)
	SampleSize uint32
	// Number of channels (1-mono, 2-stereo)
	Channels uint32
	_        [4]byte
}

// Checks if an audio stream is valid (buffers initialized)
func (s AudioStream) IsValid() bool {
	return s.Buffer != nil && // Validate stream buffer
		s.SampleRate > 0 && // Validate sample rate is supported
		s.SampleSize > 0 && // Validate sample size is supported
		s.Channels > 0 // Validate number of channels supported
}

func (a *AudioStream) Unload() {
	UnloadAudioStream(a)
}

type maDataConverter struct {
	FormatIn                uint32
	FormatOut               uint32
	ChannelsIn              uint32
	ChannelsOut             uint32
	SampleRateIn            uint32
	SampleRateOut           uint32
	DitherMode              uint32
	ExecutionPath           uint32
	ChannelConverter        maChannelConverter
	Resampler               maResampler
	HasPreFormatConversion  uint8
	HasPostFormatConversion uint8
	HasChannelConverter     uint8
	HasResampler            uint8
	IsPassthrough           uint8
	X_ownsHeap              uint8
	X_pHeap                 *byte
}

type maChannelConverter struct {
	Format         uint32
	ChannelsIn     uint32
	ChannelsOut    uint32
	MixingMode     uint32
	ConversionPath uint32
	PChannelMapIn  *uint8
	PChannelMapOut *uint8
	PShuffleTable  *uint8
	Weights        [8]byte
	X_pHeap        *byte
	X_ownsHeap     uint32
	Pad_cgo_0      [4]byte
}

type maResampler struct {
	PBackend         *byte
	PBackendVTable   *maResamplingBackendVtable
	PBackendUserData *byte
	Format           uint32
	Channels         uint32
	SampleRateIn     uint32
	SampleRateOut    uint32
	State            [136]byte
	X_pHeap          *byte
	X_ownsHeap       uint32
	Pad_cgo_0        [4]byte
}

type maResamplingBackendVtable struct {
	OnGetHeapSize                 *[0]byte
	OnInit                        *[0]byte
	OnUninit                      *[0]byte
	OnProcess                     *[0]byte
	OnSetRate                     *[0]byte
	OnGetInputLatency             *[0]byte
	OnGetOutputLatency            *[0]byte
	OnGetRequiredInputFrameCount  *[0]byte
	OnGetExpectedOutputFrameCount *[0]byte
	OnReset                       *[0]byte
}

type AudioBuffer struct {
	Converter            maDataConverter
	Callback             *[0]byte
	Processor            *AudioProcessor
	Volume               float32
	Pitch                float32
	Pan                  float32
	Playing              bool
	Paused               bool
	Looping              bool
	Usage                int32
	IsSubBufferProcessed [2]bool
	SizeInFrames         uint32
	FrameCursorPos       uint32
	FramesProcessed      uint32
	Data                 *uint8
	Next                 *AudioBuffer
	Prev                 *AudioBuffer
}

type AudioProcessor struct {
	Process *[0]byte
	Next    *AudioProcessor
	Prev    *AudioProcessor
}

// AutomationEvent - Automation event
//type AutomationEvent struct {
//	Frame  uint32
//	Type   uint32
//	Params [4]int32
//}

// AutomationEventList - Automation event list
//type AutomationEventList struct {
//	Capacity uint32
//	Count    uint32
//	// Events array (c array)
//	//
//	// Use AutomationEventList.GetEvents instead (go slice)
//	Events *AutomationEvent
//}

//func (a *AutomationEventList) GetEvents() []AutomationEvent {
//	return unsafe.Slice(a.Events, a.Count)
//}

// CameraMode type
type CameraMode int32

// Camera system modes
const (
	CameraCustom CameraMode = iota
	CameraFree
	CameraOrbital
	CameraFirstPerson
	CameraThirdPerson
)

// CameraProjection type
type CameraProjection int32

// Camera projection modes
const (
	CameraPerspective CameraProjection = iota
	CameraOrthographic
)

// Some basic Defines
const (
	Pi      = 3.1415927
	Deg2rad = 0.017453292
	Rad2deg = 57.295776
)

// Raylib Config Flags

type ConfigFlags = uint32

const (
	FlagVsyncHint              ConfigFlags = 0x00000040 // Set to try enabling V-Sync on GPU
	FlagFullscreenMode         ConfigFlags = 0x00000002 // Set to run program in fullscreen
	FlagWindowResizable        ConfigFlags = 0x00000004 // Set to allow resizable window
	FlagWindowUndecorated      ConfigFlags = 0x00000008 // Set to disable window decoration (frame and buttons)
	FlagWindowHidden           ConfigFlags = 0x00000080 // Set to hide window
	FlagWindowMinimized        ConfigFlags = 0x00000200 // Set to minimize window (iconify)
	FlagWindowMaximized        ConfigFlags = 0x00000400 // Set to maximize window (expanded to monitor)
	FlagWindowUnfocused        ConfigFlags = 0x00000800 // Set to window non focused
	FlagWindowTopmost          ConfigFlags = 0x00001000 // Set to window always on top
	FlagWindowAlwaysRun        ConfigFlags = 0x00000100 // Set to allow windows running while minimized
	FlagWindowTransparent      ConfigFlags = 0x00000010 // Set to allow transparent window
	FlagWindowHighdpi          ConfigFlags = 0x00002000 // Set to support HighDPI
	FlagWindowMousePassthrough ConfigFlags = 0x00004000 // Set to support mouse passthrough, only supported when FLAG_WINDOW_UNDECORATED
	FlagBorderlessWindowedMode ConfigFlags = 0x00008000 // Set to run program in borderless windowed mode
	FlagMsaa4xHint             ConfigFlags = 0x00000020 // Set to try enabling MSAA 4X
	FlagInterlacedHint         ConfigFlags = 0x00010000 // Set to try enabling interlaced video format (for V3D)
)

type InputEventType int

const (
	InputEventPressed = iota // Default input event
	InputEventDown
	InputEventReleased
	InputEventUp
)

type KeyType int32

const (
	// KeyNull is used for no key pressed
	KeyNull KeyType = 0

	// Keyboard Function Keys
	KeySpace        KeyType = 32
	KeyEscape       KeyType = 256
	KeyEnter        KeyType = 257
	KeyTab          KeyType = 258
	KeyBackspace    KeyType = 259
	KeyInsert       KeyType = 260
	KeyDelete       KeyType = 261
	KeyRight        KeyType = 262
	KeyLeft         KeyType = 263
	KeyDown         KeyType = 264
	KeyUp           KeyType = 265
	KeyPageUp       KeyType = 266
	KeyPageDown     KeyType = 267
	KeyHome         KeyType = 268
	KeyEnd          KeyType = 269
	KeyCapsLock     KeyType = 280
	KeyScrollLock   KeyType = 281
	KeyNumLock      KeyType = 282
	KeyPrintScreen  KeyType = 283
	KeyPause        KeyType = 284
	KeyF1           KeyType = 290
	KeyF2           KeyType = 291
	KeyF3           KeyType = 292
	KeyF4           KeyType = 293
	KeyF5           KeyType = 294
	KeyF6           KeyType = 295
	KeyF7           KeyType = 296
	KeyF8           KeyType = 297
	KeyF9           KeyType = 298
	KeyF10          KeyType = 299
	KeyF11          KeyType = 300
	KeyF12          KeyType = 301
	KeyLeftShift    KeyType = 340
	KeyLeftControl  KeyType = 341
	KeyLeftAlt      KeyType = 342
	KeyLeftSuper    KeyType = 343
	KeyRightShift   KeyType = 344
	KeyRightControl KeyType = 345
	KeyRightAlt     KeyType = 346
	KeyRightSuper   KeyType = 347
	KeyKbMenu       KeyType = 348
	KeyLeftBracket  KeyType = 91
	KeyBackSlash    KeyType = 92
	KeyRightBracket KeyType = 93
	KeyGrave        KeyType = 96

	// Keyboard Number Pad Keys
	KeyKp0        KeyType = 320
	KeyKp1        KeyType = 321
	KeyKp2        KeyType = 322
	KeyKp3        KeyType = 323
	KeyKp4        KeyType = 324
	KeyKp5        KeyType = 325
	KeyKp6        KeyType = 326
	KeyKp7        KeyType = 327
	KeyKp8        KeyType = 328
	KeyKp9        KeyType = 329
	KeyKpDecimal  KeyType = 330
	KeyKpDivide   KeyType = 331
	KeyKpMultiply KeyType = 332
	KeyKpSubtract KeyType = 333
	KeyKpAdd      KeyType = 334
	KeyKpEnter    KeyType = 335
	KeyKpEqual    KeyType = 336

	// Keyboard Alpha Numeric Keys
	KeyApostrophe KeyType = 39
	KeyComma      KeyType = 44
	KeyMinus      KeyType = 45
	KeyPeriod     KeyType = 46
	KeySlash      KeyType = 47
	KeyZero       KeyType = 48
	KeyOne        KeyType = 49
	KeyTwo        KeyType = 50
	KeyThree      KeyType = 51
	KeyFour       KeyType = 52
	KeyFive       KeyType = 53
	KeySix        KeyType = 54
	KeySeven      KeyType = 55
	KeyEight      KeyType = 56
	KeyNine       KeyType = 57
	KeySemicolon  KeyType = 59
	KeyEqual      KeyType = 61
	KeyA          KeyType = 65
	KeyB          KeyType = 66
	KeyC          KeyType = 67
	KeyD          KeyType = 68
	KeyE          KeyType = 69
	KeyF          KeyType = 70
	KeyG          KeyType = 71
	KeyH          KeyType = 72
	KeyI          KeyType = 73
	KeyJ          KeyType = 74
	KeyK          KeyType = 75
	KeyL          KeyType = 76
	KeyM          KeyType = 77
	KeyN          KeyType = 78
	KeyO          KeyType = 79
	KeyP          KeyType = 80
	KeyQ          KeyType = 81
	KeyR          KeyType = 82
	KeyS          KeyType = 83
	KeyT          KeyType = 84
	KeyU          KeyType = 85
	KeyV          KeyType = 86
	KeyW          KeyType = 87
	KeyX          KeyType = 88
	KeyY          KeyType = 89
	KeyZ          KeyType = 90

	// Android keys
	KeyBack       KeyType = 4
	KeyMenu       KeyType = 5
	KeyVolumeUp   KeyType = 24
	KeyVolumeDown KeyType = 25
)

// Mouse Buttons
type MouseButtonType int32

const (
	MouseButtonLeft MouseButtonType = iota
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonSide
	MouseButtonExtra
	MouseButtonForward
	MouseButtonBack
	MouseButtonNone // keep last

	MouseLeftButton   = MouseButtonLeft
	MouseRightButton  = MouseButtonRight
	MouseMiddleButton = MouseButtonMiddle
)

// Mouse cursor
type MouseCursorType int32

const (
	MouseCursorDefault      MouseCursorType = iota // Default pointer shape
	MouseCursorArrow                               // Arrow shape
	MouseCursorIBeam                               // Text writing cursor shape
	MouseCursorCrosshair                           // Cross shape
	MouseCursorPointingHand                        // Pointing hand cursor
	MouseCursorResizeEW                            // Horizontal resize/move arrow shape
	MouseCursorResizeNS                            // Vertical resize/move arrow shape
	MouseCursorResizeNWSE                          // Top-left to bottom-right diagonal resize/move arrow shape
	MouseCursorResizeNESW                          // The top-right to bottom-left diagonal resize/move arrow shape
	MouseCursorResizeAll                           // The omni-directional resize/move cursor shape
	MouseCursorNotAllowed                          // The operation-not-allowed shape
)

// Gamepad Buttons
type GamepadButtonType int32

const (
	GamepadButtonUnknown        GamepadButtonType = iota // Unknown button, just for error checking
	GamepadButtonLeftFaceUp                              // Gamepad left DPAD up button
	GamepadButtonLeftFaceRight                           // Gamepad left DPAD right button
	GamepadButtonLeftFaceDown                            // Gamepad left DPAD down button
	GamepadButtonLeftFaceLeft                            // Gamepad left DPAD left button
	GamepadButtonRightFaceUp                             // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	GamepadButtonRightFaceRight                          // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	GamepadButtonRightFaceDown                           // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	GamepadButtonRightFaceLeft                           // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
	GamepadButtonLeftTrigger1                            // Gamepad top/back trigger left (first), it could be a trailing button
	GamepadButtonLeftTrigger2                            // Gamepad top/back trigger left (second), it could be a trailing button
	GamepadButtonRightTrigger1                           // Gamepad top/back trigger right (one), it could be a trailing button
	GamepadButtonRightTrigger2                           // Gamepad top/back trigger right (second), it could be a trailing button
	GamepadButtonMiddleLeft                              // Gamepad center buttons, left one (i.e. PS3: Select)
	GamepadButtonMiddle                                  // Gamepad center buttons, middle one (i.e. PS3: PS, Xbox: XBOX)
	GamepadButtonMiddleRight                             // Gamepad center buttons, right one (i.e. PS3: Start)
	GamepadButtonLeftThumb                               // Gamepad joystick pressed button left
	GamepadButtonRightThumb                              // Gamepad joystick pressed button right
)

// Gamepad Axis
type GamepadAxisType int32

const (
	GamepadAxisLeftX        GamepadAxisType = iota // Gamepad left stick X axis
	GamepadAxisLeftY                               // Gamepad left stick Y axis
	GamepadAxisRightX                              // Gamepad right stick X axis
	GamepadAxisRightY                              // Gamepad right stick Y axis
	GamepadAxisLeftTrigger                         // Gamepad back trigger left, pressure level: [1..-1]
	GamepadAxisRightTrigger                        // Gamepad back trigger right, pressure level: [1..-1]
)

type InputDeviceType int8

const (
	Keyboard InputDeviceType = iota
	Mouse
	Gamepad
	Gesture
)

type InputDeviceMask int32

const InputDeviceMaskShift = 28
const InputGamepadIndexMaskWidth = 4 // Up to 16 gamepads
const InputKeyMask = (1 << InputDeviceMaskShift) - 1

const (
	KeyboardMask InputDeviceMask = InputDeviceMask(Keyboard) << InputDeviceMaskShift
	MouseMask                    = InputDeviceMask(Mouse) << InputDeviceMaskShift
	GamepadMask                  = InputDeviceMask(Gamepad) << InputDeviceMaskShift
	GestureMask                  = InputDeviceMask(Gesture) << InputDeviceMaskShift
)

type UnifiedKeyType int32

func (k UnifiedKeyType) Device() InputDeviceType {
	return InputDeviceType(k >> InputDeviceMaskShift)
}

func (k UnifiedKeyType) Keyboard() KeyType {
	return KeyType(k & InputKeyMask)
}

func (k UnifiedKeyType) Mouse() MouseButtonType {
	return MouseButtonType(k & InputKeyMask)
}

func (k UnifiedKeyType) Gamepad() GamepadButtonType {
	return GamepadButtonType(k & InputKeyMask)
}

func (k UnifiedKeyType) GamepadIndex() int {
	return int(k&InputKeyMask) >> (InputDeviceMaskShift - InputGamepadIndexMaskWidth)
}

func (k UnifiedKeyType) Gesture() Gestures {
	switch k {
	case Gesture_Tap:
		return GestureTap
	case Gesture_DoubleTap:
		return GestureDoubleTap
	case Gesture_SwipeRight:
		return GestureSwipeRight
	case Gesture_SwipeLeft:
		return GestureSwipeLeft
	case Gesture_SwipeUp:
		return GestureSwipeUp
	case Gesture_SwipeDown:
		return GestureSwipeDown
	}

	return GestureNone
}

func (k UnifiedKeyType) IsPressed() bool {
	switch k.Device() {
	case Keyboard:
		return IsKeyPressed(k.Keyboard())
	case Mouse:
		return IsMouseButtonPressed(k.Mouse())
	case Gamepad:
		return IsGamepadButtonPressed(k.GamepadIndex(), k.Gamepad())
	case Gesture:
		return IsGestureDetected(k.Gesture())
	}
	return false
}

func (k UnifiedKeyType) IsDown() bool {
	switch k.Device() {
	case Keyboard:
		return IsKeyDown(k.Keyboard())
	case Mouse:
		return IsMouseButtonDown(k.Mouse())
	case Gamepad:
		return IsGamepadButtonDown(k.GamepadIndex(), k.Gamepad())
	case Gesture:
		return IsGestureDetected(k.Gesture())
	}
	return false
}

func (k UnifiedKeyType) IsReleased() bool {
	switch k.Device() {
	case Keyboard:
		return IsKeyReleased(k.Keyboard())
	case Mouse:
		return IsMouseButtonReleased(k.Mouse())
	case Gamepad:
		return IsGamepadButtonReleased(k.GamepadIndex(), k.Gamepad())
	case Gesture:
		return IsGestureDetected(k.Gesture())
	}
	return false
}

func (k UnifiedKeyType) IsUp() bool {
	switch k.Device() {
	case Keyboard:
		return IsKeyUp(k.Keyboard())
	case Mouse:
		return IsMouseButtonUp(k.Mouse())
	case Gamepad:
		return IsGamepadButtonUp(k.GamepadIndex(), k.Gamepad())
	case Gesture:
		return IsGestureDetected(k.Gesture())
	}
	return false
}

func (k UnifiedKeyType) IsEvent(event InputEventType) bool {
	if k == Any_KeyNone {
		return true
	}

	switch event {
	case InputEventPressed:
		return k.IsPressed()
	case InputEventDown:
		return k.IsDown()
	case InputEventReleased:
		return k.IsReleased()
	case InputEventUp:
		return k.IsUp()
	}
	return false
}

const (
	// Any_KeyNone is used for no key - any check will return true
	Any_KeyNone = 0

	// Keyboard Function Keys
	Keyboard_KeySpace        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeySpace)
	Keyboard_KeyEscape       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyEscape)
	Keyboard_KeyEnter        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyEnter)
	Keyboard_KeyTab          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyTab)
	Keyboard_KeyBackspace    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyBackspace)
	Keyboard_KeyInsert       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyInsert)
	Keyboard_KeyDelete       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyDelete)
	Keyboard_KeyRight        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRight)
	Keyboard_KeyLeft         = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeft)
	Keyboard_KeyDown         = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyDown)
	Keyboard_KeyUp           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyUp)
	Keyboard_KeyPageUp       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyPageUp)
	Keyboard_KeyPageDown     = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyPageDown)
	Keyboard_KeyHome         = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyHome)
	Keyboard_KeyEnd          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyEnd)
	Keyboard_KeyCapsLock     = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyCapsLock)
	Keyboard_KeyScrollLock   = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyScrollLock)
	Keyboard_KeyNumLock      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyNumLock)
	Keyboard_KeyPrintScreen  = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyPrintScreen)
	Keyboard_KeyPause        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyPause)
	Keyboard_KeyF1           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF1)
	Keyboard_KeyF2           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF2)
	Keyboard_KeyF3           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF3)
	Keyboard_KeyF4           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF4)
	Keyboard_KeyF5           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF5)
	Keyboard_KeyF6           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF6)
	Keyboard_KeyF7           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF7)
	Keyboard_KeyF8           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF8)
	Keyboard_KeyF9           = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF9)
	Keyboard_KeyF10          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF10)
	Keyboard_KeyF11          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF11)
	Keyboard_KeyF12          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF12)
	Keyboard_KeyLeftShift    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeftShift)
	Keyboard_KeyLeftControl  = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeftControl)
	Keyboard_KeyLeftAlt      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeftAlt)
	Keyboard_KeyLeftSuper    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeftSuper)
	Keyboard_KeyRightShift   = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRightShift)
	Keyboard_KeyRightControl = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRightControl)
	Keyboard_KeyRightAlt     = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRightAlt)
	Keyboard_KeyRightSuper   = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRightSuper)
	Keyboard_KeyKbMenu       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKbMenu)
	Keyboard_KeyLeftBracket  = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyLeftBracket)
	Keyboard_KeyBackSlash    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyBackSlash)
	Keyboard_KeyRightBracket = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyRightBracket)
	Keyboard_KeyGrave        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyGrave)

	// Keyboard Number Pad Keys
	Keyboard_KeyKp0        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp0)
	Keyboard_KeyKp1        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp1)
	Keyboard_KeyKp2        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp2)
	Keyboard_KeyKp3        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp3)
	Keyboard_KeyKp4        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp4)
	Keyboard_KeyKp5        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp5)
	Keyboard_KeyKp6        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp6)
	Keyboard_KeyKp7        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp7)
	Keyboard_KeyKp8        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp8)
	Keyboard_KeyKp9        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKp9)
	Keyboard_KeyKpDecimal  = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpDecimal)
	Keyboard_KeyKpDivide   = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpDivide)
	Keyboard_KeyKpMultiply = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpMultiply)
	Keyboard_KeyKpSubtract = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpSubtract)
	Keyboard_KeyKpAdd      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpAdd)
	Keyboard_KeyKpEnter    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpEnter)
	Keyboard_KeyKpEqual    = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyKpEqual)

	// Keyboard Alpha Numeric Keys
	Keyboard_KeyApostrophe = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyApostrophe)
	Keyboard_KeyComma      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyComma)
	Keyboard_KeyMinus      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyMinus)
	Keyboard_KeyPlus       = Keyboard_KeyEqual
	Keyboard_KeyPeriod     = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyPeriod)
	Keyboard_KeySlash      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeySlash)
	Keyboard_KeyZero       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyZero)
	Keyboard_KeyOne        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyOne)
	Keyboard_KeyTwo        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyTwo)
	Keyboard_KeyThree      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyThree)
	Keyboard_KeyFour       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyFour)
	Keyboard_KeyFive       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyFive)
	Keyboard_KeySix        = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeySix)
	Keyboard_KeySeven      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeySeven)
	Keyboard_KeyEight      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyEight)
	Keyboard_KeyNine       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyNine)
	Keyboard_KeySemicolon  = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeySemicolon)
	Keyboard_KeyEqual      = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyEqual)
	Keyboard_KeyA          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyA)
	Keyboard_KeyB          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyB)
	Keyboard_KeyC          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyC)
	Keyboard_KeyD          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyD)
	Keyboard_KeyE          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyE)
	Keyboard_KeyF          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyF)
	Keyboard_KeyG          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyG)
	Keyboard_KeyH          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyH)
	Keyboard_KeyI          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyI)
	Keyboard_KeyJ          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyJ)
	Keyboard_KeyK          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyK)
	Keyboard_KeyL          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyL)
	Keyboard_KeyM          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyM)
	Keyboard_KeyN          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyN)
	Keyboard_KeyO          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyO)
	Keyboard_KeyP          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyP)
	Keyboard_KeyQ          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyQ)
	Keyboard_KeyR          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyR)
	Keyboard_KeyS          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyS)
	Keyboard_KeyT          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyT)
	Keyboard_KeyU          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyU)
	Keyboard_KeyV          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyV)
	Keyboard_KeyW          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyW)
	Keyboard_KeyX          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyX)
	Keyboard_KeyY          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyY)
	Keyboard_KeyZ          = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyZ)

	// Android keys
	Keyboard_KeyBack       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyBack)
	Keyboard_KeyMenu       = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyMenu)
	Keyboard_KeyVolumeUp   = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyVolumeUp)
	Keyboard_KeyVolumeDown = UnifiedKeyType(KeyboardMask) | UnifiedKeyType(KeyVolumeDown)

	// Mouse buttons
	Mouse_ButtonLeft    = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonLeft)
	Mouse_ButtonRight   = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonRight)
	Mouse_ButtonMiddle  = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonMiddle)
	Mouse_ButtonSide    = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonSide)
	Mouse_ButtonExtra   = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonExtra)
	Mouse_ButtonForward = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonForward)
	Mouse_ButtonBack    = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonBack)
	Mouse_ButtonNone    = UnifiedKeyType(MouseMask) | UnifiedKeyType(MouseButtonNone) // keep last

	// Gamepad buttons
	Gamepad_ButtonUnknown        = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonUnknown)
	Gamepad_ButtonLeftFaceUp     = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftFaceUp)
	Gamepad_ButtonLeftFaceRight  = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftFaceRight)
	Gamepad_ButtonLeftFaceDown   = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftFaceDown)
	Gamepad_ButtonLeftFaceLeft   = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftFaceLeft)
	Gamepad_ButtonRightFaceUp    = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightFaceUp)
	Gamepad_ButtonRightFaceRight = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightFaceRight)
	Gamepad_ButtonRightFaceDown  = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightFaceDown)
	Gamepad_ButtonRightFaceLeft  = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightFaceLeft)
	Gamepad_ButtonLeftTrigger1   = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftTrigger1)
	Gamepad_ButtonLeftTrigger2   = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftTrigger2)
	Gamepad_ButtonRightTrigger1  = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightTrigger1)
	Gamepad_ButtonRightTrigger2  = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightTrigger2)
	Gamepad_ButtonMiddleLeft     = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonMiddleLeft)
	Gamepad_ButtonMiddle         = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonMiddle)
	Gamepad_ButtonMiddleRight    = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonMiddleRight)
	Gamepad_ButtonLeftThumb      = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonLeftThumb)
	Gamepad_ButtonRightThumb     = UnifiedKeyType(GamepadMask) | UnifiedKeyType(GamepadButtonRightThumb)
)
const (
	// Gestures which can be processed as virtual key inputs. These will only respond any event like on Pressed event
	Gesture_Tap = iota + UnifiedKeyType(GestureMask)
	Gesture_DoubleTap
	Gesture_SwipeRight
	Gesture_SwipeLeft
	Gesture_SwipeUp
	Gesture_SwipeDown
)

// Some Basic Colors
// NOTE: Custom raylib color palette for amazing visuals on WHITE background
var (
	// Light Gray
	LightGray = NewColor(200, 200, 200, 255)
	// Gray
	Gray = NewColor(130, 130, 130, 255)
	// Dark Gray
	DarkGray = NewColor(80, 80, 80, 255)
	// Darkest Gray
	DarkestGray = NewColor(40, 40, 40, 255)
	// Black Gray - the darkset of all Grays
	BlackGray = NewColor(20, 20, 20, 255)
	// Yellow
	Yellow = NewColor(253, 249, 0, 255)
	// Gold
	Gold = NewColor(255, 203, 0, 255)
	// Orange
	Orange = NewColor(255, 161, 0, 255)
	// Pink
	Pink = NewColor(255, 109, 194, 255)
	// Red
	Red = NewColor(230, 41, 55, 255)
	// Maroon
	Maroon = NewColor(190, 33, 55, 255)
	// Green
	Green = NewColor(0, 228, 48, 255)
	// Lime
	Lime = NewColor(0, 158, 47, 255)
	// Dark Green
	DarkGreen = NewColor(0, 117, 44, 255)
	// Sky Blue
	SkyBlue = NewColor(102, 191, 255, 255)
	// Blue
	Blue = NewColor(0, 121, 241, 255)
	// Dark Blue
	DarkBlue = NewColor(0, 82, 172, 255)
	// Purple
	Purple = NewColor(200, 122, 255, 255)
	// Violet
	Violet = NewColor(135, 60, 190, 255)
	// Dark Purple
	DarkPurple = NewColor(112, 31, 126, 255)
	// Beige
	Beige = NewColor(211, 176, 131, 255)
	// Brown
	Brown = NewColor(127, 106, 79, 255)
	// Dark Brown
	DarkBrown = NewColor(76, 63, 47, 255)
	// White
	White = NewColor(255, 255, 255, 255)
	// Black
	Black = NewColor(0, 0, 0, 255)
	// Blank (Transparent)
	Blank = NewColor(0, 0, 0, 0)
	// Magenta
	Magenta = NewColor(255, 0, 255, 255)
	// Ray White (RayLib Logo White)
	RayWhite = NewColor(245, 245, 245, 255)
)

type Vector4 = vector4.Float32

var (
	AnchorTopLeft      = vector2.NewFloat32(0, 0)
	AnchorTopRight     = vector2.NewFloat32(1, 0)
	AnchorTopCenter    = vector2.NewFloat32(0.5, 0)
	AnchorCenter       = vector2.NewFloat32(0.5, 0.5)
	AnchorBottomCenter = vector2.NewFloat32(0.5, 1)
	AnchorBottomLeft   = vector2.NewFloat32(0, 1)
	AnchorBottomRight  = vector2.NewFloat32(1, 1)
)

// Matrix type (OpenGL style 4x4 - right handed, column major)
type Matrix struct {
	M0, M4, M8, M12  float32
	M1, M5, M9, M13  float32
	M2, M6, M10, M14 float32
	M3, M7, M11, M15 float32
}

// NewMatrix - Returns new Matrix
func NewMatrix(m0, m4, m8, m12, m1, m5, m9, m13, m2, m6, m10, m14, m3, m7, m11, m15 float32) Matrix {
	return Matrix{m0, m4, m8, m12, m1, m5, m9, m13, m2, m6, m10, m14, m3, m7, m11, m15}
}

// Mat2 type (used for polygon shape rotation matrix)
type Mat2 struct {
	M00 float32
	M01 float32
	M10 float32
	M11 float32
}

// NewMat2 - Returns new Mat2
func NewMat2(m0, m1, m10, m11 float32) Mat2 {
	return Mat2{m0, m1, m10, m11}
}

// Quaternion, 4 components (Vector4 alias)
type Quaternion = vector4.Float32

// NewQuaternion - Returns new Quaternion
func NewQuaternion(x, y, z, w float32) Quaternion {
	return vector4.NewFloat32(x, y, z, w)
}

// NewColor - Returns new Color
func NewColor(r, g, b, a uint8) colorex.RGBA {
	return colorex.RGBA{R: r, G: g, B: b, A: a}
}

// Rectangle type
type RectangleInt32 = rect2.Int32

// NewRectangle - Returns new Rectangle
func NewRectangle[XT, YT, WT, HT mathex.Number](x XT, y YT, width WT, height HT) rect2.Float32 {
	return rect2.New(vector2.NewFloat32(x, y), vector2.NewFloat32(width, height))
}

func NewRectangleWHV[WHT mathex.SignedNumber](wh vector2.Vector[WHT]) rect2.Float32 {
	return rect2.New(vector2.Zero[float32](), wh.ToFloat32())
}

// Camera3D type, defines a camera position/orientation in 3d space
type Camera3D struct {
	// Camera position
	Position vector3.Float32
	// Camera target it looks-at
	Target vector3.Float32
	// Camera up vector (rotation over its axis)
	Up vector3.Float32
	// Camera field-of-view apperture in Y (degrees) in perspective, used as near plane width in orthographic
	Fovy float32
	// Camera type, controlling projection type, either CameraPerspective or CameraOrthographic.
	Projection CameraProjection
}

// Camera type fallback, defaults to Camera3D
type Camera = Camera3D

// NewCamera3D - Returns new Camera3D
func NewCamera3D(pos, target, up vector3.Float32, fovy float32, ct CameraProjection) Camera3D {
	return Camera3D{pos, target, up, fovy, ct}
}

// Camera2D type, defines a 2d camera
type Camera2D struct {
	// Camera offset (displacement from target)
	Offset vector2.Float32
	// Camera target (rotation and zoom origin)
	Target vector2.Float32
	// Camera rotation in degrees
	Rotation float32
	// Camera zoom (scaling), should be 1.0f by default
	Zoom float32
}

// NewCamera2D - Returns new Camera2D
func NewCamera2D(offset, target vector2.Float32, rotation, zoom float32) Camera2D {
	return Camera2D{offset, target, rotation, zoom}
}

// BoundingBox type
type BoundingBox struct {
	// Minimum vertex box-corner
	Min vector3.Float32
	// Maximum vertex box-corner
	Max vector3.Float32
}

// NewBoundingBox - Returns new BoundingBox
func NewBoundingBox(min, max vector3.Float32) BoundingBox {
	return BoundingBox{min, max}
}

// Asset file
type Asset interface {
	io.ReadSeeker
	io.Closer

	Size() int64
}

func ReadAll(a Asset) ([]byte, error) {
	b := make([]byte, a.Size())
	if _, err := a.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Gestures type
type Gestures uint32

// Gestures types
// NOTE: It could be used as flags to enable only some gestures
const (
	GestureNone       Gestures = 0
	GestureTap        Gestures = 1
	GestureDoubleTap  Gestures = 2
	GestureHold       Gestures = 4
	GestureDrag       Gestures = 8
	GestureSwipeRight Gestures = 16
	GestureSwipeLeft  Gestures = 32
	GestureSwipeUp    Gestures = 64
	GestureSwipeDown  Gestures = 128
	GesturePinchIn    Gestures = 256
	GesturePinchOut   Gestures = 512
)

// Shader location point type
const (
	ShaderLocVertexPosition = iota
	ShaderLocVertexTexcoord01
	ShaderLocVertexTexcoord02
	ShaderLocVertexNormal
	ShaderLocVertexTangent
	ShaderLocVertexColor
	ShaderLocMatrixMvp
	ShaderLocMatrixView
	ShaderLocMatrixProjection
	ShaderLocMatrixModel
	ShaderLocMatrixNormal
	ShaderLocVectorView
	ShaderLocColorDiffuse
	ShaderLocColorSpecular
	ShaderLocColorAmbient
	ShaderLocMapAlbedo
	ShaderLocMapMetalness
	ShaderLocMapNormal
	ShaderLocMapRoughness
	ShaderLocMapOcclusion
	ShaderLocMapEmission
	ShaderLocMapHeight
	ShaderLocMapCubemap
	ShaderLocMapIrradiance
	ShaderLocMapPrefilter
	ShaderLocMapBrdf

	ShaderLocMapDiffuse  = ShaderLocMapAlbedo
	ShaderLocMapSpecular = ShaderLocMapMetalness
)

// ShaderUniformDataType type
type ShaderUniformDataType int32

// ShaderUniformDataType enumeration
const (
	// Shader uniform type: float
	ShaderUniformFloat ShaderUniformDataType = iota
	// Shader uniform type: vec2 (2 float)
	ShaderUniformVec2
	// Shader uniform type: vec3 (3 float)
	ShaderUniformVec3
	// Shader uniform type: vec4 (4 float)
	ShaderUniformVec4
	// Shader uniform type: int
	ShaderUniformInt
	// Shader uniform type: ivec2 (2 int)
	ShaderUniformIvec2
	// Shader uniform type: ivec2 (3 int)
	ShaderUniformIvec3
	// Shader uniform type: ivec2 (4 int)
	ShaderUniformIvec4
	// Shader uniform type: unsigned int
	ShaderUniformUint
	// Shader uniform type: uivec2 (2 unsigned int)
	ShaderUniformUivec2
	// Shader uniform type: uivec3 (3 unsigned int)
	ShaderUniformUivec3
	// Shader uniform type: uivec4 (4 unsigned int)
	ShaderUniformUivec4
	// Shader uniform type: sampler2d
	ShaderUniformSampler2d
)

// Material map index
const (
	MapAlbedo = iota
	MapMetalness
	MapNormal
	MapRoughness
	MapOcclusion
	MapEmission
	MapHeight
	MapBrdg
	MapCubemap
	MapIrradiance
	MapPrefilter
	MapBrdf

	MapDiffuse  = MapAlbedo
	MapSpecular = MapMetalness
)

// Shader and material limits
const (
	// Maximum number of predefined locations stored in shader struct
	MaxShaderLocations = 32
	// Maximum number of texture maps stored in shader struct
	MaxMaterialMaps = 12
)

// Mesh - Vertex data definning a mesh
type Mesh struct {
	// Number of vertices stored in arrays
	VertexCount int32
	// Number of triangles stored (indexed or not)
	TriangleCount int32
	// Vertex position (XYZ - 3 components per vertex) (shader-location = 0)
	Vertices *float32
	// Vertex texture coordinates (UV - 2 components per vertex) (shader-location = 1)
	Texcoords *float32
	// Vertex second texture coordinates (useful for lightmaps) (shader-location = 5)
	Texcoords2 *float32
	// Vertex normals (XYZ - 3 components per vertex) (shader-location = 2)
	Normals *float32
	// Vertex tangents (XYZ - 3 components per vertex) (shader-location = 4)
	Tangents *float32
	// Vertex colors (RGBA - 4 components per vertex) (shader-location = 3)
	Colors *uint8
	// Vertex indices (in case vertex data comes indexed)
	Indices *uint16
	// AnimVertices
	AnimVertices *float32
	// AnimNormals
	AnimNormals *float32
	// BoneIds
	BoneIds *int32
	// BoneWeights
	BoneWeights *float32
	// Bones animated transformation matrices
	BoneMatrices *Matrix
	// Number of bones
	BoneCount int32
	// OpenGL Vertex Array Object id
	VaoID uint32
	// OpenGL Vertex Buffer Objects id (7 types of vertex data)
	VboID *uint32
}

func (m Mesh) IsValid() bool {
	if m.Vertices != nil && unsafe.Slice(m.VboID, 1)[0] == 0 {
		return false
	} // Vertex position buffer not uploaded to GPU
	if m.Texcoords != nil && unsafe.Slice(m.VboID, 2)[1] == 0 {
		return false
	} // Vertex textcoords buffer not uploaded to GPU
	if m.Normals != nil && unsafe.Slice(m.VboID, 3)[2] == 0 {
		return false
	} // Vertex normals buffer not uploaded to GPU
	if m.Colors != nil && unsafe.Slice(m.VboID, 4)[3] == 0 {
		return false
	} // Vertex colors buffer not uploaded to GPU
	if m.Tangents != nil && unsafe.Slice(m.VboID, 5)[4] == 0 {
		return false
	} // Vertex tangents buffer not uploaded to GPU
	if m.Texcoords2 != nil && unsafe.Slice(m.VboID, 6)[5] == 0 {
		return false
	} // Vertex texcoords2 buffer not uploaded to GPU
	if m.Indices != nil && unsafe.Slice(m.VboID, 7)[6] == 0 {
		return false
	} // Vertex indices buffer not uploaded to GPU
	if m.BoneIds != nil && unsafe.Slice(m.VboID, 8)[7] == 0 {
		return false
	} // Vertex boneIds buffer not uploaded to GPU
	if m.BoneWeights != nil && unsafe.Slice(m.VboID, 9)[8] == 0 {
		return false
	} // Vertex boneWeights buffer not uploaded to GPU

	// NOTE: Some OpenGL versions do not support VAO, so we don't check it
	//if m.VaoId == 0 { return false; break }

	return true
}

func (m *Mesh) Unload() {
	UnloadMesh(m)
}

// Material type
type Material struct {
	// Shader
	Shader Shader
	// Maps
	Maps *MaterialMap
	// Generic parameters (if required)
	Params [4]float32
}

// IsMaterialValid - Check if a material is valid (shader assigned, map textures loaded in GPU)
func (m Material) IsValid() bool {
	result := false

	if m.Maps != nil && // Validate material contain some map
		m.Shader.IsValid() { // Validate material shader is valid
		result = true
	}

	// TODO: Check if available maps contain loaded textures

	return result
}

func (m *Material) Unload() {
	UnloadMaterial(m)
}

// GetMap - Get pointer to MaterialMap by map type
func (m Material) GetMap(index int32) *MaterialMap {
	return (*MaterialMap)(unsafe.Pointer(uintptr(unsafe.Pointer(m.Maps)) + uintptr(index)*unsafe.Sizeof(MaterialMap{})))
}

// MaterialMap type
type MaterialMap struct {
	// Texture
	Texture Texture2D
	// Color
	Color colorex.RGBA
	// Value
	Value float32
}

// Model is struct of model, meshes, materials and animation data
type Model struct {
	// Local transform matrix
	Transform Matrix
	// Number of meshes
	MeshCount int32
	// Number of materials
	MaterialCount int32
	// Meshes array (c array)
	//
	// Use Model.GetMeshes instead (go slice)
	Meshes *Mesh
	// Materials array (c array)
	//
	// Use Model.GetMaterials instead (go slice)
	Materials *Material
	// Mesh material number
	MeshMaterial *int32
	// Number of bones
	BoneCount int32
	// Bones information (skeleton) (c array)
	//
	// Use Model.GetBones instead (go slice)
	Bones *BoneInfo
	// Bones base transformation (pose) (c array)
	//
	// Use Model.GetBindPose instead (go slice)
	BindPose *Transform
}

func (model Model) IsValid() bool {
	result := false

	if (model.Meshes != nil) && // Validate model contains some mesh
		(model.Materials != nil) && // Validate model contains some material (at least default one)
		(model.MeshMaterial != nil) && // Validate mesh-material linkage
		(model.MeshCount > 0) && // Validate mesh count
		(model.MaterialCount > 0) { // Validate material count
		result = true
	}

	// NOTE: Many elements could be validated from a model, including every model mesh VAO/VBOs
	// but some VBOs could not be used, it depends on Mesh vertex data
	for _, mesh := range model.GetMeshes() {
		if !mesh.IsValid() {
			return false
		}
	}

	return result
}

func (m *Model) Unload() {
	UnloadModel(m)
}

// GetMeshes returns the meshes of a model as go slice
func (m Model) GetMeshes() []Mesh {
	return unsafe.Slice(m.Meshes, m.MeshCount)
}

// GetMaterials returns the materials of a model as go slice
func (m Model) GetMaterials() []Material {
	return unsafe.Slice(m.Materials, m.MaterialCount)
}

// GetBones returns the bones information (skeleton) of a model as go slice
func (m Model) GetBones() []BoneInfo {
	return unsafe.Slice(m.Bones, m.BoneCount)
}

// GetBindPose returns the bones base transformation of a model as go slice
func (m Model) GetBindPose() []Transform {
	return unsafe.Slice(m.BindPose, m.BoneCount)
}

// BoneInfo type
type BoneInfo struct {
	Name   [32]int8
	Parent int32
}

// Transform type
type Transform struct {
	Translation vector3.Float32
	Rotation    Quaternion
	Scale       vector3.Float32
}

// Ray type (useful for raycast)
type Ray struct {
	// Ray position (origin)
	Position vector3.Float32
	// Ray direction
	Direction vector3.Float32
}

// NewRay - Returns new Ray
func NewRay(position, direction vector3.Float32) Ray {
	return Ray{position, direction}
}

// ModelAnimation type
type ModelAnimation struct {
	BoneCount  int32
	FrameCount int32
	Bones      *BoneInfo
	FramePoses **Transform
	Name       [32]uint8
}

func (m *ModelAnimation) Unload() {
	UnloadModelAnimation(m)
}

// GetBones returns the bones information (skeleton) of a ModelAnimation as go slice
func (m ModelAnimation) GetBones() []BoneInfo {
	return unsafe.Slice(m.Bones, m.BoneCount)
}

// GetFramePose returns the Transform for a specific bone at a specific frame
func (m ModelAnimation) GetFramePose(frame, bone int) Transform {
	framePoses := unsafe.Slice(m.FramePoses, m.FrameCount)
	return unsafe.Slice(framePoses[frame], m.BoneCount)[bone]
}

// GetName returns the ModelAnimation's name as go string
func (m ModelAnimation) GetName() string {
	var end int
	for end = range m.Name {
		if m.Name[end] == 0 {
			break
		}
	}
	return string(m.Name[:end])
}

// RayCollision type - ray hit information
type RayCollision struct {
	Hit      bool
	Distance float32
	Point    vector3.Float32
	Normal   vector3.Float32
}

// NewRayCollision - Returns new RayCollision
func NewRayCollision(hit bool, distance float32, point, normal vector3.Float32) RayCollision {
	return RayCollision{hit, distance, point, normal}
}

// BlendMode type
type BlendMode int32

// Color blending modes (pre-defined)
const (
	BlendAlpha            BlendMode = iota // Blend textures considering alpha (default)
	BlendAdditive                          // Blend textures adding colors
	BlendMultiplied                        // Blend textures multiplying colors
	BlendAddColors                         // Blend textures adding colors (alternative)
	BlendSubtractColors                    // Blend textures subtracting colors (alternative)
	BlendAlphaPremultiply                  // Blend premultiplied textures considering alpha
	BlendCustom                            // Blend textures using custom src/dst factors
	BlendCustomSeparate                    // Blend textures using custom rgb/alpha separate src/dst factors
)

// Shader type (generic shader)
type Shader struct {
	// Shader program id
	ID uint32
	// Shader locations array
	Locs *int32
}

// NewShader - Returns new Shader
func NewShader(id uint32, locs *int32) Shader {
	return Shader{id, locs}
}

func (s Shader) IsValid() bool {
	return s.ID > 0 && // Validate shader id (GPU loaded successfully)
		s.Locs != nil
}

func (m *Shader) Unload() {
	UnloadShader(m)
}

// GetLocation - Get shader value's location
func (s Shader) GetLocation(index int32) int32 {
	return *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s.Locs)) + uintptr(index*4)))
}

// UpdateLocation - Update shader value's location
func (s *Shader) UpdateLocation(index int32, loc int32) {
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s.Locs)) + uintptr(index*4))) = loc
}

// GlyphInfo - Font character info
type GlyphInfo struct {
	// Character value (Unicode)
	Value int32
	// Character offset X when drawing
	OffsetX int32
	// Character offset Y when drawing
	OffsetY int32
	// Character advance position X
	AdvanceX int32
	// Character image data
	Image Image
}

// NewGlyphInfo - Returns new CharInfo
func NewGlyphInfo(value int32, offsetX, offsetY, advanceX int32, image Image) GlyphInfo {
	return GlyphInfo{value, offsetX, offsetY, advanceX, image}
}

// Font type, defines generation method
const (
	FontDefault = iota // Default font generation, anti-aliased
	FontBitmap         // Bitmap font generation, no anti-aliasing
	FontSdf            // SDF font generation, requires external shader
)

// Font type, includes texture and charSet array data
type Font struct {
	// Base size (default chars height)
	BaseSize int32
	// Number of characters
	GlyphCount int32
	// Padding around the chars
	GlyphPadding int32
	// Characters texture atlas
	Texture Texture2D
	// Characters rectangles in texture
	Recs *rect2.Float32
	// Characters info data
	Glyphs *GlyphInfo
}

func (f Font) IsValid() bool {
	return f.Texture.IsValid() && // Validate OpenGL id fot font texture atlas
		f.BaseSize != 0 && // Validate font size
		f.GlyphCount != 0 && // Validate font contains some glyph
		f.Recs != nil && // Validate font recs defining glyphs on texture atlas
		f.Glyphs != nil // Validate glyph data is loaded
}

func (m *Font) Unload() {
	UnloadFont(m)
}

// DrawTextEx - Draw text using Font and additional parameters
func (f Font) DrawEx(text string, position vector2.Float32, fontSize float32, spacing float32, tint colorex.RGBA) {
	DrawTextEx(f, text, position, fontSize, spacing, tint)
}

func (f Font) DrawLayout(text string, fontSize float32, spacing float32, tint colorex.RGBA, layoutFn func(wh vector2.Float32) rect2.Float32) {
	DrawTextLayout(f, text, fontSize, spacing, tint, layoutFn)
}

// MeasureTextEx - Measure string size for Font
func (f Font) MeasureEx(text string, fontSize float32, spacing float32) vector2.Float32 {
	return MeasureTextEx(f, text, fontSize, spacing)
}

func (f Font) GetRecs() []rect2.Float32 {
	return unsafe.Slice(f.Recs, f.GlyphCount)
}

// Font type, includes texture and charSet array data
type FontPreset struct {
	Font

	FontSize    uint16
	Spacing     uint16
	LineSpacing uint16
}

func (f FontPreset) DrawEx(text string, position vector2.Float32, tint colorex.RGBA) {
	DrawTextEx(f.Font, text, position, float32(f.FontSize), float32(f.Spacing), tint)
}

func (f FontPreset) DrawLayout(text string, tint colorex.RGBA, layoutFn func(wh vector2.Float32) rect2.Float32) {
	DrawTextLayout(f.Font, text, float32(f.FontSize), float32(f.Spacing), tint, layoutFn)
}

// MeasureTextEx - Measure string size for Font
func (f FontPreset) MeasureEx(text string) vector2.Float32 {
	return MeasureTextEx(f.Font, text, float32(f.FontSize), float32(f.Spacing))
}

// PixelFormat - Texture format
type PixelFormat int32

// Texture formats
// NOTE: Support depends on OpenGL version and platform
const (
	UncompressedGrayscale    PixelFormat = iota + 1 // 8 bit per pixel (no alpha)
	UncompressedGrayAlpha                           // 8*2 bpp (2 channels)
	UncompressedR5G6B5                              // 16 bpp
	UncompressedR8G8B8                              // 24 bpp
	UncompressedR5G5B5A1                            // 16 bpp (1 bit alpha)
	UncompressedR4G4B4A4                            // 16 bpp (4 bit alpha)
	UncompressedR8G8B8A8                            // 32 bpp
	UncompressedR32                                 // 32 bpp (1 channel - float)
	UncompressedR32G32B32                           // 32*3 bpp (3 channels - float)
	UncompressedR32G32B32A32                        // 32*4 bpp (4 channels - float)
	UncompressedR16                                 // 16 bpp (1 channel - half float)
	UncompressedR16G16B16                           // 16*3 bpp (3 channels - half float)
	UncompressedR16G16B16A16                        // 16*4 bpp (4 channels - half float)
	UncompressedR16UI                               // 16 bit per pixel (no alpha)
	UncompressedR32UI                               // 32 bit per pixel (no alpha)
	CompressedDxt1Rgb                               // 4 bpp (no alpha)
	CompressedDxt1Rgba                              // 4 bpp (1 bit alpha)
	CompressedDxt3Rgba                              // 8 bpp
	CompressedDxt5Rgba                              // 8 bpp
	CompressedEtc1Rgb                               // 4 bpp
	CompressedEtc2Rgb                               // 4 bpp
	CompressedEtc2EacRgba                           // 8 bpp
	CompressedPvrtRgb                               // 4 bpp
	CompressedPvrtRgba                              // 4 bpp
	CompressedAstc4x4Rgba                           // 8 bpp
	CompressedAstc8x8Rgba                           // 2 bpp
)

// TextureFilterMode - Texture filter mode
type TextureFilterMode int32

// Texture parameters: filter mode
// NOTE 1: Filtering considers mipmaps if available in the texture
// NOTE 2: Filter is accordingly set for minification and magnification
const (
	// No filter, just pixel aproximation
	FilterPoint TextureFilterMode = iota
	// Linear filtering
	FilterBilinear
	// Trilinear filtering (linear with mipmaps)
	FilterTrilinear
	// Anisotropic filtering 4x
	FilterAnisotropic4x
	// Anisotropic filtering 8x
	FilterAnisotropic8x
	// Anisotropic filtering 16x
	FilterAnisotropic16x
)

// TextureWrapMode - Texture wrap mode
type TextureWrapMode int32

// Texture parameters: wrap mode
const (
	WrapRepeat TextureWrapMode = iota
	WrapClamp
	WrapMirrorRepeat
	WrapMirrorClamp
)

// Cubemap layouts
const (
	CubemapLayoutAutoDetect       = iota // Automatically detect layout type
	CubemapLayoutLineVertical            // Layout is defined by a vertical line with faces
	CubemapLayoutLineHorizontal          // Layout is defined by a horizontal line with faces
	CubemapLayoutCrossThreeByFour        // Layout is defined by a 3x4 cross with cubemap faces
	CubemapLayoutCrossFourByThree        // Layout is defined by a 4x3 cross with cubemap faces
)

// Image type, bpp always RGBA (32bit)
// NOTE: Data stored in CPU memory (RAM)
type Image struct {
	// Image raw Data
	Data unsafe.Pointer
	// Image base width
	Width int32
	// Image base height
	Height int32
	// Mipmap levels, 1 by default
	Mipmaps int32
	// Data format (PixelFormat)
	Format PixelFormat
}

// NewImage - Returns new Image
func NewImage(data []byte, width, height, mipmaps int32, format PixelFormat) *Image {
	d := unsafe.Pointer(&data[0])

	return &Image{d, width, height, mipmaps, format}
}

// IsImageValid - Check if an image is valid (data and parameters)
func (image Image) IsValid() bool {
	return (image.Data != nil) && // Validate pixel data available
		(image.Width > 0) && // Validate image width
		(image.Height > 0) && // Validate image height
		(image.Format > 0) && // Validate image format
		(image.Mipmaps > 0) // Validate image mipmaps (at least 1 for basic mipmap level)
}

func (i *Image) Unload() {
	UnloadImage(i)
}

func (t Image) IsNull() bool {
	return t.cptr().data == nil
}

func (t Image) IsReady() bool {
	return !t.IsNull()
}

func (i Image) GetSize() vector2.Float32 {
	return vector2.NewFloat32(i.Width, i.Height)
}

func (i Image) GetRect() rect2.Float32 {
	return NewRectangle(0, 0, i.Width, i.Height)
}

func (i *Image) DrawDef(dst *Image, dstRect rect2.Float32) {
	ImageDraw(dst, *i, i.GetRect(), dstRect, White)
}

// Texture2D type, bpp always RGBA (32bit)
// NOTE: Data stored in GPU memory
type Texture2D struct {
	// OpenGL texture id
	ID uint32
	// Texture base width
	Width int32
	// Texture base height
	Height int32
	// Mipmap levels, 1 by default
	Mipmaps int32
	// Data format (PixelFormat)
	Format PixelFormat
}

// NewTexture2D - creates new empty texture of given size and format
func NewTexture2D[T constraints.Integer](width, height, mipmaps T, format PixelFormat) Texture2D {
	return LoadTextureFromImage(
		Image{
			Width:   int32(width),
			Height:  int32(height),
			Mipmaps: int32(mipmaps),
			Format:  format,
		})
}

func (t *Texture2D) Unload() {
	UnloadTexture(t)
}

func (t Texture2D) IsValidFast() bool {
	return t.ID > 0
}

func (t Texture2D) IsValid() bool {
	return t.ID > 0 && // Validate OpenGL id
		t.Width > 0 &&
		t.Height > 0 && // Validate texture size
		t.Format > 0 && // Validate texture pixel format
		t.Mipmaps > 0
}

func (t Texture2D) GetSize() vector2.Float32 {
	return vector2.NewFloat32(t.Width, t.Height)
}

func (t Texture2D) GetRect() rect2.Float32 {
	return NewRectangle(0, 0, t.Width, t.Height)
}

func (t Texture2D) Draw(posX int, posY int, tint colorex.RGBA) {
	DrawTexture(t, posX, posY, tint)
}

func (t Texture2D) DrawDef(posX int, posY int) {
	DrawTexture(t, posX, posY, White)
}

func (t Texture2D) DrawV(position vector2.Float32, tint colorex.RGBA) {
	DrawTextureV(t, position, tint)
}

func (t Texture2D) DrawVDef(position vector2.Float32) {
	DrawTextureV(t, position, White)
}

func (t Texture2D) DrawEx(position vector2.Float32, rotation, scale float32, tint colorex.RGBA) {
	DrawTextureEx(t, position, rotation, scale, tint)
}

func (t Texture2D) DrawExDef(position vector2.Float32) {
	DrawTextureEx(t, position, 0, 1, White)
}

func (t Texture2D) DrawRec(sourceRec rect2.Float32, position vector2.Float32, tint colorex.RGBA) {
	DrawTextureRec(t, sourceRec, position, tint)
}

func (t Texture2D) DrawPro(sourceRec, destRec rect2.Float32, origin vector2.Float32, rotation float32, tint colorex.RGBA) {
	DrawTexturePro(t, sourceRec, destRec, origin, rotation, tint)
}

func (t Texture2D) DrawFlippedPro(sourceRec, destRec rect2.Float32, origin vector2.Float32, rotation float32, tint colorex.RGBA) {
	sourceRec = sourceRec.ScaleByVectorF(vector2.NewFloat32(1, -1))
	sourceRec = sourceRec.SetY(float32(t.Height) + sourceRec.Height())
	DrawTexturePro(t, sourceRec, destRec, origin, rotation, tint)
}

func (t Texture2D) DrawProDef(destRec rect2.Float32) {
	DrawTexturePro(t, t.GetRect(), destRec, vector2.Zero[float32](), 0, White)
}

func (t Texture2D) DrawProFlippedDef(destRec rect2.Float32) {
	DrawTexturePro(t, t.GetRect().ScaleByVectorF(vector2.NewFloat32(1, -1)), destRec, vector2.Zero[float32](), 0, White)
}

func (t Texture2D) DrawTiled(source, dest rect2.Float32, origin vector2.Float32, rotation, scale float32, tint colorex.RGBA) {
	DrawTextureTiled(t, source, dest, origin, rotation, scale, tint)
}

func (t Texture2D) DrawTiledDef(dest rect2.Float32) {
	DrawTextureTiled(t, t.GetRect(), dest, vector2.Zero[float32](), 0, 1, White)
}

// RenderTexture2D type, for texture rendering
type RenderTexture2D struct {
	// Render texture (fbo) id
	ID uint32
	// Color buffer attachment texture
	Texture Texture2D
	// Depth buffer attachment texture
	Depth Texture2D
}

// NewRenderTexture2D - Returns new RenderTexture2D
func NewRenderTexture2D(id uint32, texture, depth Texture2D) *RenderTexture2D {
	return &RenderTexture2D{id, texture, depth}
}

func (r RenderTexture2D) IsValidFast() bool {
	return r.ID > 0
}

func (r RenderTexture2D) IsValid() bool {
	return r.ID > 0 &&
		r.Texture.IsValid() &&
		r.Depth.IsValid()
}

func (r *RenderTexture2D) Unload() {
	UnloadRenderTexture(r)
}

// TraceLogCallbackFun - function that will recive the trace log messages
type TraceLogCallbackFun func(int, string)

// TraceLogLevel parameter of trace log message
type TraceLogLevel int

// Trace log level
// NOTE: Organized by priority level
const (
	// Display all logs
	LogAll TraceLogLevel = iota
	// Trace logging, intended for internal use only
	LogTrace
	// Debug logging, used for internal debugging, it should be disabled on release builds
	LogDebug
	// Info logging, used for program execution info
	LogInfo
	// Warning logging, used on recoverable failures
	LogWarning
	// Error logging, used on unrecoverable failures
	LogError
	// Fatal logging, used to abort program: exit(EXIT_FAILURE)
	LogFatal
	// Disable logging
	LogNone
)

// N-patch layout
type NPatchLayout int32

const (
	NPatchNinePatch            NPatchLayout = iota // Npatch layout: 3x3 tiles
	NPatchThreePatchVertical                       // Npatch layout: 1x3 tiles
	NPatchThreePatchHorizontal                     // Npatch layout: 3x1 tiles
)

// NPatchInfo type, n-patch layout info
type NPatchInfo struct {
	Source rect2.Float32 // Texture source rectangle
	Left   int32         // Left border offset
	Top    int32         // Top border offset
	Right  int32         // Right border offset
	Bottom int32         // Bottom border offset
	Layout NPatchLayout  // Layout of the n-patch: 3x3, 1x3 or 3x1
}

func (i NPatchInfo) Default(d any) NPatchInfo {
	switch d := d.(type) {
	case Texture2D:
		i.Source = d.GetRect()
		iss := i.Source.Size
		if iss.X == iss.Y {
			ps := int32(iss.X / 3)
			i.Layout = NPatchNinePatch
			i.Left = ps
			i.Top = ps
			i.Right = ps
			i.Bottom = ps
		} else if iss.X > iss.Y {
			ps := int32(iss.X / 3)
			i.Layout = NPatchThreePatchHorizontal
			i.Left = ps
			i.Top = ps
			i.Right = ps
			i.Bottom = ps
		} else {
			ps := int32(iss.Y / 3)
			i.Layout = NPatchThreePatchVertical
			i.Left = ps
			i.Top = ps
			i.Right = ps
			i.Bottom = ps
		}
	}
	return i
}

// VrStereoConfig, VR stereo rendering configuration for simulator
type VrStereoConfig struct {
	Projection        [2]Matrix  // VR projection matrices (per eye)
	ViewOffset        [2]Matrix  // VR view offset matrices (per eye)
	LeftLensCenter    [2]float32 // VR left lens center
	RightLensCenter   [2]float32 // VR right lens center
	LeftScreenCenter  [2]float32 // VR left screen center
	RightScreenCenter [2]float32 // VR right screen center
	Scale             [2]float32 // VR distortion scale
	ScaleIn           [2]float32 // VR distortion scale in
}

// VrDeviceInfo, Head-Mounted-Display device parameters
type VrDeviceInfo struct {
	HResolution            int32      // Horizontal resolution in pixels
	VResolution            int32      // Vertical resolution in pixels
	HScreenSize            float32    // Horizontal size in meters
	VScreenSize            float32    // Vertical size in meters
	VScreenCenter          float32    // Screen center in meters
	EyeToScreenDistance    float32    // Distance between eye and display in meters
	LensSeparationDistance float32    // Lens separation distance in meters
	InterpupillaryDistance float32    // IPD (distance between pupils) in meters
	LensDistortionValues   [4]float32 // Lens distortion constant parameters
	ChromaAbCorrection     [4]float32 // Chromatic aberration correction parameters
}
