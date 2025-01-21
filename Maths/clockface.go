package clockface

import (
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

// SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	io.WriteString(w, svgEnd)
}

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

/*
SecondHand is the unit vector of the second hand of an analogue clock at time `t`
represented as a Point.

scale -> Multiply the cos and sin from secondHandPoint with the length of second hands
For example: p = (0.5 * 90, 0.866 * 90)
p = Point{X: 45, Y: 77.94}

flip -> Since the SVG system has an inverted Y axis (0 starts at the top), the Y coordinates are reversed.
For example: p = Point{X: 45, Y: -77.94}

translate -> Moves the needle position from the origin (0,0) to the hour center (150, 150)
For example: p = (45 + 150, -77.94 + 150)
p = Point{X: 195, Y: 72.06}
*/
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

/*
secondHandPoint turns the radians into sin and cos

x = sin(radians) -> Give the horizontal position of X axis
y = cos(radians) -> Give the vertical position of Y axis
*/
func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

/*
secondsInRadians turns t into radian

1 second is 2π/60 radians.. cancel out the 2 and we get π/30 radians
OLD FORMULA: radians = t.Second() * (pi / 30)

but since float is horrible, we can write the formula as
NEW FORMULA: radians = π / (30 / t.Second())
*/
func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}
