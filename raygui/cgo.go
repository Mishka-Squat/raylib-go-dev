package raygui

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/raygui/src -std=gnu99 -Wno-unused-result

#include "raygui.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
)

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

/*
// govec4ptr - Returns new Vector4 from pointer
func govec4ptr(v *C.Vector4) *rl.Vector4 {
	return (*rl.Vector4)(unsafe.Pointer(v))
}

// cvec4 returns C pointer
func cvec4ptr(v *rl.Vector4) *C.Vector4 {
	return (*C.Vector4)(unsafe.Pointer(v))
}
*/

// gorec2ptr - Returns new Rectangle from pointer
func gorec2ptr(ptr *C.Rectangle) *rect2.Float32 {
	return (*rect2.Float32)(unsafe.Pointer(ptr))
}

// cptr returns C pointer
func crect2ptr(r *rect2.Float32) *C.Rectangle {
	return (*C.Rectangle)(unsafe.Pointer(r))
}

func textAlloc(text string) *C.char {
	ctext := (*C.char)(unsafe.Pointer(unsafe.StringData(text)))
	clen := (C.int)(len(text))
	return C.TextAlloc(ctext, &clen)
}
