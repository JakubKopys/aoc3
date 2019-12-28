package utils

import (
	"log"
	"strconv"
	"strings"
)

type Movement struct {
	Direction byte
	Distance  int
}

func MovementsFromString(str string) []Movement {
	segments := strings.Split(str, ",")

	movements := []Movement{}
	for _, segment := range segments {
		distance, err := strconv.Atoi(string(segment[1:]))
		if err != nil {
			log.Fatalf("Error when converting uint8 to int")
		}

		movements = append(
			movements,
			Movement{segment[0], distance},
		)
	}

	return movements
}
