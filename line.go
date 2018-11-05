package gogeom

import "math"

//Ax+By+c =0
type GeneralLine struct {
	A, B, C float64
}

//Ax+By+c =0
//Dx+Ey+F=0

type GeneralLines struct {
	A, B, C, D, E, F float64
}
type TwoPointFormLine struct {
	X1, Y1, X2, Y2 float64
}

// Slope of the line
func (l *GeneralLine) SlopeOfLine() float64 {
	return -(l.A / l.B)
}

//Y Intencept of the line
func (l *GeneralLine) YIntercept() float64 {
	return -(l.C / l.B)
}

//X Intencept of the line
func (l *GeneralLine) XIntercept() float64 {
	return -(l.C / l.A)
}

//Intersection of two lines Ax + By + C = 0 and Dx + Ey + F = 0
func (l *GeneralLines) IntersectionOfLines() (float64, float64) {
	denominator := (l.A*l.E - l.B*l.D)
	xNumerator := (l.B*l.F - l.C*l.E)
	yNumerator := (l.C*l.D - l.A*l.F)
	return xNumerator / denominator, yNumerator / denominator
}

// Slope of the line
func (l *TwoPointFormLine) SlopeOfLine() float64 {
	return (l.Y2 - l.Y1) / (l.X2 - l.X1)
}

// Middle Point of the line
func (l *TwoPointFormLine) MidPoints() (float64, float64) {
	return (l.X1 + l.X2) / 2, (l.Y1 + l.Y2) / 2
}

//Point (x, y) which divides the line connecting two points (x1 , y1) and (x2 , y2) in the ratio p:q
func (l *TwoPointFormLine) DividingPoints(p, q float64) (float64, float64) {
	denominator := (p + q)
	return ((p*l.X2 + q*l.X1) / denominator), ((p*l.Y2 + q*l.Y1) / denominator)
}

//Point (x, y) which divides the line connecting two points (x1 , y1) and (x2 , y2) externally at a ratio p:q
func (l *TwoPointFormLine) ExternalDividingPoints(p, q float64) (float64, float64) {
	denominator := (p - q)
	return ((p*l.X1 - q*l.X2) / denominator), ((p*l.Y1 + q*l.Y2) / denominator)
}

//A point (x, y) which is located at a distance d from a point (x1 , y1) on the line
func (l *GeneralLines) AngleBetweenTwoLines() float64 {
	numerator := (l.A*l.D + l.B*l.E)
	denominator := (math.Sqrt(PowerFunction(l.A, 2)+PowerFunction(l.B, 2)) * math.Sqrt(PowerFunction(l.D, 2)+PowerFunction(l.E, 2)))
	return math.Acos(numerator / denominator)
}

//Equation of a line passing through two points (x1 , y1), (x2 , y2)
func (l *TwoPointFormLine) LineThroughTwoPoint() string {
	slope := (l.Y2 - l.Y1) / (l.X2 - l.X1)
	constant := (l.X1 * slope) + l.Y1
	return "y = " + FloatToString(slope) + " x - ( " + FloatToString(constant) + " )"
}

//Equation of a line parallel to the line Ax + By + C = 0 and at a distance d from it.
func (l *GeneralLine) EquiDistantParallelLine(d float64) (string, string) {
	constant := FloatToString(d * math.Sqrt(PowerFunction(l.A, 2)+PowerFunction(l.B, 2)))
	return FloatToString(l.A) + "x + " + FloatToString(l.B) + "y + " + FloatToString(l.C) + " - " + constant, FloatToString(l.A) + "x + " + FloatToString(l.B) + "y + " + FloatToString(l.C) + " + " + constant
}

//Distance between two points (D)
func (l *TwoPointFormLine) DistanceBetweenTwoPoints() float64 {
	return math.Sqrt((PowerFunction((l.X2 - l.X1), 2)) + (PowerFunction((l.Y2 - l.Y1), 2)))
}

//Distance between intercepts xi and yi
func (l *TwoPointFormLine) DistanceBetweenInterecepts() float64 {
	return math.Sqrt((PowerFunction((l.X1), 2)) + (PowerFunction((l.Y1), 2)))
}
