package rl

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/vector2"
)

type Monitor struct {
	Index       int
	Name        string
	Resolution  vector2.Int
	Position    vector2.Float32
	Dimensions  vector2.Int
	RefreshRate int
}

func MakeMonitor(i int) Monitor {
	return Monitor{
		Index: i,
		Name:  GetMonitorName(i),
		Resolution: vector2.New(
			GetMonitorWidth(i),
			GetMonitorHeight(i),
		),
		Position: GetMonitorPosition(i),
		Dimensions: vector2.New(
			GetMonitorPhysicalWidth(i),
			GetMonitorPhysicalHeight(i),
		),
		RefreshRate: GetMonitorRefreshRate(i),
	}
}

func EnumMonitors() iter.Seq[Monitor] {
	return func(yild func(Monitor) bool) {
		for i := range GetMonitorCount() {
			if !yild(MakeMonitor(i)) {
				return
			}
		}
	}
}

func (m Monitor) SetWindowMonitor() {
	SetWindowMonitor(m.Index)
}
