package gogeom

import "math"

// if center is 0,0
type Ellipse struct {
	A, B float64
}

// Get Eccentricity of Locus
func (e *Ellipse) GetEccentricity() float64 {
	sqrt := math.Sqrt(PowerFunction(e.A, 2) - PowerFunction(e.B, 2))
	return sqrt / e.A
}

//Get Shape of Locus based on Eccentricity
func (e *Ellipse) GetShapeOfLocus() string {
	var output string
	if e.GetEccentricity() == 0 {
		output = "circle"
	} else if 0 < e.GetEccentricity() && e.GetEccentricity() < 1 {
		output = "ellipse"
	} else if e.GetEccentricity() == 1 {
		output = "parabola"
	} else if e.GetEccentricity() > 1 {
		output = "hyperbola"
	}
	return output
}

// Get Slope of Tangent of Line at x1,y1
func (e *Ellipse) GetSlopeOfTangentLine(x1, y1 float64) float64 {
	numerator := PowerFunction(e.B, 2) * x1
	denominator := PowerFunction(e.A, 2) * y1
	return (-(numerator / denominator))

}

// Eqn of Tangent at point x1,y1
func (e *Ellipse) GetTangentLineEquation(x1, y1 float64) string {
	xCoeff := (-(x1 * PowerFunction(e.B, 2) / (y1 * PowerFunction(e.A, 2))))
	constant := PowerFunction(e.B, 2) / y1
	return " y = " + FloatToString(xCoeff) + " x + " + FloatToString(constant)
}

// Ramunajan Approx Circumference
func (e *Ellipse) GetRamanujanApproxCircumference() float64 {
	h := PowerFunction((e.A-e.B), 2) / PowerFunction((e.A+e.B), 2)
	fract := 3 * h / (10 + math.Sqrt(4-3*h))
	return (math.Pi * (e.A + e.B) * (1 + fract))
}
