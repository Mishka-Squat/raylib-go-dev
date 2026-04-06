package rl

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
)

func DrawTextLayout(font Font, text string, fontSize float32, spacing float32, tint colorex.RGBA, layoutFn func(wh vector2.Float32) rect2.Float32) {
	rect := layoutFn(MeasureTextEx(font, text, fontSize, spacing))
	DrawTextEx(font, text, rect.Position, fontSize, spacing, tint)
}
