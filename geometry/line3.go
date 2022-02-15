package geometry

import (
	"github.com/gopherd/doge/math/mathutil"

	"github.com/gopherd/three/core"
)

type Line3 struct {
	Start, End core.Vector3
}

func (line Line3) Center() core.Vector3         { return line.Start.Add(line.End).Div(2) }
func (line Line3) Direction() core.Vector3      { return line.End.Sub(line.Start) }
func (line Line3) Square() core.Float           { return line.End.Sub(line.Start).Square() }
func (line Line3) Length() core.Float           { return line.End.Sub(line.Start).Length() }
func (line Line3) At(t core.Float) core.Vector3 { return line.Direction().Mul(t).Add(line.Start) }

func (line Line3) ClosestPointToPointParameter(point core.Vector3, clampToLine bool) core.Float {
	var startPoint = point.Sub(line.Start)
	var startEnd = line.Direction()
	var startEnd2 = startEnd.Dot(startEnd)
	var startEndStartPoint = startEnd.Dot(startPoint)
	var t = startEndStartPoint / startEnd2
	if clampToLine {
		t = mathutil.Clamp(t, 0, 1)
	}
	return t
}

func (line Line3) ClosestPointToPoint(point core.Vector3, clampToLine bool) core.Vector3 {
	var t = line.ClosestPointToPointParameter(point, clampToLine)
	return line.Direction().Mul(t).Add(line.Start)
}
