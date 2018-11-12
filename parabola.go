package gogeom

//y^2 = 4*a*x
type Parabola struct {
	A       float64
	IsYAxis bool
}
type ParabolaWithOrigin struct {
	P       *Parabola
	H       float64
	K       float64
	IsYAxis bool
}

func (p *Parabola) LenghtOfLatusRation() float64 {
	return (p.A * 4)
}
func (p *ParabolaWithOrigin) LenghtOfLatusRation() float64 {
	return (p.P.A * 4)
}
func (p *Parabola) FocusOfParabola() (float64, float64) {
	if p.IsYAxis == false {
		return p.A, 0
	} else {
		return 0, p.A
	}
}

func (p *Parabola) DirectrixEquation() string {
	if p.IsYAxis == true {
		return ("y = " + FloatToString(-(p.A)))
	} else {
		return ("x = " + FloatToString(-(p.A)))
	}
}
func (p *ParabolaWithOrigin) Vertex() (float64, float64) {
	return p.H, p.K
}

func (p *ParabolaWithOrigin) FocusOfParabola() (float64, float64) {
	if p.IsYAxis == false {
		return (p.H + p.P.A), p.K
	} else {
		return (p.H), (p.K + p.P.A)
	}
}
