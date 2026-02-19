package rl

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/raylib/src/external/glfw/include -std=gnu99 -Wno-missing-braces -Wno-unused-result -Wno-implicit-function-declaration

#cgo nocallback BeginDrawing
#cgo nocallback BeginMode2D
#cgo nocallback BeginMode3D
#cgo nocallback BeginScissorMode
#cgo nocallback BeginTextureMode
#cgo nocallback BeginVrStereoMode
#cgo nocallback ClearBackground
#cgo nocallback ClearWindowState
#cgo nocallback CloseWindow
#cgo nocallback ColorAlphaBlend
#cgo nocallback ColorBrightness
#cgo nocallback ColorContrast
#cgo nocallback ColorFromHSV
#cgo nocallback ColorFromNormalized
#cgo nocallback ColorLerp
#cgo nocallback ColorTint
#cgo nocallback ColorToHSV
#cgo nocallback ColorToInt
#cgo nocallback CString
#cgo nocallback DisableEventWaiting
#cgo nocallback EnableEventWaiting
#cgo nocallback EndDrawing
#cgo nocallback EndMode2D
#cgo nocallback EndMode3D
#cgo nocallback EndScissorMode
#cgo nocallback EndTextureMode
#cgo nocallback EndVrStereoMode
#cgo nocallback ExportAutomationEventList
#cgo nocallback Fade
#cgo nocallback GetCameraMatrix
#cgo nocallback GetCameraMatrix2D
#cgo nocallback GetCharPressed
#cgo nocallback GetClipboardImage
#cgo nocallback GetClipboardText
#cgo nocallback GetColor
#cgo nocallback GetCurrentMonitor
#cgo nocallback GetFPS
#cgo nocallback GetFrameTime
#cgo nocallback GetGamepadAxisCount
#cgo nocallback GetGamepadAxisMovement
#cgo nocallback GetGamepadButtonPressed
#cgo nocallback GetGamepadName
#cgo nocallback GetKeyDownCount
#cgo nocallback GetKeyPressed
#cgo nocallback GetMonitorCount
#cgo nocallback GetMonitorHeight
#cgo nocallback GetMonitorName
#cgo nocallback GetMonitorPhysicalHeight
#cgo nocallback GetMonitorPhysicalWidth
#cgo nocallback GetMonitorPosition
#cgo nocallback GetMonitorRefreshRate
#cgo nocallback GetMonitorWidth
#cgo nocallback GetMouseDelta
#cgo nocallback GetMousePosition
#cgo nocallback GetMouseWheelMove
#cgo nocallback GetMouseWheelMoveV
#cgo nocallback GetMouseX
#cgo nocallback GetMouseY
#cgo nocallback GetPixelDataSize
#cgo nocallback GetRandomValue
#cgo nocallback GetRenderHeight
#cgo nocallback GetRenderWidth
#cgo nocallback GetScreenHeight
#cgo nocallback GetScreenToWorld2D
#cgo nocallback GetScreenToWorldRay
#cgo nocallback GetScreenToWorldRayEx
#cgo nocallback GetScreenWidth
#cgo nocallback GetShaderLocation
#cgo nocallback GetShaderLocationAttrib
#cgo nocallback GetTime
#cgo nocallback GetTouchPointCount
#cgo nocallback GetTouchPointId
#cgo nocallback GetTouchPosition
#cgo nocallback GetTouchX
#cgo nocallback GetTouchY
#cgo nocallback GetWindowHandle
#cgo nocallback GetWindowPosition
#cgo nocallback GetWindowScaleDPI
#cgo nocallback GetWorldToScreen
#cgo nocallback GetWorldToScreen2D
#cgo nocallback GetWorldToScreenEx
#cgo nocallback GoString
#cgo nocallback IsGamepadAvailable
#cgo nocallback IsGamepadButtonDown
#cgo nocallback IsGamepadButtonPressed
#cgo nocallback IsGamepadButtonReleased
#cgo nocallback IsGamepadButtonUp
#cgo nocallback IsKeyDown
#cgo nocallback IsKeyPressed
#cgo nocallback IsKeyPressedRepeat
#cgo nocallback IsKeyReleased
#cgo nocallback IsKeyUp
#cgo nocallback IsMouseButtonDown
#cgo nocallback IsMouseButtonPressed
#cgo nocallback IsMouseButtonReleased
#cgo nocallback IsMouseButtonUp
#cgo nocallback IsShaderValid
#cgo nocallback IsWindowFocused
#cgo nocallback IsWindowFullscreen
#cgo nocallback IsWindowHidden
#cgo nocallback IsWindowMaximized
#cgo nocallback IsWindowMinimized
#cgo nocallback IsWindowReady
#cgo nocallback IsWindowResized
#cgo nocallback IsWindowState
#cgo nocallback LoadAutomationEventList
#cgo nocallback LoadShaderFromMemory
#cgo nocallback LoadVrStereoConfig
#cgo nocallback MaximizeWindow
#cgo nocallback MinimizeWindow
#cgo nocallback OpenURL
#cgo nocallback PlayAutomationEvent
#cgo nocallback PollInputEvents
#cgo nocallback RestoreWindow
#cgo nocallback SetAutomationEventBaseFrame
#cgo nocallback SetAutomationEventList
#cgo nocallback SetClipboardText
#cgo nocallback SetConfigFlags
#cgo nocallback SetExitKey
#cgo nocallback SetGamepadMappings
#cgo nocallback SetGamepadVibration
#cgo nocallback SetMouseCursor
#cgo nocallback SetMouseOffset
#cgo nocallback SetMousePosition
#cgo nocallback SetMouseScale
#cgo nocallback SetShaderValue
#cgo nocallback SetShaderValueMatrix
#cgo nocallback SetShaderValueTexture
#cgo nocallback SetShaderValueV
#cgo nocallback SetTargetFPS
#cgo nocallback SetWindowIcon
#cgo nocallback SetWindowIcons
#cgo nocallback SetWindowMaxSize
#cgo nocallback SetWindowMinSize
#cgo nocallback SetWindowMonitor
#cgo nocallback SetWindowOpacity
#cgo nocallback SetWindowPosition
#cgo nocallback SetWindowSize
#cgo nocallback SetWindowState
#cgo nocallback SetWindowTitle
#cgo nocallback StartAutomationEventRecording
#cgo nocallback StopAutomationEventRecording
#cgo nocallback SwapScreenBuffer
#cgo nocallback TakeScreenshot
#cgo nocallback ToggleBorderlessWindowed
#cgo nocallback ToggleFullscreen
#cgo nocallback UnloadAutomationEventList
#cgo nocallback UnloadShader
#cgo nocallback UnloadVrStereoConfig
#cgo nocallback WaitTime
#cgo nocallback WindowShouldClose

#cgo noescape BeginDrawing
#cgo noescape BeginMode2D
#cgo noescape BeginMode3D
#cgo noescape BeginScissorMode
#cgo noescape BeginTextureMode
#cgo noescape BeginVrStereoMode
#cgo noescape ClearBackground
#cgo noescape ClearWindowState
#cgo noescape CloseWindow
#cgo noescape ColorAlphaBlend
#cgo noescape ColorBrightness
#cgo noescape ColorContrast
#cgo noescape ColorFromHSV
#cgo noescape ColorFromNormalized
#cgo noescape ColorLerp
#cgo noescape ColorTint
#cgo noescape ColorToHSV
#cgo noescape ColorToInt
#cgo noescape CString
#cgo noescape DisableEventWaiting
#cgo noescape EnableEventWaiting
#cgo noescape EndDrawing
#cgo noescape EndMode2D
#cgo noescape EndMode3D
#cgo noescape EndScissorMode
#cgo noescape EndTextureMode
#cgo noescape EndVrStereoMode
#cgo noescape ExportAutomationEventList
#cgo noescape Fade
#cgo noescape GetCameraMatrix
#cgo noescape GetCameraMatrix2D
#cgo noescape GetCharPressed
#cgo noescape GetClipboardImage
#cgo noescape GetClipboardText
#cgo noescape GetColor
#cgo noescape GetCurrentMonitor
#cgo noescape GetFPS
#cgo noescape GetFrameTime
#cgo noescape GetGamepadAxisCount
#cgo noescape GetGamepadAxisMovement
#cgo noescape GetGamepadButtonPressed
#cgo noescape GetGamepadName
#cgo noescape GetKeyDownCount
#cgo noescape GetKeyPressed
#cgo noescape GetMonitorCount
#cgo noescape GetMonitorHeight
#cgo noescape GetMonitorName
#cgo noescape GetMonitorPhysicalHeight
#cgo noescape GetMonitorPhysicalWidth
#cgo noescape GetMonitorPosition
#cgo noescape GetMonitorRefreshRate
#cgo noescape GetMonitorWidth
#cgo noescape GetMouseDelta
#cgo noescape GetMousePosition
#cgo noescape GetMouseWheelMove
#cgo noescape GetMouseWheelMoveV
#cgo noescape GetMouseX
#cgo noescape GetMouseY
#cgo noescape GetPixelDataSize
#cgo noescape GetRandomValue
#cgo noescape GetRenderHeight
#cgo noescape GetRenderWidth
#cgo noescape GetScreenHeight
#cgo noescape GetScreenToWorld2D
#cgo noescape GetScreenToWorldRay
#cgo noescape GetScreenToWorldRayEx
#cgo noescape GetScreenWidth
#cgo noescape GetShaderLocation
#cgo noescape GetShaderLocationAttrib
#cgo noescape GetTime
#cgo noescape GetTouchPointCount
#cgo noescape GetTouchPointId
#cgo noescape GetTouchPosition
#cgo noescape GetTouchX
#cgo noescape GetTouchY
#cgo noescape GetWindowHandle
#cgo noescape GetWindowPosition
#cgo noescape GetWindowScaleDPI
#cgo noescape GetWorldToScreen
#cgo noescape GetWorldToScreen2D
#cgo noescape GetWorldToScreenEx
#cgo noescape GoString
#cgo noescape IsGamepadAvailable
#cgo noescape IsGamepadButtonDown
#cgo noescape IsGamepadButtonPressed
#cgo noescape IsGamepadButtonReleased
#cgo noescape IsGamepadButtonUp
#cgo noescape IsKeyDown
#cgo noescape IsKeyPressed
#cgo noescape IsKeyPressedRepeat
#cgo noescape IsKeyReleased
#cgo noescape IsKeyUp
#cgo noescape IsMouseButtonDown
#cgo noescape IsMouseButtonPressed
#cgo noescape IsMouseButtonReleased
#cgo noescape IsMouseButtonUp
#cgo noescape IsShaderValid
#cgo noescape IsWindowFocused
#cgo noescape IsWindowFullscreen
#cgo noescape IsWindowHidden
#cgo noescape IsWindowMaximized
#cgo noescape IsWindowMinimized
#cgo noescape IsWindowReady
#cgo noescape IsWindowResized
#cgo noescape IsWindowState
#cgo noescape LoadAutomationEventList
#cgo noescape LoadShader
#cgo noescape LoadShaderFromMemory
#cgo noescape LoadVrStereoConfig
#cgo noescape MaximizeWindow
#cgo noescape MinimizeWindow
#cgo noescape OpenURL
#cgo noescape PlayAutomationEvent
#cgo noescape PollInputEvents
#cgo noescape RestoreWindow
#cgo noescape SetAutomationEventBaseFrame
#cgo noescape SetAutomationEventList
#cgo noescape SetClipboardText
#cgo noescape SetConfigFlags
#cgo noescape SetExitKey
#cgo noescape SetGamepadMappings
#cgo noescape SetGamepadVibration
#cgo noescape SetMouseCursor
#cgo noescape SetMouseOffset
#cgo noescape SetMousePosition
#cgo noescape SetMouseScale
#cgo noescape SetShaderValue
#cgo noescape SetShaderValueMatrix
#cgo noescape SetShaderValueTexture
#cgo noescape SetShaderValueV
#cgo noescape SetTargetFPS
#cgo noescape SetWindowIcon
#cgo noescape SetWindowIcons
#cgo noescape SetWindowMaxSize
#cgo noescape SetWindowMinSize
#cgo noescape SetWindowMonitor
#cgo noescape SetWindowOpacity
#cgo noescape SetWindowPosition
#cgo noescape SetWindowSize
#cgo noescape SetWindowState
#cgo noescape SetWindowTitle
#cgo noescape StartAutomationEventRecording
#cgo noescape StopAutomationEventRecording
#cgo noescape SwapScreenBuffer
#cgo noescape TakeScreenshot
#cgo noescape ToggleBorderlessWindowed
#cgo noescape ToggleFullscreen
#cgo noescape UnloadAutomationEventList
#cgo noescape UnloadShader
#cgo noescape UnloadVrStereoConfig
#cgo noescape WaitTime
#cgo noescape WindowShouldClose

#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/igadmg/gamemath/vector2"
	"github.com/igadmg/gamemath/vector3"
	"github.com/igadmg/gamemath/vector4"
	"github.com/igadmg/goex/image/colorex"
	"golang.org/x/exp/constraints"
)

// AutomationEvent - Automation event
type AutomationEvent C.AutomationEvent
type AutomationEventList C.AutomationEventList

// WindowShouldClose - Check if KeyEscape pressed or Close icon pressed
func WindowShouldClose() bool {
	ret := C.WindowShouldClose()
	v := bool(ret)
	return v
}

// CloseWindow - Close Window and Terminate Context
func CloseWindow() {
	C.CloseWindow()
}

// IsWindowReady - Check if window has been initialized successfully
func IsWindowReady() bool {
	ret := C.IsWindowReady()
	v := bool(ret)
	return v
}

// IsWindowFullscreen - Check if window is currently fullscreen
func IsWindowFullscreen() bool {
	ret := C.IsWindowFullscreen()
	v := bool(ret)
	return v
}

// IsWindowHidden - Check if window is currently hidden
func IsWindowHidden() bool {
	ret := C.IsWindowHidden()
	v := bool(ret)
	return v
}

// IsWindowMinimized - Check if window is currently minimized
func IsWindowMinimized() bool {
	ret := C.IsWindowMinimized()
	v := bool(ret)
	return v
}

// IsWindowMaximized - Check if window is currently maximized
func IsWindowMaximized() bool {
	ret := C.IsWindowMaximized()
	v := bool(ret)
	return v
}

// IsWindowFocused - Check if window is currently focused
func IsWindowFocused() bool {
	ret := C.IsWindowFocused()
	v := bool(ret)
	return v
}

// IsWindowResized - Check if window has been resized
func IsWindowResized() bool {
	ret := C.IsWindowResized()
	v := bool(ret)
	return v
}

// IsWindowState - Check if one specific window flag is enabled
func IsWindowState(flag uint32) bool {
	cflag := (C.uint)(flag)
	ret := C.IsWindowState(cflag)
	v := bool(ret)
	return v
}

// SetWindowState - Set window configuration state using flags
func SetWindowState(flags uint32) {
	cflags := (C.uint)(flags)
	C.SetWindowState(cflags)
}

// ClearWindowState - Clear window configuration state flags
func ClearWindowState(flags uint32) {
	cflags := (C.uint)(flags)
	C.ClearWindowState(cflags)
}

// ToggleFullscreen - Fullscreen toggle (only PLATFORM_DESKTOP)
func ToggleFullscreen() {
	C.ToggleFullscreen()
}

// ToggleBorderlessWindowed - Borderless fullscreen toggle (only PLATFORM_DESKTOP)
func ToggleBorderlessWindowed() {
	C.ToggleBorderlessWindowed()
}

// MaximizeWindow - Set window state: maximized, if resizable
func MaximizeWindow() {
	C.MaximizeWindow()
}

// MinimizeWindow - Set window state: minimized, if resizable
func MinimizeWindow() {
	C.MinimizeWindow()
}

// RestoreWindow - Set window state: not minimized/maximized
func RestoreWindow() {
	C.RestoreWindow()
}

// SetWindowIcon - Set icon for window (single image, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcon(image Image) {
	cimage := image.cptr()
	C.SetWindowIcon(*cimage)
}

// SetWindowIcons - Set icon for window (multiple images, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcons(images []Image, count int32) {
	cimages := (&images[0]).cptr()
	cimagesCount := C.int(count)
	C.SetWindowIcons(cimages, cimagesCount)
}

// SetWindowTitle - Set title for window (only PLATFORM_DESKTOP)
func SetWindowTitle(title string) {
	ctitle := textAlloc(title)
	C.SetWindowTitle(ctitle)
}

// SetWindowPosition - Set window position on screen (only PLATFORM_DESKTOP)
func SetWindowPosition(x, y int) {
	cx := (C.int)(x)
	cy := (C.int)(y)
	C.SetWindowPosition(cx, cy)
}

// SetWindowMonitor - Set monitor for the current window (fullscreen mode)
func SetWindowMonitor(monitor int) {
	cmonitor := (C.int)(monitor)
	C.SetWindowMonitor(cmonitor)
}

// SetWindowMinSize - Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(w, h int) {
	cw := (C.int)(w)
	ch := (C.int)(h)
	C.SetWindowMinSize(cw, ch)
}

// SetWindowMaxSize - Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMaxSize(w, h int) {
	cw := (C.int)(w)
	ch := (C.int)(h)
	C.SetWindowMaxSize(cw, ch)
}

// SetWindowSize - Set window dimensions
func SetWindowSize(w, h int) {
	cw := (C.int)(w)
	ch := (C.int)(h)
	C.SetWindowSize(cw, ch)
}

// SetWindowOpacity - Set window opacity [0.0f..1.0f] (only PLATFORM_DESKTOP)
func SetWindowOpacity(opacity float32) {
	copacity := (C.float)(opacity)
	C.SetWindowOpacity(copacity)
}

// GetWindowHandle - Get native window handle
func GetWindowHandle() unsafe.Pointer {
	v := C.GetWindowHandle()
	return v
}

// GetScreenWidth - Get current screen width
func GetScreenWidth() int {
	ret := C.GetScreenWidth()
	v := (int)(ret)
	return v
}

// GetScreenHeight - Get current screen height
func GetScreenHeight() int {
	ret := C.GetScreenHeight()
	v := (int)(ret)
	return v
}

func GetScreenSize() vector2.Int {
	return vector2.NewInt(
		GetScreenWidth(),
		GetScreenHeight(),
	)
}

// GetRenderWidth - Get current render width (it considers HiDPI)
func GetRenderWidth() int {
	ret := C.GetRenderWidth()
	v := (int)(ret)
	return v
}

// GetRenderHeight - Get current render height (it considers HiDPI)
func GetRenderHeight() int {
	ret := C.GetRenderHeight()
	v := (int)(ret)
	return v
}

func GetRenderSize() vector2.Int {
	return vector2.NewInt(
		GetRenderWidth(),
		GetRenderHeight(),
	)
}

// GetMonitorCount - Get number of connected monitors
func GetMonitorCount() int {
	ret := C.GetMonitorCount()
	v := (int)(ret)
	return v
}

// GetCurrentMonitor - Get current monitor where window is placed
func GetCurrentMonitor() int {
	ret := C.GetCurrentMonitor()
	v := (int)(ret)
	return v
}

// GetMonitorPosition - Get specified monitor position
func GetMonitorPosition(monitor int) vector2.Float32 {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorPosition(cmonitor)
	return *govec2ptr(&ret)
}

// GetMonitorWidth - Get primary monitor width
func GetMonitorWidth(monitor int) int {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorWidth(cmonitor)
	v := (int)(ret)
	return v
}

// GetMonitorHeight - Get primary monitor height
func GetMonitorHeight(monitor int) int {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorHeight(cmonitor)
	v := (int)(ret)
	return v
}

// GetMonitorPhysicalWidth - Get primary monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) int {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorPhysicalWidth(cmonitor)
	v := (int)(ret)
	return v
}

// GetMonitorPhysicalHeight - Get primary monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) int {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorPhysicalHeight(cmonitor)
	v := (int)(ret)
	return v
}

// GetMonitorRefreshRate - Get specified monitor refresh rate
func GetMonitorRefreshRate(monitor int) int {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorRefreshRate(cmonitor)
	v := (int)(ret)
	return v
}

// GetWindowPosition - Get window position XY on monitor
func GetWindowPosition() vector2.Float32 {
	ret := C.GetWindowPosition()
	return *govec2ptr(&ret)
}

// GetWindowScaleDPI - Get window scale DPI factor
func GetWindowScaleDPI() vector2.Float32 {
	ret := C.GetWindowScaleDPI()
	return *govec2ptr(&ret)
}

// GetDisplayScaleDPI - Get window scale DPI factor
func GetDisplayScaleDPI() vector2.Float32 {
	ret := C.GetDisplayScaleDPI()
	return *govec2ptr(&ret)
}

// GetMonitorName - Get the human-readable, UTF-8 encoded name of the primary monitor
func GetMonitorName(monitor int) string {
	cmonitor := (C.int)(monitor)
	ret := C.GetMonitorName(cmonitor)
	v := C.GoString(ret)
	return v
}

// SetClipboardText - Set clipboard text content
func SetClipboardText(data string) {
	cdata := textAlloc(data)
	C.SetClipboardText(cdata)
}

// GetClipboardText - Get clipboard text content
func GetClipboardText() string {
	ret := C.GetClipboardText()
	v := C.GoString(ret)
	return v
}

// GetClipboardImage - Get clipboard image content
//
// Only works with SDL3 backend or Windows with GLFW/RGFW
func GetClipboardImage() Image {
	ret := C.GetClipboardImage()
	v := newImageFromPointer(&ret)
	return *v
}

// EnableEventWaiting - Enable waiting for events on EndDrawing(), no automatic event polling
func EnableEventWaiting() {
	C.EnableEventWaiting()
}

// DisableEventWaiting - Disable waiting for events on EndDrawing(), automatic events polling
func DisableEventWaiting() {
	C.DisableEventWaiting()
}

// ClearBackground - Sets Background Color
func ClearBackground(col colorex.RGBA) {
	ccolor := ccolorptr(&col)
	C.ClearBackground(*ccolor)
}

// BeginDrawing - Setup drawing canvas to start drawing
func BeginDrawing() {
	C.BeginDrawing()
}

// EndDrawing - End canvas drawing and Swap Buffers (Double Buffering)
func EndDrawing() {
	C.EndDrawing()
}

// BeginMode2D - Initialize 2D mode with custom camera
func BeginMode2D(camera Camera2D) {
	ccamera := camera.cptr()
	C.BeginMode2D(*ccamera)
}

// EndMode2D - Ends 2D mode custom camera usage
func EndMode2D() {
	C.EndMode2D()
}

// BeginMode3D - Initializes 3D mode for drawing (Camera setup)
func BeginMode3D(camera Camera) {
	ccamera := camera.cptr()
	C.BeginMode3D(*ccamera)
}

// EndMode3D - Ends 3D mode and returns to default 2D orthographic mode
func EndMode3D() {
	C.EndMode3D()
}

// #cgo noescape BeginTextureMode
// #cgo nocallback BeginTextureMode
// BeginTextureMode - Initializes render texture for drawing
func BeginTextureMode(target RenderTexture2D) {
	ctarget := target.cptr()
	C.BeginTextureMode(*ctarget)
}

// EndTextureMode - Ends drawing to render texture
func EndTextureMode() {
	C.EndTextureMode()
}

// BeginScissorMode - Begins scissor mode (define screen area for following drawing)
func BeginScissorMode(x, y, width, height int32) {
	cx := (C.int)(x)
	cy := (C.int)(y)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	C.BeginScissorMode(cx, cy, cwidth, cheight)
}

func BeginScissorModeRec(r RectangleInt32) {
	cx := (C.int)(r.X())
	cy := (C.int)(r.Y())
	cwidth := (C.int)(r.Width())
	cheight := (C.int)(r.Height())
	C.BeginScissorMode(cx, cy, cwidth, cheight)
}

// EndScissorMode - Ends scissor mode
func EndScissorMode() {
	C.EndScissorMode()
}

// LoadShader - Load a custom shader and bind default locations
func LoadShader(vsFileName string, fsFileName string) Shader {
	cvsFileName := textAlloc(vsFileName)
	cfsFileName := textAlloc(fsFileName)

	if vsFileName == "" {
		cvsFileName = nil
	}

	if fsFileName == "" {
		cfsFileName = nil
	}

	ret := C.LoadShader(cvsFileName, cfsFileName)
	return *newShaderFromPointer(&ret)
}

// LoadShaderFromMemory - Load shader from code strings and bind default locations
func LoadShaderFromMemory(vsCode string, fsCode string) Shader {
	var cvsCode *C.char = nil
	var cfsCode *C.char = nil

	if vsCode != "" {
		cvsCode = C.CString(vsCode)
		defer C.free(unsafe.Pointer(cvsCode))
	}

	if fsCode != "" {
		cfsCode = C.CString(fsCode)
		defer C.free(unsafe.Pointer(cfsCode))
	}

	ret := C.LoadShaderFromMemory(cvsCode, cfsCode)
	return *newShaderFromPointer(&ret)
}

// IsShaderValid - Check if a shader is valid (loaded on GPU)
func IsShaderValid(shader Shader) bool {
	cshader := shader.cptr()
	ret := C.IsShaderValid(*cshader)
	v := bool(ret)
	return v
}

// GetShaderLocation - Get shader uniform location
func GetShaderLocation(shader Shader, uniformName string) int32 {
	cshader := shader.cptr()
	cuniformName := textAlloc(uniformName)

	ret := C.GetShaderLocation(*cshader, cuniformName)
	v := (int32)(ret)
	return v
}

// GetShaderLocationAttrib - Get shader attribute location
func GetShaderLocationAttrib(shader Shader, attribName string) int32 {
	cshader := shader.cptr()
	cuniformName := textAlloc(attribName)

	ret := C.GetShaderLocationAttrib(*cshader, cuniformName)
	v := (int32)(ret)
	return v
}

// SetShaderValue - Set shader uniform value (float)
func SetShaderValue(shader Shader, locIndex int32, value []float32, uniformType ShaderUniformDataType) {
	cshader := shader.cptr()
	clocIndex := (C.int)(locIndex)
	cvalue := (*C.float)(unsafe.Pointer(&value[0]))
	cuniformType := (C.int)(uniformType)
	C.SetShaderValue(*cshader, clocIndex, unsafe.Pointer(cvalue), cuniformType)
}

// SetShaderValueV - Set shader uniform value (float)
func SetShaderValueV(shader Shader, locIndex int32, value []float32, uniformType ShaderUniformDataType, count int32) {
	cshader := shader.cptr()
	clocIndex := (C.int)(locIndex)
	cvalue := (*C.float)(unsafe.Pointer(&value[0]))
	cuniformType := (C.int)(uniformType)
	ccount := (C.int)(count)
	C.SetShaderValueV(*cshader, clocIndex, unsafe.Pointer(cvalue), cuniformType, ccount)
}

// SetShaderValueMatrix - Set shader uniform value (matrix 4x4)
func SetShaderValueMatrix(shader Shader, locIndex int32, mat Matrix) {
	cshader := shader.cptr()
	clocIndex := (C.int)(locIndex)
	cmat := mat.cptr()
	C.SetShaderValueMatrix(*cshader, clocIndex, *cmat)
}

// SetShaderValueTexture - Set shader uniform value for texture (sampler2d)
func SetShaderValueTexture(shader Shader, locIndex int32, texture *Texture2D) {
	cshader := shader.cptr()
	clocIndex := (C.int)(locIndex)
	ctexture := texture.cptr()
	C.SetShaderValueTexture(*cshader, clocIndex, *ctexture)
}

// UnloadShader - Unload a custom shader from memory
func UnloadShader(shader *Shader) {
	cshader := shader.cptr()
	C.UnloadShader(cshader)
}

// GetMouseRay - Get a ray trace from mouse position
//
// Deprecated: Use [GetScreenToWorldRay] instead.
func GetMouseRay(mousePosition vector2.Float32, camera Camera) Ray {
	return GetScreenToWorldRay(mousePosition, camera)
}

// GetScreenToWorldRay - Get a ray trace from screen position (i.e mouse)
func GetScreenToWorldRay(position vector2.Float32, camera Camera) Ray {
	cposition := cvec2ptr(&position)
	ccamera := camera.cptr()
	ret := C.GetScreenToWorldRay(*cposition, *ccamera)
	return *newRayFromPointer(&ret)
}

// GetScreenToWorldRayEx - Get a ray trace from screen position (i.e mouse) in a viewport
func GetScreenToWorldRayEx(position vector2.Float32, camera Camera, width, height int32) Ray {
	cposition := cvec2ptr(&position)
	ccamera := camera.cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.GetScreenToWorldRayEx(*cposition, *ccamera, cwidth, cheight)
	return *newRayFromPointer(&ret)
}

// GetCameraMatrix - Returns camera transform matrix (view matrix)
func GetCameraMatrix(camera Camera) Matrix {
	ccamera := camera.cptr()
	ret := C.GetCameraMatrix(*ccamera)
	return *newMatrixFromPointer(&ret)
}

// GetCameraMatrix2D - Returns camera 2d transform matrix
func GetCameraMatrix2D(camera Camera2D) Matrix {
	ccamera := camera.cptr()
	ret := C.GetCameraMatrix2D(*ccamera)
	return *newMatrixFromPointer(&ret)
}

// GetWorldToScreen - Returns the screen space position from a 3d world space position
func GetWorldToScreen(position vector3.Float32, camera Camera) vector2.Float32 {
	cposition := cvec3ptr(&position)
	ccamera := camera.cptr()
	ret := C.GetWorldToScreen(*cposition, *ccamera)
	return *govec2ptr(&ret)
}

// GetScreenToWorld2D - Returns the world space position for a 2d camera screen space position
func GetScreenToWorld2D(position vector2.Float32, camera Camera2D) vector2.Float32 {
	cposition := cvec2ptr(&position)
	ccamera := camera.cptr()
	ret := C.GetScreenToWorld2D(*cposition, *ccamera)
	return *govec2ptr(&ret)
}

// GetWorldToScreenEx - Get size position for a 3d world space position
func GetWorldToScreenEx(position vector3.Float32, camera Camera, width int32, height int32) vector2.Float32 {
	cposition := cvec3ptr(&position)
	ccamera := camera.cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.GetWorldToScreenEx(*cposition, *ccamera, cwidth, cheight)
	return *govec2ptr(&ret)
}

// GetWorldToScreen2D - Returns the screen space position for a 2d camera world space position
func GetWorldToScreen2D(position vector2.Float32, camera Camera2D) vector2.Float32 {
	cposition := cvec2ptr(&position)
	ccamera := camera.cptr()
	ret := C.GetWorldToScreen2D(*cposition, *ccamera)
	return *govec2ptr(&ret)
}

// SetTargetFPS - Set target FPS (maximum)
func SetTargetFPS[T constraints.Integer](fps T) {
	cfps := (C.int)(fps)
	C.SetTargetFPS(cfps)
}

// GetFPS - Returns current FPS
func GetFPS() int32 {
	ret := C.GetFPS()
	v := (int32)(ret)
	return v
}

// GetFrameTime - Returns time in seconds for one frame
func GetFrameTime() float32 {
	ret := C.GetFrameTime()
	v := (float32)(ret)
	return v
}

// GetTime - Return time in seconds
func GetTime() float64 {
	ret := C.GetTime()
	v := (float64)(ret)
	return v
}

// Custom frame control functions
// NOTE: SwapScreenBuffer and PollInputEvents are intended for advanced users that want full control over the frame processing
// By default EndDrawing() does this job: draws everything + SwapScreenBuffer() + manage frame timing + PollInputEvents()
// To avoid that behaviour and control frame processes manually you can either enable in config.h: SUPPORT_CUSTOM_FRAME_CONTROL
// or add CGO_CFLAGS="-DSUPPORT_CUSTOM_FRAME_CONTROL=1" to your build

// SwapScreenBuffer - Swap back buffer to front buffer
func SwapScreenBuffer() {
	C.SwapScreenBuffer()
}

// Register all input events
func PollInputEvents() {
	C.PollInputEvents()
}

// WaitTime - Wait for some time (halt program execution)
func WaitTime(seconds float64) {
	cseconds := (C.double)(seconds)
	C.WaitTime(cseconds)
}

// Fade - Returns color with alpha applied, alpha goes from 0.0f to 1.0f
func Fade(col colorex.RGBA, alpha float32) colorex.RGBA {
	ccolor := ccolorptr(&col)
	calpha := (C.float)(alpha)
	ret := C.Fade(*ccolor, calpha)
	return *gocolorptr(&ret)
}

// ColorToInt - Get hexadecimal value for a Color (0xRRGGBBAA)
func ColorToInt(col colorex.RGBA) int32 {
	ccolor := ccolorptr(&col)
	ret := C.ColorToInt(*ccolor)
	v := (int32)(ret)
	return v
}

// ColorNormalize - Returns color normalized as float [0..1]
func ColorNormalize(col colorex.RGBA) vector4.Float32 {
	return vector4.NewFloat32(
		float32(col.R)/255,
		float32(col.G)/255,
		float32(col.B)/255,
		float32(col.A)/255)
}

// ColorFromNormalized - Returns Color from normalized values [0..1]
func ColorFromNormalized(normalized vector4.Float32) colorex.RGBA {
	cnormalized := cvec4ptr(&normalized)
	ret := C.ColorFromNormalized(*cnormalized)
	return *gocolorptr(&ret)
}

// ColorToHSV - Returns HSV values for a Color, hue [0..360], saturation/value [0..1]
func ColorToHSV(col colorex.RGBA) vector3.Float32 {
	ccolor := ccolorptr(&col)
	ret := C.ColorToHSV(*ccolor)
	return *govec3ptr(&ret)
}

// ColorFromHSV - Returns a Color from HSV values, hue [0..360], saturation/value [0..1]
func ColorFromHSV(hue, saturation, value float32) colorex.RGBA {
	chue := (C.float)(hue)
	csaturation := (C.float)(saturation)
	cvalue := (C.float)(value)
	ret := C.ColorFromHSV(chue, csaturation, cvalue)
	return *gocolorptr(&ret)
}

// ColorTint - Get color multiplied with another color
func ColorTint(col colorex.RGBA, tint colorex.RGBA) colorex.RGBA {
	ccolor := ccolorptr(&col)
	ctint := ccolorptr(&tint)
	ret := C.ColorTint(*ccolor, *ctint)
	return *gocolorptr(&ret)
}

// ColorBrightness - Get color with brightness correction, brightness factor goes from -1.0f to 1.0f
func ColorBrightness(col colorex.RGBA, factor float32) colorex.RGBA {
	ccolor := ccolorptr(&col)
	cfactor := C.float(factor)
	ret := C.ColorBrightness(*ccolor, cfactor)
	return *gocolorptr(&ret)
}

// ColorContrast - Get color with contrast correction, contrast values between -1.0f and 1.0f
func ColorContrast(col colorex.RGBA, contrast float32) colorex.RGBA {
	ccolor := ccolorptr(&col)
	ccontrast := C.float(contrast)
	ret := C.ColorContrast(*ccolor, ccontrast)
	return *gocolorptr(&ret)
}

// ColorAlpha - Returns color with alpha applied, alpha goes from 0.0f to 1.0f
func ColorAlpha(col colorex.RGBA, alpha float32) colorex.RGBA {
	return Fade(col, alpha)
}

// ColorAlphaBlend - Returns src alpha-blended into dst color with tint
func ColorAlphaBlend(src, dst, tint colorex.RGBA) colorex.RGBA {
	csrc := ccolorptr(&src)
	cdst := ccolorptr(&dst)
	ctint := ccolorptr(&tint)
	ret := C.ColorAlphaBlend(*csrc, *cdst, *ctint)
	return *gocolorptr(&ret)
}

// ColorLerp - Get color lerp interpolation between two colors, factor [0.0f..1.0f]
func ColorLerp(col1, col2 colorex.RGBA, factor float32) colorex.RGBA {
	ccol1 := ccolorptr(&col1)
	ccol2 := ccolorptr(&col2)
	ret := C.ColorLerp(*ccol1, *ccol2, C.float(factor))
	return *gocolorptr(&ret)
}

// GetColor - Returns a Color struct from hexadecimal value
func GetColor(hexValue uint) colorex.RGBA {
	chexValue := (C.uint)(hexValue)
	ret := C.GetColor(chexValue)
	return *gocolorptr(&ret)
}

// GetPixelDataSize - Get pixel data size in bytes for certain format
func GetPixelDataSize(width, height, format int32) int32 {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cformat := (C.int)(format)
	ret := C.GetPixelDataSize(cwidth, cheight, cformat)
	v := (int32)(ret)
	return v
}

// Vector3ToFloat - Converts Vector3 to float32 slice
func Vector3ToFloat(vec vector3.Float32) []float32 {
	data := make([]float32, 0)
	data[0] = vec.X
	data[1] = vec.Y
	data[2] = vec.Z

	return data
}

// GetRandomValue - Returns a random value between min and max (both included)
func GetRandomValue(min, max int32) int32 {
	cmin := (C.int)(min)
	cmax := (C.int)(max)
	ret := C.GetRandomValue(cmin, cmax)
	v := (int32)(ret)
	return v
}

// OpenURL - Open URL with default system browser (if available)
func OpenURL(url string) {
	curl := textAlloc(url)
	C.OpenURL(curl)
}

// SetConfigFlags - Setup some window configuration flags
func SetConfigFlags(flags ConfigFlags) {
	cflags := (C.uint)(flags)
	C.SetConfigFlags(cflags)
}

// TakeScreenshot - Takes a screenshot of current screen (saved a .png)
func TakeScreenshot(name string) {
	cname := textAlloc(name)
	C.TakeScreenshot(cname)
}

// LoadAutomationEventList - Load automation events list from file, NULL for empty list, capacity = MAX_AUTOMATION_EVENTS
func LoadAutomationEventList(fileName string) AutomationEventList {
	cfileName := textAlloc(fileName)

	ret := C.LoadAutomationEventList(cfileName)
	return *newAutomationEventListFromPointer(&ret)
}

// UnloadAutomationEventList - Unload automation events list from file
func UnloadAutomationEventList(list *AutomationEventList) {
	C.UnloadAutomationEventList((*C.AutomationEventList)(list))
}

// ExportAutomationEventList - Export automation events list as text file
func ExportAutomationEventList(list AutomationEventList, fileName string) bool {
	cfileName := textAlloc(fileName)

	ret := C.ExportAutomationEventList(*list.cptr(), cfileName)
	v := bool(ret)

	return v
}

// SetAutomationEventList - Set automation event list to record to
func SetAutomationEventList(list *AutomationEventList) {
	C.SetAutomationEventList(list.cptr())
}

// SetAutomationEventBaseFrame - Set automation event internal base frame to start recording
func SetAutomationEventBaseFrame(frame int) {
	cframe := (C.int)(frame)
	C.SetAutomationEventBaseFrame(cframe)
}

// StartAutomationEventRecording - Start recording automation events (AutomationEventList must be set)
func StartAutomationEventRecording() {
	C.StartAutomationEventRecording()
}

// StopAutomationEventRecording - Stop recording automation events
func StopAutomationEventRecording() {
	C.StopAutomationEventRecording()
}

// PlayAutomationEvent - Play a recorded automation event
func PlayAutomationEvent(event AutomationEvent) {
	C.PlayAutomationEvent(*event.cptr())
}

// IsKeyPressed - Detect if a key has been pressed once
func IsKeyPressed(key KeyType) bool {
	ckey := (C.int)(key)
	ret := C.IsKeyPressed(ckey)
	v := bool(ret)
	return v
}

// IsKeyPressedRepeat - Detect if a key has been pressed again (Only PLATFORM_DESKTOP)
func IsKeyPressedRepeat(key KeyType) bool {
	ckey := (C.int)(key)
	ret := C.IsKeyPressedRepeat(ckey)
	v := bool(ret)
	return v
}

// IsKeyDown - Detect if a key is being pressed
func IsKeyDown(key KeyType) bool {
	ckey := (C.int)(key)
	ret := C.IsKeyDown(ckey)
	v := bool(ret)
	return v
}

// IsKeyReleased - Detect if a key has been released once
func IsKeyReleased(key KeyType) bool {
	ckey := (C.int)(key)
	ret := C.IsKeyReleased(ckey)
	v := bool(ret)
	return v
}

// IsKeyUp - Detect if a key is NOT being pressed
func IsKeyUp(key KeyType) bool {
	ckey := (C.int)(key)
	ret := C.IsKeyUp(ckey)
	v := bool(ret)
	return v
}

func GetKeyDownCount() int {
	ret := C.GetKeyDownCount()
	v := (int)(ret)
	return v
}

// GetKeyPressed - Get latest key pressed
func GetKeyPressed() KeyType {
	ret := C.GetKeyPressed()
	v := (KeyType)(ret)
	return v
}

// GetCharPressed - Get the last char pressed
func GetCharPressed() KeyType {
	ret := C.GetCharPressed()
	v := (KeyType)(ret)
	return v
}

// SetExitKey - Set a custom key to exit program (default is ESC)
func SetExitKey(key KeyType) {
	ckey := (C.int)(key)
	C.SetExitKey(ckey)
}

// IsGamepadAvailable - Detect if a gamepad is available
func IsGamepadAvailable[GT constraints.Integer](gamepad GT) bool {
	cgamepad := (C.int)(gamepad)
	ret := C.IsGamepadAvailable(cgamepad)
	v := bool(ret)
	return v
}

// GetGamepadName - Return gamepad internal name id
func GetGamepadName[GT constraints.Integer](gamepad GT) string {
	cgamepad := (C.int)(gamepad)
	ret := C.GetGamepadName(cgamepad)
	v := C.GoString(ret)
	return v
}

// IsGamepadButtonPressed - Detect if a gamepad button has been pressed once
func IsGamepadButtonPressed[GT constraints.Integer](gamepad GT, button GamepadButtonType) bool {
	cgamepad := (C.int)(gamepad)
	cbutton := (C.int)(button)
	ret := C.IsGamepadButtonPressed(cgamepad, cbutton)
	v := bool(ret)
	return v
}

// IsGamepadButtonDown - Detect if a gamepad button is being pressed
func IsGamepadButtonDown[GT constraints.Integer](gamepad GT, button GamepadButtonType) bool {
	cgamepad := (C.int)(gamepad)
	cbutton := (C.int)(button)
	ret := C.IsGamepadButtonDown(cgamepad, cbutton)
	v := bool(ret)
	return v
}

// IsGamepadButtonReleased - Detect if a gamepad button has been released once
func IsGamepadButtonReleased[GT constraints.Integer](gamepad GT, button GamepadButtonType) bool {
	cgamepad := (C.int)(gamepad)
	cbutton := (C.int)(button)
	ret := C.IsGamepadButtonReleased(cgamepad, cbutton)
	v := bool(ret)
	return v
}

// IsGamepadButtonUp - Detect if a gamepad button is NOT being pressed
func IsGamepadButtonUp[GT constraints.Integer](gamepad GT, button GamepadButtonType) bool {
	cgamepad := (C.int)(gamepad)
	cbutton := (C.int)(button)
	ret := C.IsGamepadButtonUp(cgamepad, cbutton)
	v := bool(ret)
	return v
}

// GetGamepadButtonPressed - Get the last gamepad button pressed
func GetGamepadButtonPressed() int32 {
	ret := C.GetGamepadButtonPressed()
	v := (int32)(ret)
	return v
}

// GetGamepadAxisCount - Return gamepad axis count for a gamepad
func GetGamepadAxisCount[GT constraints.Integer](gamepad GT) int32 {
	cgamepad := (C.int)(gamepad)
	ret := C.GetGamepadAxisCount(cgamepad)
	v := (int32)(ret)
	return v
}

// GetGamepadAxisMovement - Return axis movement value for a gamepad axis
func GetGamepadAxisMovement[GT constraints.Integer](gamepad GT, axis GamepadAxisType) float32 {
	cgamepad := (C.int)(gamepad)
	caxis := (C.int)(axis)
	ret := C.GetGamepadAxisMovement(cgamepad, caxis)
	v := (float32)(ret)
	return v
}

// SetGamepadMappings - Set internal gamepad mappings (SDL_GameControllerDB)
func SetGamepadMappings(mappings string) int32 {
	cmappings := textAlloc(mappings)
	//defer C.free(unsafe.Pointer(cmappings))  TODO: possible mappings truncation
	ret := C.SetGamepadMappings(cmappings)
	v := (int32)(ret)
	return v
}

// SetGamepadVibration - Set gamepad vibration for both motors (duration in seconds)
func SetGamepadVibration(gamepad int32, leftMotor, rightMotor, duration float32) {
	C.SetGamepadVibration(C.int(gamepad), C.float(leftMotor), C.float(rightMotor), C.float(duration))
}

// IsMouseButtonPressed - Detect if a mouse button has been pressed once
func IsMouseButtonPressed(button MouseButtonType) bool {
	cbutton := (C.int)(button)
	ret := C.IsMouseButtonPressed(cbutton)
	v := bool(ret)
	return v
}

// IsMouseButtonDown - Detect if a mouse button is being pressed
func IsMouseButtonDown(button MouseButtonType) bool {
	cbutton := (C.int)(button)
	ret := C.IsMouseButtonDown(cbutton)
	v := bool(ret)
	return v
}

// IsMouseButtonReleased - Detect if a mouse button has been released once
func IsMouseButtonReleased(button MouseButtonType) bool {
	cbutton := (C.int)(button)
	ret := C.IsMouseButtonReleased(cbutton)
	v := bool(ret)
	return v
}

// IsMouseButtonUp - Detect if a mouse button is NOT being pressed
func IsMouseButtonUp(button MouseButtonType) bool {
	cbutton := (C.int)(button)
	ret := C.IsMouseButtonUp(cbutton)
	v := bool(ret)
	return v
}

// GetMouseX - Returns mouse position X
func GetMouseX() int32 {
	ret := C.GetMouseX()
	v := (int32)(ret)
	return v
}

// GetMouseY - Returns mouse position Y
func GetMouseY() int32 {
	ret := C.GetMouseY()
	v := (int32)(ret)
	return v
}

// GetMousePosition - Returns mouse position XY
func GetMousePosition() vector2.Float32 {
	ret := C.GetMousePosition()
	return *govec2ptr(&ret)
}

// GetMouseDelta - Get mouse delta between frames
func GetMouseDelta() vector2.Float32 {
	ret := C.GetMouseDelta()
	return *govec2ptr(&ret)
}

// SetMousePosition - Set mouse position XY
func SetMousePosition(x, y int) {
	cx := (C.int)(x)
	cy := (C.int)(y)
	C.SetMousePosition(cx, cy)
}

// SetMouseOffset - Set mouse offset
func SetMouseOffset(offsetX, offsetY int) {
	ox := (C.int)(offsetX)
	oy := (C.int)(offsetY)
	C.SetMouseOffset(ox, oy)
}

// SetMouseScale - Set mouse scaling
func SetMouseScale(scaleX, scaleY float32) {
	cscaleX := (C.float)(scaleX)
	cscaleY := (C.float)(scaleY)
	C.SetMouseScale(cscaleX, cscaleY)
}

// GetMouseWheelMove - Get mouse wheel movement for X or Y, whichever is larger
func GetMouseWheelMove() float32 {
	ret := C.GetMouseWheelMove()
	v := (float32)(ret)
	return v
}

// GetMouseWheelMoveV - Get mouse wheel movement for both X and Y
func GetMouseWheelMoveV() vector2.Float32 {
	ret := C.GetMouseWheelMoveV()
	return *govec2ptr(&ret)
}

// SetMouseCursor - Set mouse cursor
func SetMouseCursor(cursor MouseCursorType) {
	ccursor := (C.int)(cursor)
	C.SetMouseCursor(ccursor)
}

// GetTouchX - Returns touch position X for touch point 0 (relative to screen size)
func GetTouchX() int32 {
	ret := C.GetTouchX()
	v := (int32)(ret)
	return v
}

// GetTouchY - Returns touch position Y for touch point 0 (relative to screen size)
func GetTouchY() int32 {
	ret := C.GetTouchY()
	v := (int32)(ret)
	return v
}

// GetTouchPosition - Returns touch position XY for a touch point index (relative to screen size)
func GetTouchPosition[IT constraints.Integer](index IT) vector2.Float32 {
	cindex := (C.int)(index)
	ret := C.GetTouchPosition(cindex)
	return *govec2ptr(&ret)
}

// GetTouchPointId - Get touch point identifier for given index
func GetTouchPointId[IT constraints.Integer](index IT) int32 {
	cindex := (C.int)(index)
	ret := C.GetTouchPointId(cindex)
	v := (int32)(ret)
	return v
}

// GetTouchPointCount - Get number of touch points
func GetTouchPointCount() int32 {
	ret := C.GetTouchPointCount()
	v := (int32)(ret)
	return v
}

// BeginVrStereoMode - Begin stereo rendering (requires VR simulator)
func BeginVrStereoMode(config VrStereoConfig) {
	C.BeginVrStereoMode(*(*C.VrStereoConfig)(unsafe.Pointer(&config)))
}

// EndVrStereoMode - End stereo rendering (requires VR simulator)
func EndVrStereoMode() {
	C.EndVrStereoMode()
}

// LoadVrStereoConfig - Load VR stereo config for VR simulator device parameters
func LoadVrStereoConfig(device VrDeviceInfo) VrStereoConfig {
	ret := C.LoadVrStereoConfig(*(*C.VrDeviceInfo)(unsafe.Pointer(&device)))
	return *(*VrStereoConfig)(unsafe.Pointer(&ret))
}

// UnloadVrStereoConfig - Unload VR stereo config
func UnloadVrStereoConfig(config VrStereoConfig) {
	C.UnloadVrStereoConfig(*(*C.VrStereoConfig)(unsafe.Pointer(&config)))
}
