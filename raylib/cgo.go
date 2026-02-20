package rl

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/raylib/src/external/glfw/include -std=gnu99 -Wno-missing-braces -Wno-unused-result -Wno-implicit-function-declaration
#cgo android CFLAGS: -DPLATFORM_ANDROID -DPLATFORM_ANDROID_NOMAIN -DGRAPHICS_API_OPENGL_ES2

#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	"github.com/Mishka-Squat/gamemath/vector4"
	"github.com/Mishka-Squat/goex/image/colorex"
)

// core

// govec2ptr - Returns new Vector2 from pointer
func govec2ptr(v *C.Vector2) *vector2.Float32 {
	return (*vector2.Float32)(unsafe.Pointer(v))
}

// cptr returns C pointer
func cvec2ptr(v *vector2.Float32) *C.Vector2 {
	return (*C.Vector2)(unsafe.Pointer(v))
}

// govec3ptr - Returns new Vector3 from pointer
func govec3ptr(v *C.Vector3) *vector3.Float32 {
	return (*vector3.Float32)(unsafe.Pointer(v))
}

// cvec3 returns C pointer
func cvec3ptr(v *vector3.Float32) *C.Vector3 {
	return (*C.Vector3)(unsafe.Pointer(v))
}

// govec4ptr - Returns new Vector4 from pointer
func govec4ptr(v *C.Vector4) *vector4.Float32 {
	return (*vector4.Float32)(unsafe.Pointer(v))
}

// cvec4 returns C pointer
func cvec4ptr(v *vector4.Float32) *C.Vector4 {
	return (*C.Vector4)(unsafe.Pointer(v))
}

// newMatrixFromPointer - Returns new Matrix from pointer
func newMatrixFromPointer(ptr *C.Matrix) *Matrix {
	return (*Matrix)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (m *Matrix) cptr() *C.Matrix {
	return (*C.Matrix)(unsafe.Pointer(m))
}

// gocolorptr - Returns new Color from pointer
func gocolorptr(ptr *C.Color) *colorex.RGBA {
	return (*colorex.RGBA)(unsafe.Pointer(ptr))
}

// ccolorptr returns color C pointer
func ccolorptr(col *colorex.RGBA) *C.Color {
	return (*C.Color)(unsafe.Pointer(col))
}

// gorec2ptr - Returns new Rectangle from pointer
func gorec2ptr(ptr *C.Rectangle) *rect2.Float32 {
	return (*rect2.Float32)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func crect2ptr(r *rect2.Float32) *C.Rectangle {
	return (*C.Rectangle)(unsafe.Pointer(r))
}

// newCamera3DFromPointer - Returns new Camera3D from pointer
func newCamera3DFromPointer(ptr unsafe.Pointer) Camera3D {
	return *(*Camera3D)(ptr)
}

// cptr returns C pointer
func (c *Camera) cptr() *C.Camera {
	return (*C.Camera)(unsafe.Pointer(c))
}

// newCamera2DFromPointer - Returns new Camera2D from pointer
func newCamera2DFromPointer(ptr unsafe.Pointer) Camera2D {
	return *(*Camera2D)(ptr)
}

// cptr returns C pointer
func (c *Camera2D) cptr() *C.Camera2D {
	return (*C.Camera2D)(unsafe.Pointer(c))
}

// newBoundingBoxFromPointer - Returns new BoundingBox from pointer
func newBoundingBoxFromPointer(ptr *C.BoundingBox) *BoundingBox {
	return (*BoundingBox)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (b *BoundingBox) cptr() *C.BoundingBox {
	return (*C.BoundingBox)(unsafe.Pointer(b))
}

// newShaderFromPointer - Returns new Shader from pointer
func newShaderFromPointer(ptr *C.Shader) *Shader {
	return (*Shader)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (s *Shader) cptr() *C.Shader {
	return (*C.Shader)(unsafe.Pointer(s))
}

// newAutomationEventFromPointer - Returns new AutomationEvent from pointer
func newAutomationEventFromPointer(ptr unsafe.Pointer) AutomationEvent {
	return *(*AutomationEvent)(ptr)
}

// cptr returns C pointer
func (a *AutomationEvent) cptr() *C.AutomationEvent {
	return (*C.AutomationEvent)(unsafe.Pointer(a))
}

// newAutomationEventListFromPointer - Returns new AutomationEventList from pointer
func newAutomationEventListFromPointer(ptr *C.AutomationEventList) *AutomationEventList {
	return (*AutomationEventList)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (a *AutomationEventList) cptr() *C.AutomationEventList {
	return (*C.AutomationEventList)(unsafe.Pointer(a))
}

// model

// newMeshFromPointer - Returns new Mesh from pointer
func newMeshFromPointer(ptr *C.Mesh) *Mesh {
	return (*Mesh)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (m *Mesh) cptr() *C.Mesh {
	return (*C.Mesh)(unsafe.Pointer(m))
}

// newMaterialFromPointer - Returns new Material from pointer
func newMaterialFromPointer(ptr *C.Material) *Material {
	return (*Material)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (m *Material) cptr() *C.Material {
	return (*C.Material)(unsafe.Pointer(m))
}

// newModelFromPointer - Returns new Model from pointer
func newModelFromPointer(ptr *C.Model) *Model {
	return (*Model)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (m *Model) cptr() *C.Model {
	return (*C.Model)(unsafe.Pointer(m))
}

// newRayFromPointer - Returns new Ray from pointer
func newRayFromPointer(ptr *C.Ray) *Ray {
	return (*Ray)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (r *Ray) cptr() *C.Ray {
	return (*C.Ray)(unsafe.Pointer(r))
}

// newModelAnimationFromPointer - Returns new ModelAnimation from pointer
func newModelAnimationFromPointer(ptr *C.ModelAnimation) *ModelAnimation {
	return (*ModelAnimation)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (r *ModelAnimation) cptr() *C.ModelAnimation {
	return (*C.ModelAnimation)(unsafe.Pointer(r))
}

// newRayCollisionFromPointer - Returns new RayCollision from pointer
func newRayCollisionFromPointer(ptr *C.RayCollision) *RayCollision {
	return (*RayCollision)(unsafe.Pointer(ptr))
}

// texture

// newImageFromPointer - Returns new Image from pointer
func newImageFromPointer(ptr *C.Image) *Image {
	return (*Image)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (i *Image) cptr() *C.Image {
	return (*C.Image)(unsafe.Pointer(i))
}

// newTexture2DFromPointer - Returns new Texture2D from pointer
func newTexture2DFromPointer(ptr *C.Texture2D) *Texture2D {
	return (*Texture2D)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (t *Texture2D) cptr() *C.Texture2D {
	return (*C.Texture2D)(unsafe.Pointer(t))
}

// newRenderTexture2DFromPointer - Returns new RenderTexture2D from pointer
func newRenderTexture2DFromPointer(ptr *C.RenderTexture2D) *RenderTexture2D {
	return (*RenderTexture2D)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (r *RenderTexture2D) cptr() *C.RenderTexture2D {
	return (*C.RenderTexture2D)(unsafe.Pointer(r))
}

// cptr returns C pointer
func (n *NPatchInfo) cptr() *C.NPatchInfo {
	return (*C.NPatchInfo)(unsafe.Pointer(n))
}

// text

// newGlyphInfoFromPointer - Returns new GlyphInfo from pointer
func newGlyphInfoFromPointer(ptr *C.GlyphInfo) *GlyphInfo {
	return (*GlyphInfo)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (c *GlyphInfo) cptr() *C.GlyphInfo {
	return (*C.GlyphInfo)(unsafe.Pointer(c))
}

// newFontFromPointer - Returns new Font from pointer
func newFontFromPointer(ptr *C.Font) *Font {
	return (*Font)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (s *Font) cptr() *C.Font {
	return (*C.Font)(unsafe.Pointer(s))
}

// audio

// newWaveFromPointer - Returns new Wave from pointer
func newWaveFromPointer(ptr *C.Wave) *Wave {
	return (*Wave)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (w *Wave) cptr() *C.Wave {
	return (*C.Wave)(unsafe.Pointer(w))
}

// newSoundFromPointer - Returns new Sound from pointer
func newSoundFromPointer(ptr *C.Sound) *Sound {
	return (*Sound)(unsafe.Pointer(ptr))
}

func (s *Sound) cptr() *C.Sound {
	return (*C.Sound)(unsafe.Pointer(s))
}

// newAudioStreamFromPointer - Returns new AudioStream from pointer
func newAudioStreamFromPointer(ptr *C.AudioStream) *AudioStream {
	return (*AudioStream)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func (a *AudioStream) cptr() *C.AudioStream {
	return (*C.AudioStream)(unsafe.Pointer(a))
}

// newMusicFromPointer - Returns new Music from pointer
func newMusicFromPointer(ptr *C.Music) *Music {
	return (*Music)(unsafe.Pointer(ptr))
}

func (s *Music) cptr() *C.Music {
	return (*C.Music)(unsafe.Pointer(s))
}

func textAlloc(text string) *C.char {
	ctext := (*C.char)(unsafe.Pointer(unsafe.StringData(text)))
	clen := (C.int)(len(text))
	return C.TextAlloc(ctext, (*C.int)(unsafe.Pointer(&clen)))
}
