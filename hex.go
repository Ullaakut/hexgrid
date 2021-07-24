package hexgrid

import (
	"fmt"
	"math"
)

// Direction describes a direction.
type Direction string

const (
	// DirectionSE Southeast direction.
	DirectionSE Direction = "SE"
	// DirectionNE Northeast direction.
	DirectionNE = "NE"
	// DirectionN North direction.
	DirectionN = "N"
	// DirectionNW Northwest direction.
	DirectionNW = "NW"
	// DirectionSW Southwest direction.
	DirectionSW = "SW"
	// DirectionS South direction.
	DirectionS = "S"
)

var Directions = map[Direction]Hex{
	DirectionSE: NewHex(1, 0),
	DirectionNE: NewHex(1, -1),
	DirectionN:  NewHex(0, -1),
	DirectionNW: NewHex(-1, 0),
	DirectionSW: NewHex(-1, +1),
	DirectionS:  NewHex(0, +1),
}

// Hex describes a regular hexagon with Cube Coordinates (although the S coordinate is computed on the constructor)
// It's also easy to reference them as axial (trapezoidal coordinates):
// - R represents the vertical axis
// - Q the diagonal one
// - S can be ignored
// For additional reference on these coordinate systems: http://www.redblobgames.com/grids/hexagons/#coordinates
//           _ _
//         /     \
//    _ _ /(0,-1) \ _ _
//  /     \  -R   /     \
// /(-1,0) \ _ _ /(1,-1) \
// \  -Q   /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \  +Q   /
//  \ _ _ / (0,1) \ _ _ /
//        \  +R   /
//         \ _ _ /
type Hex struct {
	Q int // x axis
	R int // y axis
	S int // z axis
}

// NewHex constructs new Hex value with specified q and r.
func NewHex(q, r int) Hex {
	h := Hex{Q: q, R: r, S: -q - r}
	return h
}

func (h Hex) String() string {
	return fmt.Sprintf("(%d,%d)", h.Q, h.R)
}

// FractionalHex provides a more precise representation for hexagons when precision is required.
// It's also represented in Cube Coordinates
type FractionalHex struct {
	q float64
	r float64
	s float64
}

// NewFractionalHex constructs new FractionalHex value with specified q and r.
func NewFractionalHex(q, r float64) FractionalHex {
	h := FractionalHex{q: q, r: r, s: -q - r}
	return h
}

// Round rounds a FractionalHex to a Regular Hex
func (h FractionalHex) Round() Hex {
	roundToInt := func(a float64) int {
		if a < 0 {
			return int(a - 0.5)
		}
		return int(a + 0.5)
	}

	q := roundToInt(h.q)
	r := roundToInt(h.r)
	s := roundToInt(h.s)

	qDiff := math.Abs(float64(q) - h.q)
	rDiff := math.Abs(float64(r) - h.r)
	sDiff := math.Abs(float64(s) - h.s)

	if qDiff > rDiff && qDiff > sDiff {
		q = -r - s
	} else if rDiff > sDiff {
		r = -q - s
	} else {
		s = -q - r
	}
	return Hex{q, r, s}

}

// Add adds two hexagons
func Add(a, b Hex) Hex {
	return NewHex(a.Q+b.Q, a.R+b.R)
}

// Sub subtracts two hexagons
func Sub(a, b Hex) Hex {
	return NewHex(a.Q-b.Q, a.R-b.R)
}

// Scale scales an hexagon by a k factor. If factor k is 1 there's no change
func Scale(a Hex, k int) Hex {
	return NewHex(a.Q*k, a.R*k)
}

// Length returns a length of hex.
func Length(hex Hex) int {
	return int((math.Abs(float64(hex.Q)) + math.Abs(float64(hex.R)) + math.Abs(float64(hex.S))) / 2.)
}

// Distance returns a distance between two hexes.
func Distance(a, b Hex) int {
	return Length(Sub(a, b))
}

// Neighbor returns the neighbor hexagon at a certain direction
func Neighbor(h Hex, direction Direction) Hex {
	directionOffset := Directions[direction]
	return NewHex(h.Q+directionOffset.Q, h.R+directionOffset.R)
}

// Line returns the slice of hexagons that exist on a line that goes from hexagon a to hexagon b
func Line(a, b Hex) []Hex {
	hexLerp := func(a FractionalHex, b FractionalHex, t float64) FractionalHex {
		return NewFractionalHex(a.q*(1-t)+b.q*t, a.r*(1-t)+b.r*t)
	}

	N := Distance(a, b)

	// Sometimes the hexLerp will output a point that’s on an edge.
	// On some systems, the rounding code will push that to one side or the other,
	// somewhat unpredictably and inconsistently.
	// To make it always push these points in the same direction, add an “epsilon” value to a.
	// This will “nudge” things in the same direction when it’s on an edge, and leave other points unaffected.

	aNudge := NewFractionalHex(float64(a.Q)+0.000001, float64(a.R)+0.000001)
	bNudge := NewFractionalHex(float64(b.Q)+0.000001, float64(b.R)+0.000001)

	results := make([]Hex, 0)
	step := 1. / math.Max(float64(N), 1)

	for i := 0; i <= N; i++ {
		results = append(results, hexLerp(aNudge, bNudge, step*float64(i)).Round())
	}
	return results
}

// Range returns the set of hexagons around a certain center for a given radius
func Range(center Hex, radius int) []Hex {
	var results = make([]Hex, 0)

	if radius >= 0 {
		for dx := -radius; dx <= radius; dx++ {
			for dy := math.Max(float64(-radius), float64(-dx-radius)); dy <= math.Min(float64(radius), float64(-dx+radius)); dy++ {
				results = append(results, Add(center, NewHex(int(dx), int(dy))))
			}
		}
	}

	return results
}

// HasLineOfSight determines if a given hexagon is visible from another hexagon, taking into consideration a set of blocking hexagons
func HasLineOfSight(center Hex, target Hex, blocking []Hex) bool {
	contains := func(s []Hex, e Hex) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	for _, h := range Line(center, target) {
		if contains(blocking, h) {
			return false
		}
	}

	return true
}

// FieldOfView returns the list of hexagons that are visible from a given hexagon
func FieldOfView(source Hex, candidates []Hex, blocking []Hex) []Hex {
	results := make([]Hex, 0)

	for _, h := range candidates {
		distance := Distance(source, h)

		if len(blocking) == 0 || distance <= 1 || HasLineOfSight(source, h, blocking) {
			results = append(results, h)
		}
	}

	return results
}
