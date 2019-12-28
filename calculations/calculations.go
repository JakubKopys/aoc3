package calculations

import "math"

type Point struct {
	X int
	Y int
}

type Segment struct {
	A Point
	B Point
}

type Line struct {
	A int
	B int
	C int
}

func NewLine(segment Segment) *Line {
	line := new(Line)
	line.A = segment.A.Y - segment.B.Y
	line.B = segment.B.X - segment.A.X
	line.C = -(segment.A.X*segment.B.Y - segment.B.X*segment.A.Y)
	return line
}

func SegmentIntersection(segment1, segment2 Segment) *Point {
	line1 := *NewLine(segment1)
	line2 := *NewLine(segment2)
	point := intersection(line1, line2)
	return point
}

func intersection(line1, line2 Line) (intersection *Point) {
	D := line1.A*line2.B - line1.B*line2.A
	Dx := line1.C*line2.B - line1.B*line2.C
	Dy := line1.A*line2.C - line1.C*line2.A

	if D == 0 {
		return nil
	}

	x := Dx / D
	y := Dy / D
	return &Point{x, y}
}

func ManhattanDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}
