package calculations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntersectionFound(t *testing.T) {
	line1 := *NewLine(Segment{Point{1, 2}, Point{5, 2}})
	line2 := *NewLine(Segment{Point{4, 1}, Point{4, 3}})

	result := intersection(line1, line2)

	if !cmp.Equal(*result, Point{4, 2}) {
		t.Errorf("FindIntersection failed, result: %v\n", result)
	}
}

func TestIntersectionNotFound(t *testing.T) {
	line1 := *NewLine(Segment{Point{1, 1}, Point{2, 1}})
	line2 := *NewLine(Segment{Point{1, 2}, Point{2, 2}})

	result := intersection(line1, line2)

	if result != nil {
		t.Errorf("FindIntersection failed, result: %v\n", result)
	}
}
