package gogeom

import (
	"math"
)

//sides of the right triangle
type RightTriangleSides struct {
	Base, Height, Hypotenuse float64
}

//sides of the isosceles triangle
type IsoscelesTriangleSides struct {
	Base, Slants float64
}

//sides of the equilateral triangle
type EquilateralTriangleSides struct {
	Side float64
}

type ScaleneTriangleSides struct {
	Side_a, Side_b, Side_c float64
}

//perimeter of the right angled triangle
func (t *RightTriangleSides) PerimeterOfRightTriangle() float64 {
	return t.Base + t.Height + t.Hypotenuse
}

//area of the right angled triangle
func (t *RightTriangleSides) AreaOfRightTriangle() float64 {
	return (0.5)*(t.Base * t.Height)
}

//altitude of base of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleBase() float64 {
	return (t.Height * t.Hypotenuse) / t.Base
}

//altitude of height of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleHeight() float64 {
	return (t.Base * t.Hypotenuse) / t.Height
}

//altitude of hypotenuse of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleHypotenuse() float64 {
	return (t.Height * t.Base) / t.Hypotenuse
}

//perimeter of the isosceles triangle
func (t *IsoscelesTriangleSides) PerimeterOfIsoscelesTriangle() float64 {
	return t.Base + (2 * t.Slants)
}

//altitude of base (default altitude) of isosceles triangle
func (t *IsoscelesTriangleSides) AltitudeOfIsoscelesTriangle() float64 {
	return math.Sqrt(math.Pow(t.Slants, 2) - math.Pow(t.Base/2, 2))
}

//area of isosceles triangle
func (t *IsoscelesTriangleSides) AreaOfIsoscelesTriangle() float64 {
	return (0.5) * (t.Base * t.AltitudeOfIsoscelesTriangle())
}

//perimeter of equilateral triangle
func (t *EquilateralTriangleSides) PerimeterOfEquilateralTriangle() float64 {
	return 3 * t.Side
}

//area of equilateral triangle
func (t *EquilateralTriangleSides) AreaOfEquilateralTriangle() float64 {
	return math.Sqrt(3) * (math.Pow(t.Side,2) / 4)
}

//altitude/median/angle bisector/perpendicular bisector of equilateral triangle
func (t *EquilateralTriangleSides) AltitudeOfEquilateralTriangle() float64 {
	return math.Sqrt(3) * t.Side / 2
}

//perimeter of scalene triangle
func (t *ScaleneTriangleSides) PerimeterOfScaleneTriangle() float64 {
	return t.Side_a + t.Side_b + t.Side_c
}

//area of scalene triangle using heron's formula
func (t *ScaleneTriangleSides) AreaOfScaleneTriangle() float64 {
	p := t.PerimeterOfScaleneTriangle() / 2
	return math.Sqrt(p * (p - t.Side_a) * (p - t.Side_b) * (p - t.Side_c))
}
