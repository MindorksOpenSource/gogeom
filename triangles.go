//for triangle
/*
right->
sides(Base,Height,hypotenus)
Perimeter,Area,altitudeof(base),altitudeof(height),altitudeof(hyprothenus)

isosceles->
sides(base,slants)
Perimeter,Area,altitudeof(base),altitudeof(slants)

equilateral->
side(a)
Perimeter,Area,altitude,median,angle bisector,

scalene(other random triangles)->
sides(a,b)
Perimeter,Area
*/

//package gogeom
package main

import (
	"fmt"
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

func main() {
	inputSides := IsoscelesTriangleSides {
		8.0, 5.0,
	}
	fmt.Println(inputSides.PerimeterOfIsoscelesTriangle())
	fmt.Println(inputSides.AltitudeOfIsoscelesTriangle())
}
