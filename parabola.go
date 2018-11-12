package gogeom

//y^2 = 4*a*x
type Parabola struct {
	A       float64
	IsYAxis bool
}

//(y - k)^2 = 4*a*(x - h)
type ParabolaWithOrigin struct {
	P       *Parabola
	H       float64
	K       float64
	IsYAxis bool
}

//Length of latus ration
func (p *Parabola) LenghtOfLatusRation() float64 {
	return (p.A * 4)
}

//Length of latus ration for shifted origin
func (p *ParabolaWithOrigin) LenghtOfLatusRation() float64 {
	return (p.P.A * 4)
}

// find the focus of a parabola
func (p *Parabola) FocusOfParabola() (float64, float64) {
	if p.IsYAxis == false {
		return p.A, 0
	} else {
		return 0, p.A
	}
}

// focus of shifted origin parabola
func (p *ParabolaWithOrigin) FocusOfParabola() (float64, float64) {
	if p.IsYAxis == false {
		return (p.H + p.P.A), p.K
	} else {
		return (p.H), (p.K + p.P.A)
	}
}

//Equation of directix
func (p *Parabola) DirectrixEquation() string {
	if p.IsYAxis == true {
		return ("y = " + FloatToString(-(p.A)))
	} else {
		return ("x = " + FloatToString(-(p.A)))
	}
}

//axis of normal parabola
func (p *Parabola) AxisEquation() string {
	if p.IsYAxis == true {
		return ("x = 0")
	} else {
		return ("y = 0")
	}
}

//axis of normal ParabolaWithOrigin
func (p *ParabolaWithOrigin) AxisEquation() string {
	if p.IsYAxis == true {
		return ("x = " + FloatToString(p.H-p.P.A))
	} else {
		return ("y = " + FloatToString(p.K-p.P.A))
	}
}

//Equation of directix ParabolaWithOrigin
func (p *ParabolaWithOrigin) DirectrixEquation() string {
	if p.IsYAxis == true {
		return ("y = " + FloatToString(p.K-(p.P.A)))
	} else {
		return ("x = " + FloatToString(p.H-(p.P.A)))
	}
}

//Equation of Vertex ParabolaWithOrigin
func (p *ParabolaWithOrigin) Vertex() (float64, float64) {
	return p.H, p.K
}

//Equation of Position Of Point
func (p *Parabola) PositionOfPoint(x, y float64) string {
	output := ""
	if PowerFunction(y, 2) == (4 * p.A * x) {
		output = "Point lies on Parabola"
	} else if PowerFunction(y, 2) > (4 * p.A * x) {
		output = "Point lies outside of Parabola"
	} else if PowerFunction(y, 2) < (4 * p.A * x) {
		output = "Point lies inside of Parabola"
	}
	return output
}

//Line y =mx+c is intersecting with the parabola
func (p *Parabola) PointOfInteresction(m, c float64) string {
	output := ""
	denominator := p.A / m
	if c == denominator {
		output = "meet the parabola at coincident points"
	} else if c < denominator {
		output = "intersect to parabola at two points"
	} else if c > denominator {
		output = "doesn't cut the parabola or touch it"
	}
	return output
}

//Equation of Tangent Equation of parabola
func (p *Parabola) TangentEquation(x, y float64) string {
	output := ""
	if p.IsYAxis == true {
		output = FloatToString(x) + "x = " + FloatToString(2*p.A) + " (x +" + FloatToString(x) + ")"
	} else {
		output = FloatToString(y) + "y = " + FloatToString(2*p.A) + " (y +" + FloatToString(y) + ")"
	}
	return output
}

//Equation of Normal Equation of parabola
func (p *Parabola) NormalEquation(x, y float64) string {
	LHS := FloatToString(y) + "y"
	RHS := FloatToString(-(y / 2 * p.A)) + " ( x - " + FloatToString(x) + " ) "
	return LHS + " = " + RHS
}

//Equation of Chord Of Contact Equation
func (p *Parabola) ChordOfContactEquation(x, y float64) string {
	return FloatToString(x) + "x = " + FloatToString(2) + " (x +" + FloatToString(x) + ")"
}

//Equation of Polar Of Point Equation
func (p *Parabola) PolarOfPoint(x, y float64) string {
	return FloatToString(x) + "x = " + FloatToString(2) + " (x +" + FloatToString(x) + ")"
}

// Point of Pole in Line
func (p *Parabola) PoleOfline(l, m float64) (float64, float64) {
	return m / l, (-2 * p.A * m) / l
}
