package gogeom

import (
	"math"
)

//sides of the right triangle
type RightTriangleSides struct {
	base, height, hypotenuse float64
}

//sides of the isosceles triangle
type IsoscelesTriangleSides struct {
	base, slants float64
}

//sides of the equilateral triangle
type EquilateralTriangleSides struct {
	side float64
}

type ScaleneTriangleSides struct {
	side_a, side_b, side_c float64
}

//perimeter of the right angled triangle
func (t *RightTriangleSides) PerimeterOfRightTriangle() float64 {
	return t.base + t.height + t.hypotenuse
}

//area of the right angled triangle
func (t *RightTriangleSides) AreaOfRightTriangle() float64 {
	return (0.5)*(t.base * t.height)
}

//altitude of base of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleBase() float64 {
	return (t.height * t.hypotenuse) / t.base
}

//altitude of height of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleHeight() float64 {
	return (t.base * t.hypotenuse) / t.height
}

//altitude of hypotenuse of the right angled triangle
func (t *RightTriangleSides) AltitudeOfRightTriangleHypotenuse() float64 {
	return (t.height * t.base) / t.hypotenuse
}

//perimeter of the isosceles triangle
func (t *IsoscelesTriangleSides) PerimeterOfIsoscelesTriangle() float64 {
	return t.base + (2 * t.slants)
}

//altitude of base (default altitude) of isosceles triangle
func (t *IsoscelesTriangleSides) AltitudeOfIsoscelesTriangle() float64 {
	return math.Sqrt(math.Pow(t.slants, 2) - math.Pow(t.base/2, 2))
}

//area of isosceles triangle
func (t *IsoscelesTriangleSides) AreaOfIsoscelesTriangle() float64 {
	return (0.5) * (t.base * t.AltitudeOfIsoscelesTriangle())
}

//perimeter of equilateral triangle
func (t *EquilateralTriangleSides) PerimeterOfEquilateralTriangle() float64 {
	return 3 * t.side
}

//area of equilateral triangle
func (t *EquilateralTriangleSides) AreaOfEquilateralTriangle() float64 {
	return math.Sqrt(3) * (math.Pow(t.side,2) / 4)
}

//altitude/median/angle bisector/perpendicular bisector of equilateral triangle
func (t *EquilateralTriangleSides) AltitudeOfEquilateralTriangle() float64 {
	return math.Sqrt(3) * t.side / 2
}

//perimeter of scalene triangle
func (t *ScaleneTriangleSides) PerimeterOfScaleneTriangle() float64 {
	return t.side_a + t.side_b + t.side_c
}

//area of scalene triangle using heron's formula
func (t *ScaleneTriangleSides) AreaOfScaleneTriangle() float64 {
	p := t.PerimeterOfScaleneTriangle() / 2
	return math.Sqrt(p * (p - t.side_a) * (p - t.side_b) * (p - t.side_c))
}
