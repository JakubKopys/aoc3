package main

import (
	"aoc3/calculations"
	"aoc3/utils"
	"fmt"
	"log"
)

func main() {
	segments1 := segmentsFromString("R1005,U563,R417,U509,L237,U555,R397,U414,L490,U336,L697,D682,L180,U951,L189,D547,R697,U583,L172,D859,L370,D114,L519,U829,R389,U608,R66,D634,L320,D49,L931,U137,L349,D689,L351,D829,R819,D138,L118,D849,R230,U858,L509,D311,R815,U217,R359,U840,R77,U230,R361,U322,R300,D646,R348,U815,R793,D752,R967,U128,R948,D499,R359,U572,L566,U815,R630,D290,L829,D736,R358,U778,R891,U941,R544,U889,L920,U913,L447,D604,R538,U818,L215,D437,R447,U576,R452,D794,R864,U269,L325,D35,L268,D639,L101,U777,L776,U958,R105,U517,R667,D423,R603,U469,L125,D919,R879,U994,R665,D377,R456,D570,L685,U291,R261,U846,R840,U418,L974,D270,L312,D426,R621,D334,L855,D378,R694,U845,R481,U895,L362,D840,L712,U57,R276,D643,R566,U348,R361,D144,L287,D864,L556,U610,L927,U322,R271,D90,L741,U446,R181,D527,R56,U805,L907,D406,L286,U873,L79,D280,L153,D377,R253,D61,R475,D804,R788,U393,L660,U314,R489,D491,L234,D712,L253,U651,L777,D726,R146,U47,R630,U517,R226,U624,L834,D153,L513,U799,R287,D868,R982,U390,L296,D373,R9,U994,R105,D673,L657,D868,R738,D277,R374,U828,R860,U247,R484,U986,L723,D847,L578,U487,L51,D865,L328,D199,R812,D726,R355,D463,R761,U69,R508,D753,L81,D50,L345,D66,L764,D466,L975,U619,R59,D788,L737,D360,R14,D253,L512,D417,R828,D188,L394,U212,R658,U369,R920,U927,L339,U552,R856,D458,R407,U41,L930,D460,R809,U467,L410,D800,L135,D596,R678,D4,L771,D637,L876,U192,L406,D136,R666,U730,R711,D291,L586,U845,R606,U2,L228,D759,R244,U946,R948,U160,R397,U134,R188,U850,R623,D315,L219,D450,R489,U374,R299,D474,L767,D679,L160,D403,L708")
	segments2 := segmentsFromString("L1003,D878,R937,D979,R921,U572,R4,D959,L884,U394,R221,U206,R806,U912,R345,D290,R65,D996,L411,D157,R590,D557,L32,D360,L691,D861,L156,D603,R733,U444,L433,U144,L238,U213,R827,U949,R384,D409,L727,U923,L98,U781,L201,D200,R749,U288,L486,U158,L494,D522,R636,D330,L507,U691,R918,D706,R163,U609,R559,U674,R784,D87,R670,U401,L85,U981,R848,D579,L882,U777,R671,D385,R913,D899,R92,D780,L795,U821,R956,U446,L109,D955,L570,D874,R499,U845,R769,U88,L529,U657,R553,D357,L83,D324,L273,U689,L715,U933,R161,U561,L603,U349,L445,U781,R299,U26,L212,U429,R763,U116,R961,D258,L518,D668,L767,U587,L654,D24,R318,U35,L9,D199,L161,U419,R6,D707,R944,U499,R207,D349,L727,D637,R735,D137,R18,D214,L531,D327,L916,U440,R859,U483,R952,D631,L96,D320,L192,D985,R330,D196,L345,D575,L535,D868,R376,D126,R903,D619,L126,D624,L990,D67,L927,U685,L200,D759,L157,D816,L585,U910,R587,D598,L398,U706,R847,U682,L919,D291,L932,D54,L314,U430,L60,U206,L997,D487,L874,U957,L753,U999,R156,U102,L826,U923,L204,U293,L244,U787,L273,D687,R134,D167,L287,D459,R875,D32,R635,D400,L179,D19,L576,U60,L182,D409,R114,U329,R207,U525,L295,U305,L861,U280,R531,D49,L890,U521,L283,U37,R344,D867,L474,U893,R140,U289,L67,U490,R121,D34,L696,U902,R288,U249,R107,D750,R389,U125,L406,U950,R932,U795,R205,U583,L665,U214,R806,D409,R832,D39,R207,D977,L873,U645,L762,U847,L725,U397,R414,D558,L669,D736,R897,U464,R207,U359,R257,U304,L932,U240,L582,U409,L493,D481,R48,D537,R893,U48,R707,U630,L70,D289,L769,U98,L679,U504,L337,U117,L343,D574,R595,U168,R498")

	intersections := []calculations.Point{}
	for _, segment1 := range segments1 {
		for _, segment2 := range segments2 {
			intersection := calculations.SegmentIntersection(segment1, segment2)
			if intersection != nil {
				intersections = append(intersections, *intersection)
			}
		}
	}

	fmt.Println(intersections, len(intersections))

	startingPoint := calculations.Point{X: 0, Y: 0}
	min := 0
	for _, intersection := range intersections {
		if min == 0 {
			min = calculations.ManhattanDistance(startingPoint, intersection)
			continue
		}

		distance := calculations.ManhattanDistance(startingPoint, intersection)
		if distance < min {
			min = distance
		}
	}
	fmt.Println(min)
}

func segmentsFromString(str string) []calculations.Segment {
	movements := utils.MovementsFromString(str)
	return buildSegments(movements)
}

func buildSegments(movements []utils.Movement) []calculations.Segment {
	currentPoint := calculations.Point{X: 0, Y: 0}
	segments := []calculations.Segment{}

	addSegment := func(newPoint calculations.Point) {
		newSegment := calculations.Segment{A: currentPoint, B: newPoint}
		segments = append(segments, newSegment)
		currentPoint = newPoint
	}

	for _, movement := range movements {
		switch movement.Direction {
		case 'L':
			nextPoint := calculations.Point{
				X: currentPoint.X - movement.Distance,
				Y: currentPoint.Y,
			}
			addSegment(nextPoint)
		case 'R':
			nextPoint := calculations.Point{
				X: currentPoint.X + movement.Distance,
				Y: currentPoint.Y,
			}
			addSegment(nextPoint)
		case 'U':
			nextPoint := calculations.Point{
				X: currentPoint.X,
				Y: currentPoint.Y + movement.Distance,
			}
			addSegment(nextPoint)
		case 'D':
			nextPoint := calculations.Point{
				X: currentPoint.X,
				Y: currentPoint.Y - movement.Distance,
			}
			addSegment(nextPoint)
		default:
			log.Fatal("Invalid direction\n")
		}
	}
	return segments
}
