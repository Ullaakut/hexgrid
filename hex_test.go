package hexgrid

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		hexB     Hex
		expected Hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(4, -10)},
	}

	for _, tt := range testCases {
		actual := Add(tt.hexA, tt.hexB)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

func TestSub(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		hexB     Hex
		expected Hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(-2, 4)},
	}

	for _, tt := range testCases {
		actual := Sub(tt.hexA, tt.hexB)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

func TestScale(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		factor   int
		expected Hex
	}{
		{NewHex(1, -3), 2, NewHex(2, -6)},
		{NewHex(-2, 3), 2, NewHex(-4, 6)},
	}

	for _, tt := range testCases {
		actual := Scale(tt.hexA, tt.factor)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//        \       /
//         \ _ _ /
// Tests that the neighbors of a certain hexagon are properly computed for all directions
func TestNeighbor(t *testing.T) {
	var testCases = []struct {
		origin    Hex
		direction Direction
		expected  Hex
	}{
		{NewHex(0, -1), DirectionSE, NewHex(1, -1)},
		{NewHex(0, -1), DirectionNE, NewHex(1, -2)},
		{NewHex(0, -1), DirectionN, NewHex(0, -2)},
		{NewHex(0, -1), DirectionNW, NewHex(-1, -1)},
		{NewHex(0, -1), DirectionSW, NewHex(-1, 0)},
		{NewHex(0, -1), DirectionS, NewHex(0, 0)},
	}

	for _, tt := range testCases {
		actual := Neighbor(tt.origin, tt.direction)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

// DISTANCE TESTS

//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \       /
//  \ _ _ / (0,1) \ _ _ /
//        \       /
//         \ _ _ /

func TestDistance(t *testing.T) {
	var testCases = []struct {
		origin      Hex
		destination Hex
		expected    int
	}{
		{NewHex(-1, -1), NewHex(1, -1), 2},
		{NewHex(-1, -1), NewHex(0, 0), 2},
		{NewHex(0, -1), NewHex(0, -2), 1},
		{NewHex(-1, -1), NewHex(0, 1), 3},
		{NewHex(1, 0), NewHex(-1, -1), 3},
	}

	for _, tt := range testCases {
		actual := Distance(tt.origin, tt.destination)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

//          _____         _____         _____
//         /     \       /     \       /     \
//   _____/ -2,-2 \_____/  0,-3 \_____/  2,-4 \_____
//  /     \       /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \_____/  3,-4 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/  2,-3 \_____/
//  /     \       /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \_____/  3,-3 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/  2,-2 \_____/
//  /     \       /     \       /     \       /     \
// / -3,1  \_____/ -1,0  \_____/  1,-1 \_____/  3,-2 \
// \       /     \       /     \       /     \       /
//  \_____/       \_____/       \_____/       \_____/
func TestLine(t *testing.T) {

	var testCases = []struct {
		origin      Hex
		destination Hex
		expected    string // the expected path serialized to string
	}{
		{NewHex(-3, -1), NewHex(3, -3), "[(-3,-1) (-2,-1) (-1,-2) (0,-2) (1,-2) (2,-3) (3,-3)]"},
		{NewHex(-2, 0), NewHex(2, -2), "[(-2,0) (-1,0) (0,-1) (1,-1) (2,-2)]"},
		{NewHex(1, -1), NewHex(1, -3), "[(1,-1) (1,-2) (1,-3)]"},
	}

	for _, tt := range testCases {
		actual := fmt.Sprint(Line(tt.origin, tt.destination))

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

// Tests that the range includes the correct number of hexagons with a certain radius from the center
//                 _____
//                /     \
//          _____/ -1,-2 \_____
//         /     \       /     \
//   _____/ -2,-1 \_____/  0,-2 \_____
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \       /     \       /
//  \_____/ -2,-2 \_____/  0,-3 \_____/
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \ CENTR /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/
//  /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \
// \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/
//        \       /     \       /
//         \_____/ -1,0  \_____/
//               \       /
//                \_____/
func TestRange(t *testing.T) {
	var testCases = []struct {
		radius                   int
		expectedNumberOfHexagons int
	}{
		{0, 1},
		{1, 7},
		{2, 19},
	}

	for _, tt := range testCases {
		actual := Range(NewHex(1, -2), tt.radius)

		if len(actual) != tt.expectedNumberOfHexagons {
			t.Error("Expected:", tt.expectedNumberOfHexagons, "got:", len(actual))
		}
	}
}

//    _ _           _ _           _ _
//  /     \       /     \       /     \
// /  0 0  \ _ _ /  2-1  \ _ _ /  4-2  \ _ _
// \       /     \   X   /     \   X   /     \
//  \ _ _ /  1 0  \ _ _ /  3-1  \ _ _ /  5-2  \
//  /     \       /# # #\   X   /     \   X   /
// /  0 1  \ _ _ /# 2 0 #\ _ _ /  4-1  \ _ _ /
// \       /     \#     #/# # #\   X   /     \
//  \ _ _ /  1 1  \#_#_#/# 3 0 #\ _ _ /  5-1  \
//  /     \  |P|  /     \#  X  #/     \   X   /
// /  0 2  \ _ _ /  2 1  \#_#_#/  4 0  \ _ _ /
// \       /     \       /     \   X   /     \
//  \ _ _ /  1 2  \ _ _ /  3 1  \ _ _ /  5 0  \
//  /     \       /     \       /     \       /
// /  0 3  \ _ _ /  2 2  \ _ _ /  4 1  \ _ _ /
// \       /     \       /     \       /     \
//  \ _ _ /  1 3  \ _ _ /  3 2  \ _ _ /  5 1  \
//        \       /     \       /     \       /
//         \ _ _ /       \ _ _ /       \ _ _ /
//
// The FOV measured from the central hex at 1,1, assuming blocking hexagons at 2,0 and 3,0.
// The hexagons marked with an X are non-visible. The remaining 16 are visible.
func TestFieldOfView(t *testing.T) {
	universe := []Hex{
		NewHex(0, 0),
		NewHex(0, 1),
		NewHex(0, 2),
		NewHex(0, 3),
		NewHex(1, 0),
		NewHex(1, 1),
		NewHex(1, 2),
		NewHex(1, 3),
		NewHex(2, -1),
		NewHex(2, 0),
		NewHex(2, 1),
		NewHex(2, 2),
		NewHex(3, -1),
		NewHex(3, 0),
		NewHex(3, 1),
		NewHex(3, 2),
		NewHex(4, -2),
		NewHex(4, -1),
		NewHex(4, 0),
		NewHex(4, 1),
		NewHex(5, -2),
		NewHex(5, -1),
		NewHex(5, 0),
		NewHex(5, 1),
	}

	losBlockers := []Hex{NewHex(2, 0), NewHex(3, 0)}
	actual := FieldOfView(NewHex(1, 1), universe, losBlockers)
	if len(actual) != 16 {
		t.Error("Expected: 16 got:", len(actual))
	}
}

////////////////
// Benchmarks //
////////////////

func BenchmarkDistance(b *testing.B) {
	var testCases = []struct {
		destination Hex
	}{
		{NewHex(0, 0)},
		{NewHex(100, 100)},
		{NewHex(10000, 10000)},
	}

	for _, bm := range testCases {
		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.destination), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Distance(origin, bm.destination)
			}
		})
	}
}

func BenchmarkLine(b *testing.B) {
	var testCases = []struct {
		destination Hex
	}{
		{NewHex(0, 0)},
		{NewHex(100, 100)},
		{NewHex(10000, 10000)},
	}

	for _, bm := range testCases {
		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.destination), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Line(origin, bm.destination)
			}
		})
	}
}

func BenchmarkRange(b *testing.B) {
	var testCases = []struct {
		radius int
	}{
		{0},
		{10},
		{100},
	}

	for _, bm := range testCases {
		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.radius), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Range(NewHex(1, -2), bm.radius)
			}
		})
	}
}

func BenchmarkHasLineOfSight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasLineOfSight(NewHex(1, 1), NewHex(4, -1), []Hex{NewHex(2, 0), NewHex(3, 0)})
	}
}
