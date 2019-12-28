package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMovementsFromString(t *testing.T) {
	str := "L1,U2,R3,D40"

	result := MovementsFromString(str)

	expectedResult := []Movement{
		Movement{'L', 1},
		Movement{'U', 2},
		Movement{'R', 3},
		Movement{'D', 40},
	}
	if !cmp.Equal(result, expectedResult) {
		t.Errorf("MovementsFromString failed, result %v, expectedResult %v\n", result, expectedResult)
	}
}
