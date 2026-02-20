package rl

/*
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

// SetShapesTexture - Define default texture used to draw shapes
func SetShapesTexture(texture Texture2D, source rect2.Float32) {
	ctexture := texture.cptr()
	csource := crect2ptr(&source)
	C.SetShapesTexture(*ctexture, *csource)
}

// GetShapesTexture - Get texture that is used for shapes drawing
func GetShapesTexture() Texture2D {
	ret := C.GetShapesTexture()
	return *newTexture2DFromPointer(&ret)
}

// GetShapesTextureRectangle - Get texture source rectangle that is used for shapes drawing
func GetShapesTextureRectangle() rect2.Float32 {
	ret := C.GetShapesTextureRectangle()
	return *gorec2ptr(&ret)
}

// DrawPixel - Draw a pixel
func DrawPixel[XT, YT mathex.Number](posX XT, posY YT, col colorex.RGBA) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	ccolor := ccolorptr(&col)
	C.DrawPixel(cposX, cposY, *ccolor)
}

// DrawPixelV - Draw a pixel (Vector version)
func DrawPixelV(position vector2.Float32, col colorex.RGBA) {
	cposition := cvec2ptr(&position)
	ccolor := ccolorptr(&col)
	C.DrawPixelV(*cposition, *ccolor)
}

// DrawLine - Draw a line
func DrawLine[SXT, SYT, EXT, EYT mathex.Number](startPosX SXT, startPosY SYT, endPosX EXT, endPosY EYT, col colorex.RGBA) {
	cstartPosX := (C.int)(startPosX)
	cstartPosY := (C.int)(startPosY)
	cendPosX := (C.int)(endPosX)
	cendPosY := (C.int)(endPosY)
	ccolor := ccolorptr(&col)
	C.DrawLine(cstartPosX, cstartPosY, cendPosX, cendPosY, *ccolor)
}

// DrawLineV - Draw a line (Vector version)
func DrawLineV(startPos, endPos vector2.Float32, col colorex.RGBA) {
	cstartPos := cvec2ptr(&startPos)
	cendPos := cvec2ptr(&endPos)
	ccolor := ccolorptr(&col)
	C.DrawLineV(*cstartPos, *cendPos, *ccolor)
}

func DrawLineAB(a, b vector2.Float32) func(col colorex.RGBA) {
	return func(col colorex.RGBA) {
		la, lb := a, b
		cstartPos := cvec2ptr(&la)
		cendPos := cvec2ptr(&lb)
		ccolor := ccolorptr(&col)
		C.DrawLineV(*cstartPos, *cendPos, *ccolor)
	}
}

// DrawLineEx - Draw a line defining thickness
func DrawLineEx(startPos, endPos vector2.Float32, thick float32, col colorex.RGBA) {
	cstartPos := cvec2ptr(&startPos)
	cendPos := cvec2ptr(&endPos)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawLineEx(*cstartPos, *cendPos, cthick, *ccolor)
}

// DrawLineStrip - Draw lines sequence
func DrawLineStrip(points []vector2.Float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	ccolor := ccolorptr(&col)
	C.DrawLineStrip(cpoints, cpointCount, *ccolor)
}

// DrawLineBezier - Draw a line using cubic-bezier curves in-out
func DrawLineBezier(startPos, endPos vector2.Float32, thick float32, col colorex.RGBA) {
	cstartPos := cvec2ptr(&startPos)
	cendPos := cvec2ptr(&endPos)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawLineBezier(*cstartPos, *cendPos, cthick, *ccolor)
}

// DrawCircle - Draw a color-filled circle
func DrawCircle(centerX, centerY int32, radius float32, col colorex.RGBA) {
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradius := (C.float)(radius)
	ccolor := ccolorptr(&col)
	C.DrawCircle(ccenterX, ccenterY, cradius, *ccolor)
}

// DrawCircleSector - Draw a piece of a circle
func DrawCircleSector(center vector2.Float32, radius, startAngle, endAngle float32, segments int32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	cstartAngle := (C.float)(startAngle)
	cendAngle := (C.float)(endAngle)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawCircleSector(*ccenter, cradius, cstartAngle, cendAngle, csegments, *ccolor)
}

// DrawCircleSectorLines -
func DrawCircleSectorLines(center vector2.Float32, radius, startAngle, endAngle float32, segments int32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	cstartAngle := (C.float)(startAngle)
	cendAngle := (C.float)(endAngle)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawCircleSectorLines(*ccenter, cradius, cstartAngle, cendAngle, csegments, *ccolor)
}

// DrawCircleGradient - Draw a gradient-filled circle
func DrawCircleGradient(centerX, centerY int32, radius float32, col1, col2 colorex.RGBA) {
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradius := (C.float)(radius)
	ccolor1 := ccolorptr(&col1)
	ccolor2 := ccolorptr(&col2)
	C.DrawCircleGradient(ccenterX, ccenterY, cradius, *ccolor1, *ccolor2)
}

// DrawCircleV - Draw a color-filled circle (Vector version)
func DrawCircleV(center vector2.Float32, radius float32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	ccolor := ccolorptr(&col)
	C.DrawCircleV(*ccenter, cradius, *ccolor)
}

// DrawCircleLines - Draw circle outline
func DrawCircleLines(centerX, centerY int32, radius float32, col colorex.RGBA) {
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradius := (C.float)(radius)
	ccolor := ccolorptr(&col)
	C.DrawCircleLines(ccenterX, ccenterY, cradius, *ccolor)
}

// DrawCircleLinesV - Draw circle outline (Vector version)
func DrawCircleLinesV(center vector2.Float32, radius float32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	ccolor := ccolorptr(&col)
	C.DrawCircleLinesV(*ccenter, cradius, *ccolor)
}

// DrawEllipse - Draw ellipse
func DrawEllipse(centerX, centerY int32, radiusH, radiusV float32, col colorex.RGBA) {
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradiusH := (C.float)(radiusH)
	cradiusV := (C.float)(radiusV)
	ccolor := ccolorptr(&col)
	C.DrawEllipse(ccenterX, ccenterY, cradiusH, cradiusV, *ccolor)
}

// DrawEllipseLines - Draw ellipse outline
func DrawEllipseLines(centerX, centerY int32, radiusH, radiusV float32, col colorex.RGBA) {
	ccenterX := (C.int)(centerX)
	ccenterY := (C.int)(centerY)
	cradiusH := (C.float)(radiusH)
	cradiusV := (C.float)(radiusV)
	ccolor := ccolorptr(&col)
	C.DrawEllipseLines(ccenterX, ccenterY, cradiusH, cradiusV, *ccolor)
}

// DrawRing - Draw ring
func DrawRing(center vector2.Float32, innerRadius, outerRadius, startAngle, endAngle float32, segments int32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cinnerRadius := (C.float)(innerRadius)
	couterRadius := (C.float)(outerRadius)
	cstartAngle := (C.float)(startAngle)
	cendAngle := (C.float)(endAngle)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawRing(*ccenter, cinnerRadius, couterRadius, cstartAngle, cendAngle, csegments, *ccolor)
}

// DrawRingLines - Draw ring outline
func DrawRingLines(center vector2.Float32, innerRadius, outerRadius, startAngle, endAngle float32, segments int32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	cinnerRadius := (C.float)(innerRadius)
	couterRadius := (C.float)(outerRadius)
	cstartAngle := (C.float)(startAngle)
	cendAngle := (C.float)(endAngle)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawRingLines(*ccenter, cinnerRadius, couterRadius, cstartAngle, cendAngle, csegments, *ccolor)
}

// DrawRectangle - Draw a color-filled rectangle
func DrawRectangle[XT, YT, WT, HT mathex.Number](posX XT, posY YT, width WT, height HT, col colorex.RGBA) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor := ccolorptr(&col)
	C.DrawRectangle(cposX, cposY, cwidth, cheight, *ccolor)
}

// DrawRectangleV - Draw a color-filled rectangle (Vector version)
func DrawRectangleV(position vector2.Float32, size vector2.Float32, col colorex.RGBA) {
	cposition := cvec2ptr(&position)
	csize := cvec2ptr(&size)
	ccolor := ccolorptr(&col)
	C.DrawRectangleV(*cposition, *csize, *ccolor)
}

// DrawRectangleRec - Draw a color-filled rectangle
func DrawRectangleRec(rec rect2.Float32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	ccolor := ccolorptr(&col)
	C.DrawRectangleRec(*crec, *ccolor)
}

// DrawRectanglePro - Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec rect2.Float32, origin vector2.Float32, rotation float32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	corigin := cvec2ptr(&origin)
	crotation := (C.float)(rotation)
	ccolor := ccolorptr(&col)
	C.DrawRectanglePro(*crec, *corigin, crotation, *ccolor)
}

// DrawRectangleGradientV - Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientV(posX, posY, width, height int32, col1, col2 colorex.RGBA) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor1 := ccolorptr(&col1)
	ccolor2 := ccolorptr(&col2)
	C.DrawRectangleGradientV(cposX, cposY, cwidth, cheight, *ccolor1, *ccolor2)
}

// DrawRectangleGradientH - Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX, posY, width, height int32, col1, col2 colorex.RGBA) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor1 := ccolorptr(&col1)
	ccolor2 := ccolorptr(&col2)
	C.DrawRectangleGradientH(cposX, cposY, cwidth, cheight, *ccolor1, *ccolor2)
}

// DrawRectangleGradientEx - Draw a gradient-filled rectangle with custom vertex colors
func DrawRectangleGradientEx(rec rect2.Float32, col1, col2, col3, col4 colorex.RGBA) {
	crec := crect2ptr(&rec)
	ccolor1 := ccolorptr(&col1)
	ccolor2 := ccolorptr(&col2)
	ccolor3 := ccolorptr(&col3)
	ccolor4 := ccolorptr(&col4)
	C.DrawRectangleGradientEx(*crec, *ccolor1, *ccolor2, *ccolor3, *ccolor4)
}

// DrawRectangleLines - Draw rectangle outline
func DrawRectangleLines(posX, posY, width, height int32, col colorex.RGBA) {
	cposX := (C.int)(posX)
	cposY := (C.int)(posY)
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ccolor := ccolorptr(&col)
	C.DrawRectangleLines(cposX, cposY, cwidth, cheight, *ccolor)
}

// DrawRectangleLinesEx - Draw rectangle outline with extended parameters
func DrawRectangleLinesEx(rec rect2.Float32, lineThick float32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	clineThick := (C.float)(lineThick)
	ccolor := ccolorptr(&col)
	C.DrawRectangleLinesEx(*crec, clineThick, *ccolor)
}

// DrawRectangleRounded - Draw rectangle with rounded edges
func DrawRectangleRounded(rec rect2.Float32, roundness float32, segments int32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	croundness := (C.float)(roundness)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawRectangleRounded(*crec, croundness, csegments, *ccolor)
}

// DrawRectangleRoundedLines - Draw rectangle with rounded edges outline
func DrawRectangleRoundedLines(rec rect2.Float32, roundness float32, segments float32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	croundness := (C.float)(roundness)
	csegments := (C.int)(segments)
	ccolor := ccolorptr(&col)
	C.DrawRectangleRoundedLines(*crec, croundness, csegments, *ccolor)
}

// DrawRectangleRoundedLines - Draw rectangle with rounded edges outline
func DrawRectangleRoundedLinesEx(rec rect2.Float32, roundness float32, segments, lineThick float32, col colorex.RGBA) {
	crec := crect2ptr(&rec)
	croundness := (C.float)(roundness)
	csegments := (C.int)(segments)
	clineThick := (C.float)(lineThick)
	ccolor := ccolorptr(&col)
	C.DrawRectangleRoundedLinesEx(*crec, croundness, csegments, clineThick, *ccolor)
}

// DrawTriangle - Draw a color-filled triangle
func DrawTriangle(v1, v2, v3 vector2.Float32, col colorex.RGBA) {
	cv1 := cvec2ptr(&v1)
	cv2 := cvec2ptr(&v2)
	cv3 := cvec2ptr(&v3)
	ccolor := ccolorptr(&col)
	C.DrawTriangle(*cv1, *cv2, *cv3, *ccolor)
}

// DrawTriangleLines - Draw triangle outline
func DrawTriangleLines(v1, v2, v3 vector2.Float32, col colorex.RGBA) {
	cv1 := cvec2ptr(&v1)
	cv2 := cvec2ptr(&v2)
	cv3 := cvec2ptr(&v3)
	ccolor := ccolorptr(&col)
	C.DrawTriangleLines(*cv1, *cv2, *cv3, *ccolor)
}

// DrawTriangleFan - Draw a triangle fan defined by points
func DrawTriangleFan(points []vector2.Float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointsCount := (C.int)(len(points))
	ccolor := ccolorptr(&col)
	C.DrawTriangleFan(cpoints, cpointsCount, *ccolor)
}

// DrawTriangleStrip - Draw a triangle strip defined by points
func DrawTriangleStrip(points []vector2.Float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointsCount := (C.int)(int32(len(points)))
	ccolor := ccolorptr(&col)
	C.DrawTriangleStrip(cpoints, cpointsCount, *ccolor)
}

// DrawPoly - Draw a regular polygon (Vector version)
func DrawPoly(center vector2.Float32, sides int32, radius, rotation float32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	csides := (C.int)(sides)
	cradius := (C.float)(radius)
	crotation := (C.float)(rotation)
	ccolor := ccolorptr(&col)
	C.DrawPoly(*ccenter, csides, cradius, crotation, *ccolor)
}

// DrawPolyLines - Draw a polygon outline of n sides
func DrawPolyLines(center vector2.Float32, sides int32, radius, rotation float32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	csides := (C.int)(sides)
	cradius := (C.float)(radius)
	crotation := (C.float)(rotation)
	ccolor := ccolorptr(&col)
	C.DrawPolyLines(*ccenter, csides, cradius, crotation, *ccolor)
}

// DrawPolyLinesEx - Draw a polygon outline of n sides with extended parameters
func DrawPolyLinesEx(center vector2.Float32, sides int32, radius float32, rotation float32, lineThick float32, col colorex.RGBA) {
	ccenter := cvec2ptr(&center)
	csides := (C.int)(sides)
	cradius := (C.float)(radius)
	crotation := (C.float)(rotation)
	clineThick := (C.float)(lineThick)
	ccolor := ccolorptr(&col)
	C.DrawPolyLinesEx(*ccenter, csides, cradius, crotation, clineThick, *ccolor)
}

// DrawSplineLinear - Draw spline: Linear, minimum 2 points
func DrawSplineLinear(points []vector2.Float32, thick float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineLinear(cpoints, cpointCount, cthick, *ccolor)
}

// DrawSplineBasis - Draw spline: B-Spline, minimum 4 points
func DrawSplineBasis(points []vector2.Float32, thick float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineBasis(cpoints, cpointCount, cthick, *ccolor)
}

// DrawSplineCatmullRom - Draw spline: Catmull-Rom, minimum 4 points
func DrawSplineCatmullRom(points []vector2.Float32, thick float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineCatmullRom(cpoints, cpointCount, cthick, *ccolor)
}

// DrawSplineBezierQuadratic - Draw spline: Quadratic Bezier, minimum 3 points (1 control point): [p1, c2, p3, c4...]
func DrawSplineBezierQuadratic(points []vector2.Float32, thick float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineBezierQuadratic(cpoints, cpointCount, cthick, *ccolor)
}

// DrawSplineBezierCubic - Draw spline: Cubic Bezier, minimum 4 points (2 control points): [p1, c2, c3, p4, c5, c6...]
func DrawSplineBezierCubic(points []vector2.Float32, thick float32, col colorex.RGBA) {
	cpoints := (*C.Vector2)(unsafe.Pointer(&points[0]))
	cpointCount := (C.int)(len(points))
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineBezierCubic(cpoints, cpointCount, cthick, *ccolor)
}

// DrawSplineSegmentLinear - Draw spline segment: Linear, 2 points
func DrawSplineSegmentLinear(p1, p2 vector2.Float32, thick float32, col colorex.RGBA) {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineSegmentLinear(*cp1, *cp2, cthick, *ccolor)
}

// DrawSplineSegmentBasis - Draw spline segment: B-Spline, 4 points
func DrawSplineSegmentBasis(p1, p2, p3, p4 vector2.Float32, thick float32, col colorex.RGBA) {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineSegmentBasis(*cp1, *cp2, *cp3, *cp4, cthick, *ccolor)
}

// DrawSplineSegmentCatmullRom - Draw spline segment: Catmull-Rom, 4 points
func DrawSplineSegmentCatmullRom(p1, p2, p3, p4 vector2.Float32, thick float32, col colorex.RGBA) {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineSegmentCatmullRom(*cp1, *cp2, *cp3, *cp4, cthick, *ccolor)
}

// DrawSplineSegmentBezierQuadratic - Draw spline segment: Quadratic Bezier, 2 points, 1 control point
func DrawSplineSegmentBezierQuadratic(p1, p2, p3 vector2.Float32, thick float32, col colorex.RGBA) {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineSegmentBezierQuadratic(*cp1, *cp2, *cp3, cthick, *ccolor)
}

// DrawSplineSegmentBezierCubic - Draw spline segment: Cubic Bezier, 2 points, 2 control points
func DrawSplineSegmentBezierCubic(p1, p2, p3, p4 vector2.Float32, thick float32, col colorex.RGBA) {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	cthick := (C.float)(thick)
	ccolor := ccolorptr(&col)
	C.DrawSplineSegmentBezierCubic(*cp1, *cp2, *cp3, *cp4, cthick, *ccolor)
}

// GetSplinePointLinear - Get (evaluate) spline point: Linear
func GetSplinePointLinear(p1, p2 vector2.Float32, t float32) vector2.Float32 {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	ct := (C.float)(t)
	ret := C.GetSplinePointLinear(*cp1, *cp2, ct)
	return *govec2ptr(&ret)
}

// GetSplinePointBasis - Get (evaluate) spline point: B-Spline
func GetSplinePointBasis(p1, p2, p3, p4 vector2.Float32, t float32) vector2.Float32 {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	ct := (C.float)(t)
	ret := C.GetSplinePointBasis(*cp1, *cp2, *cp3, *cp4, ct)
	return *govec2ptr(&ret)
}

// GetSplinePointCatmullRom - Get (evaluate) spline point: Catmull-Rom
func GetSplinePointCatmullRom(p1, p2, p3, p4 vector2.Float32, t float32) vector2.Float32 {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	ct := (C.float)(t)
	ret := C.GetSplinePointCatmullRom(*cp1, *cp2, *cp3, *cp4, ct)
	return *govec2ptr(&ret)
}

// GetSplinePointBezierQuad - Get (evaluate) spline point: Quadratic Bezier
func GetSplinePointBezierQuad(p1, p2, p3 vector2.Float32, t float32) vector2.Float32 {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	ct := (C.float)(t)
	ret := C.GetSplinePointBezierQuad(*cp1, *cp2, *cp3, ct)
	return *govec2ptr(&ret)
}

// GetSplinePointBezierCubic - Get (evaluate) spline point: Cubic Bezier
func GetSplinePointBezierCubic(p1, p2, p3, p4 vector2.Float32, t float32) vector2.Float32 {
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	cp4 := cvec2ptr(&p4)
	ct := (C.float)(t)
	ret := C.GetSplinePointBezierCubic(*cp1, *cp2, *cp3, *cp4, ct)
	return *govec2ptr(&ret)
}

// CheckCollisionRecs - Check collision between two rectangles
func CheckCollisionRecs(rec1, rec2 rect2.Float32) bool {
	crec1 := crect2ptr(&rec1)
	crec2 := crect2ptr(&rec2)
	ret := C.CheckCollisionRecs(*crec1, *crec2)
	return bool(ret)
}

// CheckCollisionCircles - Check collision between two circles
func CheckCollisionCircles(center1 vector2.Float32, radius1 float32, center2 vector2.Float32, radius2 float32) bool {
	ccenter1 := cvec2ptr(&center1)
	cradius1 := (C.float)(radius1)
	ccenter2 := cvec2ptr(&center2)
	cradius2 := (C.float)(radius2)
	ret := C.CheckCollisionCircles(*ccenter1, cradius1, *ccenter2, cradius2)
	return bool(ret)
}

// CheckCollisionCircleRec - Check collision between circle and rectangle
func CheckCollisionCircleRec(center vector2.Float32, radius float32, rec rect2.Float32) bool {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	crec := crect2ptr(&rec)
	ret := C.CheckCollisionCircleRec(*ccenter, cradius, *crec)
	return bool(ret)
}

// CheckCollisionCircleLine - Check if circle collides with a line created betweeen two points [p1] and [p2]
func CheckCollisionCircleLine(center vector2.Float32, radius float32, p1, p2 vector2.Float32) bool {
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	ret := C.CheckCollisionCircleLine(*ccenter, cradius, *cp1, *cp2)
	return bool(ret)
}

// CheckCollisionPointRec - Check if point is inside rectangle
func CheckCollisionPointRec(point vector2.Float32, rec rect2.Float32) bool {
	cpoint := cvec2ptr(&point)
	crec := crect2ptr(&rec)
	ret := C.CheckCollisionPointRec(*cpoint, *crec)
	return bool(ret)
}

// CheckCollisionPointCircle - Check if point is inside circle
func CheckCollisionPointCircle(point vector2.Float32, center vector2.Float32, radius float32) bool {
	cpoint := cvec2ptr(&point)
	ccenter := cvec2ptr(&center)
	cradius := (C.float)(radius)
	ret := C.CheckCollisionPointCircle(*cpoint, *ccenter, cradius)
	return bool(ret)
}

// CheckCollisionPointTriangle - Check if point is inside a triangle
func CheckCollisionPointTriangle(point, p1, p2, p3 vector2.Float32) bool {
	cpoint := cvec2ptr(&point)
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cp3 := cvec2ptr(&p3)
	ret := C.CheckCollisionPointTriangle(*cpoint, *cp1, *cp2, *cp3)
	return bool(ret)
}

// CheckCollisionPointPoly - Check if point is within a polygon described by array of vertices
//
// NOTE: Based on http://jeffreythompson.org/collision-detection/poly-point.php
func CheckCollisionPointPoly(point vector2.Float32, points []vector2.Float32) bool {
	cpoint := cvec2ptr(&point)
	cpoints := cvec2ptr(&points[0])
	cpointCount := C.int(len(points))
	ret := C.CheckCollisionPointPoly(*cpoint, cpoints, cpointCount)
	return bool(ret)
}

// CheckCollisionLines - Check the collision between two lines defined by two points each, returns collision point by reference
func CheckCollisionLines(startPos1, endPos1, startPos2, endPos2 vector2.Float32, point *vector2.Float32) bool {
	cstartPos1 := cvec2ptr(&startPos1)
	cendPos1 := cvec2ptr(&endPos1)
	cstartPos2 := cvec2ptr(&startPos2)
	cendPos2 := cvec2ptr(&endPos2)
	cpoint := cvec2ptr(point)
	ret := C.CheckCollisionLines(*cstartPos1, *cendPos1, *cstartPos2, *cendPos2, cpoint)
	return bool(ret)
}

// CheckCollisionPointLine - Check if point belongs to line created between two points [p1] and [p2] with defined margin in pixels [threshold]
func CheckCollisionPointLine(point, p1, p2 vector2.Float32, threshold int32) bool {
	cpoint := cvec2ptr(&point)
	cp1 := cvec2ptr(&p1)
	cp2 := cvec2ptr(&p2)
	cthreshold := (C.int)(threshold)
	ret := C.CheckCollisionPointLine(*cpoint, *cp1, *cp2, cthreshold)
	return bool(ret)
}

// GetCollisionRec - Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1, rec2 rect2.Float32) rect2.Float32 {
	crec1 := crect2ptr(&rec1)
	crec2 := crect2ptr(&rec2)
	ret := C.GetCollisionRec(*crec1, *crec2)
	return *gorec2ptr(&ret)
}
