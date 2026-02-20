package rl

/*
#cgo nocallback DrawFPS
#cgo nocallback DrawText
#cgo nocallback DrawTextCodepoint
#cgo nocallback DrawTextCodepoints
#cgo nocallback DrawTextEx
#cgo nocallback DrawTextPro
#cgo nocallback GenImageFontAtlas
#cgo nocallback GetFontDefault
#cgo nocallback GetGlyphAtlasRec
#cgo nocallback GetGlyphIndex
#cgo nocallback GetGlyphInfo
#cgo nocallback IsFontValid
#cgo nocallback MeasureText
#cgo nocallback MeasureTextEx
#cgo nocallback SetTextLineSpacing

#cgo noescape DrawFPS
#cgo noescape DrawText
#cgo noescape DrawTextCodepoint
#cgo noescape DrawTextCodepoints
#cgo noescape DrawTextEx
#cgo noescape DrawTextPro
#cgo noescape GenImageFontAtlas
#cgo noescape GetFontDefault
#cgo noescape GetGlyphAtlasRec
#cgo noescape GetGlyphIndex
#cgo noescape GetGlyphInfo
#cgo noescape IsFontValid
#cgo noescape LoadFont
#cgo noescape LoadFontData
#cgo noescape LoadFontEx
#cgo noescape LoadFontFromImage
#cgo noescape LoadFontFromMemory
#cgo noescape MeasureText
#cgo noescape MeasureTextEx
#cgo noescape SetTextLineSpacing
#cgo noescape UnloadFont
#cgo noescape UnloadFontData

#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	"github.com/Mishka-Squat/goex/mathex"
)

var defaultFont Font

// GetFontDefault - Get the default Font
func GetFontDefault() Font {
	if !defaultFont.IsValid() {
		ret := C.GetFontDefault()
		defaultFont = *newFontFromPointer(&ret)
	}

	return defaultFont
}

// LoadFont - Load a Font image into GPU memory (VRAM)
func LoadFont(fileName string) Font {
	cfileName := textAlloc(fileName)
	ret := C.LoadFont(cfileName)
	return *newFontFromPointer(&ret)
}

// LoadFontEx - Load Font from file with extended parameters
func LoadFontEx(fileName string, fontSize int32, fontChars []rune, runesNumber ...int32) Font {
	var cfontChars *C.int
	var ccharsCount C.int

	cfileName := textAlloc(fileName)
	cfontSize := (C.int)(fontSize)
	if fontChars != nil {
		cfontChars = (*C.int)(unsafe.Pointer(&fontChars[0]))
		ccharsCount = (C.int)(len(fontChars))
	}
	if fontChars == nil {
		if len(runesNumber) > 0 {
			ccharsCount = (C.int)(runesNumber[0])
		}
	}
	ret := C.LoadFontEx(cfileName, cfontSize, cfontChars, ccharsCount)
	return *newFontFromPointer(&ret)
}

// LoadFontFromImage - Loads an Image font file (XNA style)
func LoadFontFromImage(image Image, key colorex.RGBA, firstChar int32) Font {
	cimage := image.cptr()
	ckey := ccolorptr(&key)
	cfirstChar := (C.int)(firstChar)
	ret := C.LoadFontFromImage(*cimage, *ckey, cfirstChar)
	return *newFontFromPointer(&ret)
}

// LoadFontFromMemory - Load font from memory buffer, fileType refers to extension: i.e. ".ttf"
func LoadFontFromMemory(fileType string, fileData []byte, fontSize int32, codepoints []rune) Font {
	cfileType := textAlloc(fileType)
	cfileData := (*C.uchar)(unsafe.Pointer(&fileData[0]))
	cdataSize := (C.int)(len(fileData))
	cfontSize := (C.int)(fontSize)
	cfontChars := (*C.int)(unsafe.SliceData(codepoints))
	ccharsCount := (C.int)(len(codepoints))
	ret := C.LoadFontFromMemory(cfileType, cfileData, cdataSize, cfontSize, cfontChars, ccharsCount)
	return *newFontFromPointer(&ret)
}

// IsFontValid - Check if a font is valid (font data loaded, WARNING: GPU texture not checked)
func IsFontValid(font Font) bool {
	cfont := font.cptr()
	ret := C.IsFontValid(*cfont)
	v := bool(ret)
	return v
}

// LoadFontData - Load font data for further use
func LoadFontData(fileData []byte, fontSize int32, codePoints []rune, codepointCount, typ int32) []GlyphInfo {
	cfileData := (*C.uchar)(unsafe.Pointer(&fileData[0]))
	cdataSize := (C.int)(len(fileData))
	cfontSize := (C.int)(fontSize)
	ccodePoints := (*C.int)(unsafe.SliceData(codePoints))
	// In case no chars count provided, default to 95
	if codepointCount <= 0 {
		codepointCount = 95
	}
	ccodePointCount := (C.int)(codepointCount)
	ctype := (C.int)(typ)
	cglyphCount := (C.int)(codepointCount)
	ret := C.LoadFontData(cfileData, cdataSize, cfontSize, ccodePoints, ccodePointCount, ctype, &cglyphCount)
	v := unsafe.Slice((*GlyphInfo)(unsafe.Pointer(ret)), cglyphCount)
	return v
}

// UnloadFontData - Unload font chars info data (RAM)
func UnloadFontData(glyphs []GlyphInfo) {
	cglyphs := (*C.GlyphInfo)(unsafe.Pointer(&glyphs[0]))
	C.UnloadFontData(cglyphs, C.int(len(glyphs)))
}

// UnloadFont - Unload Font from GPU memory (VRAM)
func UnloadFont(font *Font) {
	cfont := font.cptr()
	C.UnloadFont(cfont)
}

// DrawFPS - Shows current FPS
func DrawFPS[XT, YT mathex.Number](posX XT, posY YT) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	C.DrawFPS(cposX, cposY)
}

// DrawText - Draw text (using default font)
func DrawText[XT, YT mathex.Number](text string, posX XT, posY YT, fontSize int32, col colorex.RGBA) {
	ctext := textAlloc(text)
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	cfontSize := (C.int)(fontSize)
	ccolor := ccolorptr(&col)
	C.DrawText(ctext, cposX, cposY, cfontSize, *ccolor)
}

// DrawTextEx - Draw text using Font and additional parameters
func DrawTextEx(font Font, text string, position vector2.Float32, fontSize float32, spacing float32, tint colorex.RGBA) {
	cfont := font.cptr()
	ctext := textAlloc(text)
	cposition := cvec2ptr(&position)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ctint := ccolorptr(&tint)
	C.DrawTextEx(*cfont, ctext, *cposition, cfontSize, cspacing, *ctint)
}

func DrawTextLayout(font Font, text string, fontSize float32, spacing float32, tint colorex.RGBA, layoutFn func(wh vector2.Float32) rect2.Float32) {
	rect := layoutFn(MeasureTextEx(font, text, fontSize, spacing))
	DrawTextEx(font, text, rect.Position, fontSize, spacing, tint)
}

// DrawTextPro - Draw text using Font and pro parameters (rotation)
func DrawTextPro(font Font, text string, position vector2.Float32, origin vector2.Float32, rotation, fontSize float32, spacing float32, tint colorex.RGBA) {
	cfont := font.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cposition := cvec2ptr(&position)
	crotation := (C.float)(rotation)
	corigin := cvec2ptr(&origin)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ctint := ccolorptr(&tint)
	C.DrawTextPro(*cfont, ctext, *cposition, *corigin, crotation, cfontSize, cspacing, *ctint)
}

// SetTextLineSpacing - Set vertical line spacing when drawing with line-breaks
func SetTextLineSpacing(spacing int) {
	cspacing := (C.int)(spacing)
	C.SetTextLineSpacing(cspacing)
}

// MeasureText - Measure string width for default font
func MeasureText(text string, fontSize int32) int32 {
	ctext := textAlloc(text)
	cfontSize := (C.int)(fontSize)
	ret := C.MeasureText(ctext, cfontSize)
	return (int32)(ret)
}

// MeasureTextEx - Measure string size for Font
func MeasureTextEx(font Font, text string, fontSize float32, spacing float32) vector2.Float32 {
	cfont := font.cptr()
	ctext := textAlloc(text)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ret := C.MeasureTextEx(*cfont, ctext, cfontSize, cspacing)
	return *govec2ptr(&ret)
}

// GetGlyphIndex - Get glyph index position in font for a codepoint (unicode character), fallback to '?' if not found
func GetGlyphIndex(font Font, codepoint int32) int32 {
	cfont := font.cptr()
	ccodepoint := (C.int)(codepoint)
	ret := C.GetGlyphIndex(*cfont, ccodepoint)
	return (int32)(ret)
}

// GetGlyphInfo - Get glyph font info data for a codepoint (unicode character), fallback to '?' if not found
func GetGlyphInfo(font Font, codepoint int32) GlyphInfo {
	cfont := font.cptr()
	ccodepoint := (C.int)(codepoint)
	ret := C.GetGlyphInfo(*cfont, ccodepoint)
	return *newGlyphInfoFromPointer(&ret)
}

// GetGlyphAtlasRec - Get glyph rectangle in font atlas for a codepoint (unicode character), fallback to '?' if not found
func GetGlyphAtlasRec(font Font, codepoint int32) rect2.Float32 {
	cfont := font.cptr()
	ccodepoint := (C.int)(codepoint)
	ret := C.GetGlyphAtlasRec(*cfont, ccodepoint)
	return *gorec2ptr(&ret)
}

// GenImageFontAtlas - Generate image font atlas using chars info
func GenImageFontAtlas(glyphs []GlyphInfo, glyphRecs []*rect2.Float32, fontSize int32, padding int32, packMethod int32) Image {
	cglyphs := (*C.GlyphInfo)(unsafe.Pointer(&glyphs[0]))
	cglyphRecs := (**C.Rectangle)(unsafe.Pointer(&glyphRecs[0]))
	cglyphCount := C.int(len(glyphs))
	ret := C.GenImageFontAtlas(cglyphs, cglyphRecs, cglyphCount, C.int(fontSize), C.int(padding), C.int(packMethod))
	return *newImageFromPointer(&ret)
}

// DrawTextCodepoint - Draw one character (codepoint)
func DrawTextCodepoint(font Font, codepoint rune, position vector2.Float32, fontSize float32, tint colorex.RGBA) {
	cfont := font.cptr()
	ccodepoint := (C.int)(codepoint)
	cposition := cvec2ptr(&position)
	cfontSize := (C.float)(fontSize)
	ctint := ccolorptr(&tint)
	C.DrawTextCodepoint(*cfont, ccodepoint, *cposition, cfontSize, *ctint)
}

// DrawTextCodepoints - Draw multiple character (codepoint)
func DrawTextCodepoints(font Font, codepoints []rune, position vector2.Float32, fontSize float32, spacing float32, tint colorex.RGBA) {
	cfont := font.cptr()
	ccodepoints := (*C.int)(unsafe.SliceData(codepoints))
	ccodepointCount := C.int(len(codepoints))
	cposition := cvec2ptr(&position)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ctint := ccolorptr(&tint)
	C.DrawTextCodepoints(*cfont, ccodepoints, ccodepointCount, *cposition, cfontSize, cspacing, *ctint)
}
