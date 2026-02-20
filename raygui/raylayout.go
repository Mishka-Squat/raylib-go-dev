package raygui

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/mathex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

type Justyfy int

const (
	JustifyTop Justyfy = iota
	JustifyCenter
	JustifyBottom
	JustifyFill
	JustifyLeft
	JustifyRight
)

// Justify a segment of width v in a segment of width max
// returns new segmend width amd distance from start
//
//	/--------max-------------/
//	/-----v----/
//	JustifyTop, JustifyLeft:
//	/-----v----/                 (v, 0)
//	JustifyCenter:
//	        /-----v----/         (v, (max - v) / 2)
//	JustifyBottom, JustifyRight:
//	              /-----v----/   (v, max - v)
//	JustifyFill:
//	/---------v--------------/   (max, 0)
func (j Justyfy) Justyfy(v, max float32) (nv, dv float32) {
	switch j {
	case JustifyTop:
	case JustifyLeft:
		return v, 0
	case JustifyCenter:
		return v, (max - v) / 2
	case JustifyBottom:
	case JustifyRight:
		return v, max - v
	case JustifyFill:
		return max, 0
	}

	return v, 0
}

// DOC(iga): Spreads items in interval [a, b] with distance [mins, maxd]
func Spread[S ~[]E, E any](s S, a, b vector2.Float32, maxd, mind float32, justify Justyfy) iter.Seq2[vector2.Float32, E] {
	count := len(s)
	switch count {
	case 0:
		return func(yield func(vector2.Float32, E) bool) {}
	case 1:
		return func(yield func(vector2.Float32, E) bool) { yield(a, s[0]) }
	}

	return func(yield func(vector2.Float32, E) bool) {
		ab := b.Sub(a)
		ab_length := ab.LengthF()
		direction := ab.Normalized()
		distance := mathex.Clamp(ab_length/(float32(count)-1), mind, maxd)

		_, dv := justify.Justyfy(min(distance*(float32(count)-1), ab_length), ab_length)
		position := a.Add(direction.ScaleF(dv))
		delta := direction.ScaleF(distance)
		for _, e := range s {
			if !yield(position, e) {
				return
			}

			position = position.Add(delta)
		}
	}
}

//type ЫефслувLayout interface {
//	Layout(wh vector2.Float32, justify Justyfy) rect2.Float32
//	Fill(height float32, justify Justyfy) rect2.Float32
//	Pie(percent float32) rect2.Float32
//}

type layout struct {
	Bounds rect2.Float32
}

func (l *layout) Anchor(xy vector2.Float32) vector2.Float32 {
	return l.Bounds.Size.NormalizeF(xy)
}

type CanvasLayoutPanel struct {
	layout
}

func CanvasLayout(bounds rect2.Float32) CanvasLayoutPanel {
	return CanvasLayoutPanel{
		layout: layout{Bounds: bounds},
	}
}

func (cl *CanvasLayoutPanel) Layout(anchor vector2.Float32, pivot vector2.Float32, wh vector2.Float32) rect2.Float32 {
	anchorp := anchor.MultByVector(cl.Bounds.Size)
	pivotp := pivot.MultByVector(wh)
	return rect2.NewFloat32(cl.Bounds.Position.Add(anchorp).Sub(pivotp), wh)
}

type HorizontalLayoutPanel struct {
	layout
	spacing  int
	position int
}

func HorizontalLayout(bounds rect2.Float32, spacing int) HorizontalLayoutPanel {
	return HorizontalLayoutPanel{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (hl *HorizontalLayoutPanel) Layout(wh vector2.Float32, justify Justyfy) rect2.Float32 {
	whY, dy := justify.Justyfy(wh.Y, hl.Bounds.Height())
	if wh.X < 0 {
		wh = wh.SetX(hl.Bounds.Width() - float32(hl.position) + wh.X - float32(hl.spacing))
	}
	r := rl.MakeRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, wh.X, whY)
	hl.position += int(wh.X) + hl.spacing
	return r
}

func (hl *HorizontalLayoutPanel) Fill(height float32, justify Justyfy) rect2.Float32 {
	whY, dy := justify.Justyfy(height, hl.Bounds.Height())
	r := rl.MakeRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, hl.Bounds.Width()-float32(hl.position), whY)
	hl.position = int(hl.Bounds.Width())
	return r
}

func (hl *HorizontalLayoutPanel) Pie(percent float32) rect2.Float32 {
	r := hl.Bounds.ShrinkXYWH(float32(hl.position), 0, 0, 0)
	r = r.ShrinkXYWH(0, 0, r.Width()*(1-percent), 0).Round()
	hl.position += int(r.Width())
	return r
}

type VerticalLayoutPanel struct {
	layout
	spacing  int
	position int
}

func VerticalLayout(bounds rect2.Float32, spacing int) VerticalLayoutPanel {
	return VerticalLayoutPanel{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (vl *VerticalLayoutPanel) Layout(wh vector2.Float32, justify Justyfy) rect2.Float32 {
	whX, dx := justify.Justyfy(wh.X, vl.Bounds.Width())
	if wh.Y < 0 {
		wh = wh.SetY(vl.Bounds.Height() - float32(vl.position) + wh.Y - float32(vl.spacing))
	}
	r := rl.MakeRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, wh.Y)
	vl.position += int(wh.Y) + vl.spacing
	return r
}

func (vl *VerticalLayoutPanel) Fill(width float32, justify Justyfy) rect2.Float32 {
	whX, dx := justify.Justyfy(width, vl.Bounds.Width())
	r := rl.MakeRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, vl.Bounds.Height()-float32(vl.position))
	vl.position = int(vl.Bounds.Height())
	return r
}

func (vl *VerticalLayoutPanel) Pie(percent float32) rect2.Float32 {
	r := vl.Bounds.ShrinkXYWH(0, float32(vl.position), 0, 0)
	r = r.ShrinkXYWH(0, 0, 0, r.Height()*(1-percent)).Round()
	vl.position += int(r.Height())
	return r
}
