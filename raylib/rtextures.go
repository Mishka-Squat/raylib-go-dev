package rl

/*
// nocallback is wierd - looks like it means function have no returns.
#cgo nocallback DrawTexture
#cgo nocallback DrawTextureEx
#cgo nocallback DrawTextureNPatch
#cgo nocallback DrawTexturePro
#cgo nocallback DrawTextureRec
#cgo nocallback DrawTextureTiled
#cgo nocallback DrawTextureV
#cgo nocallback ExportImage
#cgo nocallback ExportImageToMemory
#cgo nocallback GenImageCellular
#cgo nocallback GenImageChecked
#cgo nocallback GenImageColor
#cgo nocallback GenImageGradientLinear
#cgo nocallback GenImageGradientRadial
#cgo nocallback GenImageGradientSquare
#cgo nocallback GenImagePerlinNoise
#cgo nocallback GenImageText
#cgo nocallback GenImageWhiteNoise
#cgo nocallback GenTextureMipmaps
#cgo nocallback GetImageColor
#cgo nocallback ImageAlphaClear
#cgo nocallback ImageAlphaCrop
#cgo nocallback ImageAlphaMask
#cgo nocallback ImageAlphaPremultiply
#cgo nocallback ImageBlurGaussian
#cgo nocallback ImageClearBackground
#cgo nocallback ImageColorBrightness
#cgo nocallback ImageColorContrast
#cgo nocallback ImageColorGrayscale
#cgo nocallback ImageColorInvert
#cgo nocallback ImageColorReplace
#cgo nocallback ImageColorTint
#cgo nocallback ImageCopy
#cgo nocallback ImageCrop
#cgo nocallback ImageDither
#cgo nocallback ImageDraw
#cgo nocallback ImageDrawCircle
#cgo nocallback ImageDrawCircleLines
#cgo nocallback ImageDrawCircleLinesV
#cgo nocallback ImageDrawCircleV
#cgo nocallback ImageDrawLine
#cgo nocallback ImageDrawLineEx
#cgo nocallback ImageDrawLineV
#cgo nocallback ImageDrawPixel
#cgo nocallback ImageDrawPixelV
#cgo nocallback ImageDrawRectangle
#cgo nocallback ImageDrawRectangleLines
#cgo nocallback ImageDrawRectangleRec
#cgo nocallback ImageDrawRectangleV
#cgo nocallback ImageDrawText
#cgo nocallback ImageDrawTextEx
#cgo nocallback ImageDrawTriangle
#cgo nocallback ImageDrawTriangleEx
#cgo nocallback ImageDrawTriangleFan
#cgo nocallback ImageDrawTriangleLines
#cgo nocallback ImageDrawTriangleStrip
#cgo nocallback ImageFlipHorizontal
#cgo nocallback ImageFlipVertical
#cgo nocallback ImageFormat
#cgo nocallback ImageFromChannel
#cgo nocallback ImageFromImage
#cgo nocallback ImageKernelConvolution
#cgo nocallback ImageMipmaps
#cgo nocallback ImageResize
#cgo nocallback ImageResizeCanvas
#cgo nocallback ImageResizeNN
#cgo nocallback ImageRotate
#cgo nocallback ImageRotateCCW
#cgo nocallback ImageRotateCW
#cgo nocallback ImageText
#cgo nocallback ImageTextEx
#cgo nocallback ImageToPOT
#cgo nocallback IsImageValid
#cgo nocallback IsRenderTextureValid
#cgo nocallback IsTextureValid
#cgo nocallback ReloadImageFromTexture
#cgo nocallback SetTextureFilter
#cgo nocallback SetTextureWrap

#cgo noescape DrawTexture
#cgo noescape DrawTextureEx
#cgo noescape DrawTextureNPatch
#cgo noescape DrawTexturePro
#cgo noescape DrawTextureRec
#cgo noescape DrawTextureTiled
#cgo noescape DrawTextureV
#cgo noescape ExportImage
#cgo noescape ExportImageToMemory
#cgo noescape GenImageCellular
#cgo noescape GenImageChecked
#cgo noescape GenImageColor
#cgo noescape GenImageGradientLinear
#cgo noescape GenImageGradientRadial
#cgo noescape GenImageGradientSquare
#cgo noescape GenImagePerlinNoise
#cgo noescape GenImageText
#cgo noescape GenImageWhiteNoise
#cgo noescape GenTextureMipmaps
#cgo noescape GetImageColor
#cgo noescape ImageAlphaClear
#cgo noescape ImageAlphaCrop
#cgo noescape ImageAlphaMask
#cgo noescape ImageAlphaPremultiply
#cgo noescape ImageBlurGaussian
#cgo noescape ImageClearBackground
#cgo noescape ImageColorBrightness
#cgo noescape ImageColorContrast
#cgo noescape ImageColorGrayscale
#cgo noescape ImageColorInvert
#cgo noescape ImageColorReplace
#cgo noescape ImageColorTint
#cgo noescape ImageCopy
#cgo noescape ImageCrop
#cgo noescape ImageDither
#cgo noescape ImageDraw
#cgo noescape ImageDrawCircle
#cgo noescape ImageDrawCircleLines
#cgo noescape ImageDrawCircleLinesV
#cgo noescape ImageDrawCircleV
#cgo noescape ImageDrawLine
#cgo noescape ImageDrawLineEx
#cgo noescape ImageDrawLineV
#cgo noescape ImageDrawPixel
#cgo noescape ImageDrawPixelV
#cgo noescape ImageDrawRectangle
#cgo noescape ImageDrawRectangleLines
#cgo noescape ImageDrawRectangleRec
#cgo noescape ImageDrawRectangleV
#cgo noescape ImageDrawText
#cgo noescape ImageDrawTextEx
#cgo noescape ImageDrawTriangle
#cgo noescape ImageDrawTriangleEx
#cgo noescape ImageDrawTriangleFan
#cgo noescape ImageDrawTriangleLines
#cgo noescape ImageDrawTriangleStrip
#cgo noescape ImageFlipHorizontal
#cgo noescape ImageFlipVertical
#cgo noescape ImageFormat
#cgo noescape ImageFromChannel
#cgo noescape ImageFromImage
#cgo noescape ImageKernelConvolution
#cgo noescape ImageMipmaps
#cgo noescape ImageResize
#cgo noescape ImageResizeCanvas
#cgo noescape ImageResizeNN
#cgo noescape ImageRotate
#cgo noescape ImageRotateCCW
#cgo noescape ImageRotateCW
#cgo noescape ImageText
#cgo noescape ImageTextEx
#cgo noescape ImageToPOT
#cgo noescape IsImageValid
#cgo noescape IsRenderTextureValid
#cgo noescape IsTextureValid
#cgo noescape LoadImage
#cgo noescape LoadImageAnim
#cgo noescape LoadImageAnimFromMemory
#cgo noescape LoadImageColors
#cgo noescape LoadImageFromMemory
#cgo noescape LoadImageFromScreen
#cgo noescape LoadImageFromTexture
#cgo noescape LoadImageRaw
#cgo noescape LoadRenderTexture
#cgo noescape LoadTexture
#cgo noescape LoadTextureCubemap
#cgo noescape LoadTextureFromImage
#cgo noescape ReloadImageFromTexture
#cgo noescape SetTextureFilter
#cgo noescape SetTextureWrap
#cgo noescape UnloadImage
#cgo noescape UnloadImageColors
#cgo noescape UnloadRenderTexture
#cgo noescape UnloadTexture
#cgo noescape UpdateTexture
#cgo noescape UpdateTextureRec

#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"image"
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
	"github.com/Mishka-Squat/goex/mathex"
	"golang.org/x/exp/constraints"
)

// ToImage converts a Image to Go image.Image
func (i *Image) ToImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(i.Width), int(i.Height)))

	// Get pixel data from image (RGBA 32bit)
	cimg := i.cptr()
	ret := C.LoadImageColors(*cimg)
	pixels := (*[1 << 24]uint8)(unsafe.Pointer(ret))[0 : i.Width*i.Height*4]

	img.Pix = pixels

	return img
}

// NewImageFromImage - Returns new Image from Go image.Image
func NewImageFromImage(img image.Image) Image {
	size := img.Bounds().Size()

	cx := (C.int)(size.X)
	cy := (C.int)(size.Y)
	ccolor := ccolorptr(&White)
	ret := C.GenImageColor(cx, cy, *ccolor)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rcolor := MakeColor(uint8(r), uint8(g), uint8(b), uint8(a))
			ccolor = ccolorptr(&rcolor)

			cx = (C.int)(x)
			cy = (C.int)(y)
			C.ImageDrawPixel(&ret, cx, cy, *ccolor)
		}
	}
	return *newImageFromPointer(&ret)
}

// LoadImage - Load an image into CPU memory (RAM)
func LoadImage(fileName string) Image {
	cfileName := textAlloc(fileName)
	ret := C.LoadImage(cfileName)
	return *newImageFromPointer(&ret)
}

// LoadImageRaw - Load image data from RAW file
func LoadImageRaw(fileName string, width, height int32, format PixelFormat, headerSize int32) Image {
	cfileName := textAlloc(fileName)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cformat := (C.int)(format)
	cheaderSize := (C.int)(headerSize)
	ret := C.LoadImageRaw(cfileName, cwidth, cheight, cformat, cheaderSize)
	return *newImageFromPointer(&ret)
}

// LoadImageAnim - Load image sequence from file (frames appended to image.data)
func LoadImageAnim(fileName string, frames *int32) Image {
	cfileName := textAlloc(fileName)
	cframes := (*C.int)(frames)
	ret := C.LoadImageAnim(cfileName, cframes)
	return *newImageFromPointer(&ret)
}

// LoadImageAnimFromMemory - Load image sequence from memory buffer
func LoadImageAnimFromMemory(fileType string, fileData []byte, dataSize int32, frames *int32) *Image {
	cfileType := C.CString(fileType)
	defer C.free(unsafe.Pointer(cfileType))
	cfileData := (*C.uchar)(unsafe.Pointer(&fileData[0]))
	cdataSize := (C.int)(dataSize)
	cframes := (*C.int)(frames)
	ret := C.LoadImageAnimFromMemory(cfileType, cfileData, cdataSize, cframes)
	return newImageFromPointer(&ret)
}

// LoadImageFromMemory - Load image from memory buffer, fileType refers to extension: i.e. ".png"
func LoadImageFromMemory(fileType string, fileData []byte, dataSize int32) Image {
	cfileType := textAlloc(fileType)
	cfileData := (*C.uchar)(unsafe.Pointer(&fileData[0]))
	cdataSize := (C.int)(dataSize)
	ret := C.LoadImageFromMemory(cfileType, cfileData, cdataSize)
	return *newImageFromPointer(&ret)
}

// LoadImageFromTexture - Get pixel data from GPU texture and return an Image
func LoadImageFromTexture(texture Texture2D) Image {
	ctexture := texture.cptr()
	ret := C.LoadImageFromTexture(*ctexture)
	return *newImageFromPointer(&ret)
}

func ReloadImageFromTexture(texture Texture2D, image Image) Image {
	if !image.IsValid() {
		return LoadImageFromTexture(texture)
	}

	ctexture := texture.cptr()
	cimage := image.cptr()
	ret := C.ReloadImageFromTexture(*ctexture, cimage)
	return *newImageFromPointer(ret)
}

// LoadImageFromScreen - Load image from screen buffer (screenshot)
func LoadImageFromScreen() Image {
	ret := C.LoadImageFromScreen()
	return *newImageFromPointer(&ret)
}

// IsImageValid - Check if an image is valid (data and parameters)
func IsImageValid(image *Image) bool {
	cimage := image.cptr()
	ret := C.IsImageValid(*cimage)
	v := bool(ret)
	return v
}

// LoadTexture - Load an image as texture into GPU memory
func LoadTexture(fileName string) Texture2D {
	cfileName := textAlloc(fileName)
	ret := C.LoadTexture(cfileName)
	return *newTexture2DFromPointer(&ret)
}

// LoadTextureFromImage - Load a texture from image data
func LoadTextureFromImage(image Image) Texture2D {
	cimage := image.cptr()
	ret := C.LoadTextureFromImage(*cimage)
	return *newTexture2DFromPointer(&ret)
}

func ReloadTextureFromImage(image Image, texture Texture2D) Texture2D {
	if !texture.IsValid() {
		return LoadTextureFromImage(image)
	}

	UpdateTextureFromImage(texture, image)
	return texture
}

// LoadRenderTexture - Load a texture to be used for rendering
func LoadRenderTexture[WT, HT mathex.SignedNumber](width WT, height HT) RenderTexture2D {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadRenderTexture(cwidth, cheight)
	return *newRenderTexture2DFromPointer(&ret)
}

func LoadRenderTextureV[T mathex.SignedNumber](wh vector2.Vector[T]) RenderTexture2D {
	return LoadRenderTexture(wh.X, wh.Y)
}

// LoadTextureCubemap - Loads a texture for a cubemap using given layout
func LoadTextureCubemap(image Image, layout int32) Texture2D {
	cimage := image.cptr()
	clayout := (C.int)(layout)
	ret := C.LoadTextureCubemap(*cimage, clayout)
	return *newTexture2DFromPointer(&ret)
}

// UnloadImage - Unload image from CPU memory (RAM)
func UnloadImage(image *Image) {
	cimage := image.cptr()
	C.UnloadImage(cimage)
}

// IsTextureValid - Check if a texture is valid (loaded in GPU)
func IsTextureValid(texture *Texture2D) bool {
	ctexture := texture.cptr()
	ret := C.IsTextureValid(*ctexture)
	v := bool(ret)
	return v
}

// UnloadTexture - Unload texture from GPU memory
func UnloadTexture(texture *Texture2D) {
	ctexture := texture.cptr()
	C.UnloadTexture(ctexture)
}

// IsRenderTextureValid - Check if a render texture is valid (loaded in GPU)
func IsRenderTextureValid(target *RenderTexture2D) bool {
	ctarget := target.cptr()
	ret := C.IsRenderTextureValid(*ctarget)
	v := bool(ret)
	return v
}

// UnloadRenderTexture - Unload render texture from GPU memory
func UnloadRenderTexture(target *RenderTexture2D) {
	ctarget := target.cptr()
	C.UnloadRenderTexture(ctarget)
}

// LoadImageColors - Get pixel data from image as a Color slice
func LoadImageColors(img Image) []colorex.RGBA {
	cimg := img.cptr()
	ret := C.LoadImageColors(*cimg)
	return (*[1 << 24]colorex.RGBA)(unsafe.Pointer(ret))[0 : img.Width*img.Height]
}

// UnloadImageColors - Unload color data loaded with LoadImageColors()
func UnloadImageColors(cols []colorex.RGBA) {
	C.UnloadImageColors((*C.Color)(unsafe.Pointer(&cols[0])))
}

// UpdateTexture - Update GPU texture with new data
func UpdateTexture(texture *Texture2D, pixels []colorex.RGBA) {
	ctexture := texture.cptr()
	cpixels := unsafe.Pointer(&pixels[0])
	C.UpdateTexture(*ctexture, cpixels)
}

func UpdateTextureFromImage(texture Texture2D, image Image) {
	ctexture := texture.cptr()
	C.UpdateTexture(*ctexture, image.Data)
}

// UpdateTextureRec - Update GPU texture rectangle with new data
func UpdateTextureRec(texture Texture2D, rec rect2.Float32, pixels []colorex.RGBA) {
	ctexture := texture.cptr()
	cpixels := unsafe.Pointer(&pixels[0])
	crec := crect2ptr(&rec)
	C.UpdateTextureRec(*ctexture, *crec, cpixels)
}

// ExportImage - Export image as a PNG file
func ExportImage(image Image, fileName string) bool {
	cfileName := textAlloc(fileName)
	cimage := image.cptr()
	return bool(C.ExportImage(*cimage, cfileName))
}

// ExportImageToMemory - Export image to memory buffer
func ExportImageToMemory(image Image, fileType string) []byte {
	cfileType := textAlloc(fileType)
	cimage := image.cptr()

	var size C.int
	ret := C.ExportImageToMemory(*cimage, cfileType, &size)
	v := unsafe.Slice((*byte)(unsafe.Pointer(ret)), size)
	return v
}

// ImageCopy - Create an image duplicate (useful for transformations)
func ImageCopy(image *Image) Image {
	cimage := image.cptr()
	ret := C.ImageCopy(*cimage)
	return *newImageFromPointer(&ret)
}

// Create an image from another image piece
func ImageFromImage(image *Image, rec rect2.Float32) Image {
	cimage := image.cptr()
	crec := crect2ptr(&rec)
	ret := C.ImageFromImage(*cimage, *crec)
	return *newImageFromPointer(&ret)
}

// ImageFromChannel - Create an image from a selected channel of another image (GRAYSCALE)
func ImageFromChannel(image Image, selectedChannel int32) Image {
	cimage := image.cptr()
	cselectedChannel := C.int(selectedChannel)
	ret := C.ImageFromChannel(*cimage, cselectedChannel)
	return *newImageFromPointer(&ret)
}

// ImageText - Create an image from text (default font)
func ImageText(text string, fontSize int32, col colorex.RGBA) Image {
	ctext := textAlloc(text)
	cfontSize := (C.int)(fontSize)
	ccolor := ccolorptr(&col)
	ret := C.ImageText(ctext, cfontSize, *ccolor)
	return *newImageFromPointer(&ret)
}

// ImageTextEx - Create an image from text (custom sprite font)
func ImageTextEx(font Font, text string, fontSize, spacing float32, tint colorex.RGBA) Image {
	cfont := font.cptr()
	ctext := textAlloc(text)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ctint := ccolorptr(&tint)
	ret := C.ImageTextEx(*cfont, ctext, cfontSize, cspacing, *ctint)
	return *newImageFromPointer(&ret)
}

// ImageFormat - Convert image data to desired format
func ImageFormat(image *Image, newFormat PixelFormat) {
	cimage := image.cptr()
	cnewFormat := (C.int)(newFormat)
	C.ImageFormat(cimage, cnewFormat)
}

// ImageToPOT - Convert image to POT (power-of-two)
func ImageToPOT(image *Image, fillColor colorex.RGBA) {
	cimage := image.cptr()
	cfillColor := ccolorptr(&fillColor)
	C.ImageToPOT(cimage, *cfillColor)
}

// ImageCrop - Crop an image to a defined rectangle
func ImageCrop(image *Image, crop rect2.Float32) {
	cimage := image.cptr()
	ccrop := crect2ptr(&crop)
	C.ImageCrop(cimage, *ccrop)
}

// ImageAlphaCrop - Crop image depending on alpha value
func ImageAlphaCrop(image *Image, threshold float32) {
	cimage := image.cptr()
	cthreshold := (C.float)(threshold)
	C.ImageAlphaCrop(cimage, cthreshold)
}

// ImageAlphaClear - Apply alpha mask to image
func ImageAlphaClear(image *Image, col colorex.RGBA, threshold float32) {
	cimage := image.cptr()
	ccolor := ccolorptr(&col)
	cthreshold := (C.float)(threshold)
	C.ImageAlphaClear(cimage, *ccolor, cthreshold)
}

// ImageAlphaMask - Apply alpha mask to image
func ImageAlphaMask(image, alphaMask *Image) {
	cimage := image.cptr()
	calphaMask := alphaMask.cptr()
	C.ImageAlphaMask(cimage, *calphaMask)
}

// ImageAlphaPremultiply - Premultiply alpha channel
func ImageAlphaPremultiply(image *Image) {
	cimage := image.cptr()
	C.ImageAlphaPremultiply(cimage)
}

// ImageBlurGaussian - Apply box blur
func ImageBlurGaussian(image *Image, blurSize int32) {
	cimage := image.cptr()
	cblurSize := C.int(blurSize)
	C.ImageBlurGaussian(cimage, cblurSize)
}

// ImageKernelConvolution - Apply custom square convolution kernel to image
func ImageKernelConvolution(image *Image, kernel []float32) {
	cimage := image.cptr()
	ckernel := (*C.float)(unsafe.Pointer(&kernel[0]))
	C.ImageKernelConvolution(cimage, ckernel, C.int(len(kernel)))
}

// ImageResize - Resize an image (bilinear filtering)
func ImageResize(image *Image, newWidth, newHeight int32) {
	cimage := image.cptr()
	cnewWidth := (C.int)(newWidth)
	cnewHeight := (C.int)(newHeight)
	C.ImageResize(cimage, cnewWidth, cnewHeight)
}

// ImageResizeNN - Resize an image (Nearest-Neighbor scaling algorithm)
func ImageResizeNN(image *Image, newWidth, newHeight int32) {
	cimage := image.cptr()
	cnewWidth := (C.int)(newWidth)
	cnewHeight := (C.int)(newHeight)
	C.ImageResizeNN(cimage, cnewWidth, cnewHeight)
}

// ImageResizeCanvas - Resize canvas and fill with color
func ImageResizeCanvas(image *Image, newWidth, newHeight, offsetX, offsetY int32, col colorex.RGBA) {
	cimage := image.cptr()
	cnewWidth := (C.int)(newWidth)
	cnewHeight := (C.int)(newHeight)
	coffsetX := (C.int)(offsetX)
	coffsetY := (C.int)(offsetY)
	ccolor := ccolorptr(&col)
	C.ImageResizeCanvas(cimage, cnewWidth, cnewHeight, coffsetX, coffsetY, *ccolor)
}

// ImageMipmaps - Generate all mipmap levels for a provided image
func ImageMipmaps(image *Image) {
	cimage := image.cptr()
	C.ImageMipmaps(cimage)
}

// ImageDither - Dither image data to 16bpp or lower (Floyd-Steinberg dithering)
func ImageDither(image *Image, rBpp, gBpp, bBpp, aBpp int32) {
	cimage := image.cptr()
	crBpp := (C.int)(rBpp)
	cgBpp := (C.int)(gBpp)
	cbBpp := (C.int)(bBpp)
	caBpp := (C.int)(aBpp)
	C.ImageDither(cimage, crBpp, cgBpp, cbBpp, caBpp)
}

// ImageFlipVertical - Flip image vertically
func ImageFlipVertical(image *Image) {
	cimage := image.cptr()
	C.ImageFlipVertical(cimage)
}

// ImageFlipHorizontal - Flip image horizontally
func ImageFlipHorizontal(image *Image) {
	cimage := image.cptr()
	C.ImageFlipHorizontal(cimage)
}

// ImageRotate - Rotate image by input angle in degrees (-359 to 359)
func ImageRotate(image *Image, degrees int32) {
	cimage := image.cptr()
	cdegrees := (C.int)(degrees)
	C.ImageRotate(cimage, cdegrees)
}

// ImageRotateCW - Rotate image clockwise 90deg
func ImageRotateCW(image *Image) {
	cimage := image.cptr()
	C.ImageRotateCW(cimage)
}

// ImageRotateCCW - Rotate image counter-clockwise 90deg
func ImageRotateCCW(image *Image) {
	cimage := image.cptr()
	C.ImageRotateCCW(cimage)
}

// ImageColorTint - Modify image color: tint
func ImageColorTint(image *Image, col colorex.RGBA) {
	cimage := image.cptr()
	ccolor := ccolorptr(&col)
	C.ImageColorTint(cimage, *ccolor)
}

// ImageColorInvert - Modify image color: invert
func ImageColorInvert(image *Image) {
	cimage := image.cptr()
	C.ImageColorInvert(cimage)
}

// ImageColorGrayscale - Modify image color: grayscale
func ImageColorGrayscale(image *Image) {
	cimage := image.cptr()
	C.ImageColorGrayscale(cimage)
}

// ImageColorContrast - Modify image color: contrast (-100 to 100)
func ImageColorContrast(image *Image, contrast float32) {
	cimage := image.cptr()
	ccontrast := (C.float)(contrast)
	C.ImageColorContrast(cimage, ccontrast)
}

// ImageColorBrightness - Modify image color: brightness (-255 to 255)
func ImageColorBrightness(image *Image, brightness int32) {
	cimage := image.cptr()
	cbrightness := (C.int)(brightness)
	C.ImageColorBrightness(cimage, cbrightness)
}

// ImageColorReplace - Modify image color: replace color
func ImageColorReplace(image *Image, col, replace colorex.RGBA) {
	cimage := image.cptr()
	ccolor := ccolorptr(&col)
	creplace := ccolorptr(&replace)
	C.ImageColorReplace(cimage, *ccolor, *creplace)
}

// GetImageColor - Get image pixel color at (x, y) position
func GetImageColor(image *Image, x, y int32) colorex.RGBA {
	cimage := image.cptr()
	cx := (C.int)(x)
	cy := (C.int)(y)

	ret := C.GetImageColor(*cimage, cx, cy)
	return *gocolorptr(&ret)
}

// ImageClearBackground - Clear image background with given color
func ImageClearBackground(dst *Image, col colorex.RGBA) {
	cdst := dst.cptr()
	ccolor := ccolorptr(&col)
	C.ImageClearBackground(cdst, *ccolor)
}

// ImageDraw - Draw a source image within a destination image
func ImageDraw(dst *Image, src Image, srcRec, dstRec rect2.Float32, tint colorex.RGBA) {
	cdst := dst.cptr()
	csrc := src.cptr()
	csrcRec := crect2ptr(&srcRec)
	cdstRec := crect2ptr(&dstRec)
	ctint := ccolorptr(&tint)
	C.ImageDraw(cdst, *csrc, *csrcRec, *cdstRec, *ctint)
}

// ImageDrawLine - Draw line within an image
func ImageDrawLine(dst *Image, startPosX, startPosY, endPosX, endPosY int32, col colorex.RGBA) {
	cdst := dst.cptr()
	cstartPosX := (C.int)(startPosX)
	cstartPosY := (C.int)(startPosY)
	cendPosX := (C.int)(endPosX)
	cendPosY := (C.int)(endPosY)
	ccolor := ccolorptr(&col)
	C.ImageDrawLine(cdst, cstartPosX, cstartPosY, cendPosX, cendPosY, *ccolor)
}

// ImageDrawLineV - Draw line within an image, vector version
func ImageDrawLineV(dst *Image, start, end vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cstart := cvec2ptr(&start)
	cend := cvec2ptr(&end)
	ccolor := ccolorptr(&col)
	C.ImageDrawLineV(cdst, *cstart, *cend, *ccolor)
}

// ImageDrawLineEx - Draw a line defining thickness within an image
func ImageDrawLineEx(dst *Image, start, end vector2.Float32, thick int32, col colorex.RGBA) {
	cdst := dst.cptr()
	cstart := cvec2ptr(&start)
	cend := cvec2ptr(&end)
	cthick := C.int(thick)
	ccolor := ccolorptr(&col)
	C.ImageDrawLineEx(cdst, *cstart, *cend, cthick, *ccolor)
}

// ImageDrawCircle - Draw a filled circle within an image
func ImageDrawCircle(dst *Image, centerX, centerY, radius int32, col colorex.RGBA) {
	cdst := dst.cptr()
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradius := (C.int)(radius)
	ccolor := ccolorptr(&col)
	C.ImageDrawCircle(cdst, ccenterX, ccenterY, cradius, *ccolor)
}

// ImageDrawCircleV - Draw a filled circle within an image (Vector version)
func ImageDrawCircleV(dst *Image, center vector2.Float32, radius int32, col colorex.RGBA) {
	cdst := dst.cptr()
	ccenter := cvec2ptr(&center)
	cradius := (C.int)(radius)
	ccolor := ccolorptr(&col)
	C.ImageDrawCircleV(cdst, *ccenter, cradius, *ccolor)
}

// ImageDrawCircleLines - Draw circle outline within an image
func ImageDrawCircleLines(dst *Image, centerX, centerY, radius int32, col colorex.RGBA) {
	cdst := dst.cptr()
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradius := (C.int)(radius)
	ccolor := ccolorptr(&col)
	C.ImageDrawCircleLines(cdst, ccenterX, ccenterY, cradius, *ccolor)
}

// ImageDrawCircleLinesV - Draw circle outline within an image (Vector version)
func ImageDrawCircleLinesV(dst *Image, center vector2.Float32, radius int32, col colorex.RGBA) {
	cdst := dst.cptr()
	ccenter := cvec2ptr(&center)
	cradius := (C.int)(radius)
	ccolor := ccolorptr(&col)
	C.ImageDrawCircleLinesV(cdst, *ccenter, cradius, *ccolor)
}

// ImageDrawPixel - Draw pixel within an image
func ImageDrawPixel(dst *Image, posX, posY int32, col colorex.RGBA) {
	cdst := dst.cptr()
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	ccolor := ccolorptr(&col)
	C.ImageDrawPixel(cdst, cposX, cposY, *ccolor)
}

// ImageDrawPixelV - Draw pixel within an image (Vector version)
func ImageDrawPixelV(dst *Image, position vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cposition := cvec2ptr(&position)
	ccolor := ccolorptr(&col)
	C.ImageDrawPixelV(cdst, *cposition, *ccolor)
}

// ImageDrawRectangle - Draw rectangle within an image
func ImageDrawRectangle(dst *Image, x, y, width, height int32, col colorex.RGBA) {
	cdst := dst.cptr()
	cx := (C.int)(x)
	cy := (C.int)(y)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor := ccolorptr(&col)
	C.ImageDrawRectangle(cdst, cx, cy, cwidth, cheight, *ccolor)
}

// ImageDrawRectangleV - Draw rectangle within an image (Vector version)
func ImageDrawRectangleV(dst *Image, position, size vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cposition := cvec2ptr(&position)
	csize := cvec2ptr(&size)
	ccolor := ccolorptr(&col)
	C.ImageDrawRectangleV(cdst, *cposition, *csize, *ccolor)
}

// ImageDrawRectangleLines - Draw rectangle lines within an image
func ImageDrawRectangleLines(dst *Image, rec rect2.Float32, thick int, col colorex.RGBA) {
	cdst := dst.cptr()
	crec := crect2ptr(&rec)
	cthick := (C.int)(thick)
	ccolor := ccolorptr(&col)
	C.ImageDrawRectangleLines(cdst, *crec, cthick, *ccolor)
}

// ImageDrawTriangle - Draw triangle within an image
func ImageDrawTriangle(dst *Image, v1, v2, v3 vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cv1 := cvec2ptr(&v1)
	cv2 := cvec2ptr(&v2)
	cv3 := cvec2ptr(&v3)
	ccol := ccolorptr(&col)
	C.ImageDrawTriangle(cdst, *cv1, *cv2, *cv3, *ccol)
}

// ImageDrawTriangleEx - Draw triangle with interpolated colors within an image
func ImageDrawTriangleEx(dst *Image, v1, v2, v3 vector2.Float32, c1, c2, c3 colorex.RGBA) {
	cdst := dst.cptr()
	cv1 := cvec2ptr(&v1)
	cv2 := cvec2ptr(&v2)
	cv3 := cvec2ptr(&v3)
	cc1 := ccolorptr(&c1)
	cc2 := ccolorptr(&c2)
	cc3 := ccolorptr(&c3)
	C.ImageDrawTriangleEx(cdst, *cv1, *cv2, *cv3, *cc1, *cc2, *cc3)
}

// ImageDrawTriangleLines - Draw triangle outline within an image
func ImageDrawTriangleLines(dst *Image, v1, v2, v3 vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cv1 := cvec2ptr(&v1)
	cv2 := cvec2ptr(&v2)
	cv3 := cvec2ptr(&v3)
	ccol := ccolorptr(&col)
	C.ImageDrawTriangleLines(cdst, *cv1, *cv2, *cv3, *ccol)
}

// ImageDrawTriangleFan - Draw a triangle fan defined by points within an image (first vertex is the center)
func ImageDrawTriangleFan(dst *Image, points []vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	pointCount := C.int(len(points))
	ccol := ccolorptr(&col)
	C.ImageDrawTriangleFan(cdst, cpoints, pointCount, *ccol)
}

// ImageDrawTriangleStrip - Draw a triangle strip defined by points within an image
func ImageDrawTriangleStrip(dst *Image, points []vector2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	pointCount := C.int(len(points))
	ccol := ccolorptr(&col)
	C.ImageDrawTriangleStrip(cdst, cpoints, pointCount, *ccol)
}

// ImageDrawRectangleRec - Draw rectangle within an image
func ImageDrawRectangleRec(dst *Image, rec rect2.Float32, col colorex.RGBA) {
	cdst := dst.cptr()
	crec := crect2ptr(&rec)
	ccolor := ccolorptr(&col)
	C.ImageDrawRectangleRec(cdst, *crec, *ccolor)
}

// ImageDrawText - Draw text (default font) within an image (destination)
func ImageDrawText(dst *Image, posX, posY int32, text string, fontSize int32, col colorex.RGBA) {
	cdst := dst.cptr()
	posx := (C.int)(posX)
	posy := (C.int)(posY)
	ctext := textAlloc(text)
	cfontSize := (C.int)(fontSize)
	ccolor := ccolorptr(&col)
	C.ImageDrawText(cdst, ctext, posx, posy, cfontSize, *ccolor)
}

// ImageDrawTextEx - Draw text (custom sprite font) within an image (destination)
func ImageDrawTextEx(dst *Image, position vector2.Float32, font Font, text string, fontSize, spacing float32, col colorex.RGBA) {
	cdst := dst.cptr()
	cposition := cvec2ptr(&position)
	cfont := font.cptr()
	ctext := textAlloc(text)
	cfontSize := (C.float)(fontSize)
	cspacing := (C.float)(spacing)
	ccolor := ccolorptr(&col)
	C.ImageDrawTextEx(cdst, *cfont, ctext, *cposition, cfontSize, cspacing, *ccolor)
}

// GenImageColor - Generate image: plain color
func GenImageColor(width, height int, col colorex.RGBA) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor := ccolorptr(&col)

	ret := C.GenImageColor(cwidth, cheight, *ccolor)
	return *newImageFromPointer(&ret)
}

// GenImageGradientLinear - Generate image: linear gradient, direction in degrees [0..360], 0=Vertical gradient
func GenImageGradientLinear(width, height, direction int, start, end colorex.RGBA) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cdensity := (C.int)(direction)
	cstart := ccolorptr(&start)
	cend := ccolorptr(&end)

	ret := C.GenImageGradientLinear(cwidth, cheight, cdensity, *cstart, *cend)
	return *newImageFromPointer(&ret)
}

// GenImageGradientRadial - Generate image: radial gradient
func GenImageGradientRadial(width, height int, density float32, inner, outer colorex.RGBA) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cdensity := (C.float)(density)
	cinner := ccolorptr(&inner)
	couter := ccolorptr(&outer)

	ret := C.GenImageGradientRadial(cwidth, cheight, cdensity, *cinner, *couter)
	return *newImageFromPointer(&ret)
}

// GenImageGradientSquare - Generate image: square gradient
func GenImageGradientSquare(width, height int, density float32, inner, outer colorex.RGBA) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cdensity := (C.float)(density)
	cinner := ccolorptr(&inner)
	couter := ccolorptr(&outer)

	ret := C.GenImageGradientSquare(cwidth, cheight, cdensity, *cinner, *couter)
	return *newImageFromPointer(&ret)
}

// GenImageChecked - Generate image: checked
func GenImageChecked(width, height, checksX, checksY int, col1, col2 colorex.RGBA) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cchecksX := (C.int)(checksX)
	cchecksY := (C.int)(checksY)
	ccol1 := ccolorptr(&col1)
	ccol2 := ccolorptr(&col2)

	ret := C.GenImageChecked(cwidth, cheight, cchecksX, cchecksY, *ccol1, *ccol2)
	return *newImageFromPointer(&ret)
}

// GenImageWhiteNoise - Generate image: white noise
func GenImageWhiteNoise(width, height int, factor float32) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	cfactor := (C.float)(factor)

	ret := C.GenImageWhiteNoise(cwidth, cheight, cfactor)
	return *newImageFromPointer(&ret)
}

// GenImagePerlinNoise - Generate image: perlin noise
func GenImagePerlinNoise(width, height, offsetX, offsetY int, scale float32) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	coffsetX := (C.int)(offsetX)
	coffsetY := (C.int)(offsetY)
	cscale := (C.float)(scale)

	ret := C.GenImagePerlinNoise(cwidth, cheight, coffsetX, coffsetY, cscale)
	return *newImageFromPointer(&ret)
}

// GenImageCellular - Generate image: cellular algorithm. Bigger tileSize means bigger cells
func GenImageCellular(width, height, tileSize int) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ctileSize := (C.int)(tileSize)

	ret := C.GenImageCellular(cwidth, cheight, ctileSize)
	return *newImageFromPointer(&ret)
}

// GenImageText - Generate image: grayscale image from text data
func GenImageText(width, height int, text string) Image {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ctext := textAlloc(text)

	ret := C.GenImageText(cwidth, cheight, ctext)
	return *newImageFromPointer(&ret)
}

// GenTextureMipmaps - Generate GPU mipmaps for a texture
func GenTextureMipmaps(texture *Texture2D) {
	ctexture := texture.cptr()
	C.GenTextureMipmaps(ctexture)
}

// SetTextureFilter - Set texture scaling filter mode
func SetTextureFilter(texture Texture2D, filterMode TextureFilterMode) {
	ctexture := texture.cptr()
	cfilterMode := (C.int)(filterMode)
	C.SetTextureFilter(*ctexture, cfilterMode)
}

// SetTextureWrap - Set texture wrapping mode
func SetTextureWrap(texture Texture2D, wrapMode TextureWrapMode) {
	ctexture := texture.cptr()
	cwrapMode := (C.int)(wrapMode)
	C.SetTextureWrap(*ctexture, cwrapMode)
}

// DrawTexture - Draw a Texture2D
func DrawTexture[XT, YT constraints.Integer](texture Texture2D, posX XT, posY YT, tint colorex.RGBA) {
	ctexture := texture.cptr()
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	ctint := ccolorptr(&tint)
	C.DrawTexture(*ctexture, cposX, cposY, *ctint)
}

// DrawTextureV - Draw a Texture2D with position defined as vector2.Float32
func DrawTextureV(texture Texture2D, position vector2.Float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	cposition := cvec2ptr(&position)
	ctint := ccolorptr(&tint)
	C.DrawTextureV(*ctexture, *cposition, *ctint)
}

// DrawTextureEx - Draw a Texture2D with extended parameters
func DrawTextureEx(texture Texture2D, position vector2.Float32, rotation, scale float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	cposition := cvec2ptr(&position)
	crotation := (C.float)(rotation)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawTextureEx(*ctexture, *cposition, crotation, cscale, *ctint)
}

// DrawTextureRec - Draw a part of a texture defined by a rectangle
func DrawTextureRec(texture Texture2D, sourceRec rect2.Float32, position vector2.Float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	csourceRec := crect2ptr(&sourceRec)
	cposition := cvec2ptr(&position)
	ctint := ccolorptr(&tint)
	C.DrawTextureRec(*ctexture, *csourceRec, *cposition, *ctint)
}

// DrawTexturePro - Draw a part of a texture defined by a rectangle with 'pro' parameters
func DrawTexturePro(texture Texture2D, sourceRec, destRec rect2.Float32, origin vector2.Float32, rotation float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	csourceRec := crect2ptr(&sourceRec)
	cdestRec := crect2ptr(&destRec)
	corigin := cvec2ptr(&origin)
	crotation := (C.float)(rotation)
	ctint := ccolorptr(&tint)
	C.DrawTexturePro(*ctexture, *csourceRec, *cdestRec, *corigin, crotation, *ctint)
}

func DrawTextureTiled(texture Texture2D, source, dest rect2.Float32, origin vector2.Float32, rotation, scale float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	csource := crect2ptr(&source)
	cdest := crect2ptr(&dest)
	corigin := cvec2ptr(&origin)
	crotation := (C.float)(rotation)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawTextureTiled(*ctexture, *csource, *cdest, *corigin, crotation, cscale, *ctint)
}

func DrawTextureNPatchRect(texture Texture2D, nPatchInfo NPatchInfo, dest rect2.Float32) {
	DrawTextureNPatch(texture, nPatchInfo, dest, vector2.Zero[float32](), 0, White)
}

// DrawTextureNPatch - Draws a texture (or part of it) that stretches or shrinks nicely using n-patch info
func DrawTextureNPatch(texture Texture2D, nPatchInfo NPatchInfo, dest rect2.Float32, origin vector2.Float32, rotation float32, tint colorex.RGBA) {
	ctexture := texture.cptr()
	cnPatchInfo := nPatchInfo.cptr()
	cdest := crect2ptr(&dest)
	corigin := cvec2ptr(&origin)
	crotation := (C.float)(rotation)
	ctint := ccolorptr(&tint)
	C.DrawTextureNPatch(*ctexture, *cnPatchInfo, *cdest, *corigin, crotation, *ctint)
}
