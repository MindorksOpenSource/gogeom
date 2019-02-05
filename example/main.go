package main

import (
	"fmt"

	geometry "github.com/MindorksOpenSource/gogeom"
)

func main() {
	line := geometry.GeneralLine{
		2, 9, 3,
	}
	ellipse := geometry.Ellipse{
		9, 2,
	}
	circle := geometry.RadiusFormOfCircle{
		2, 3, 4, 5, 6, 7,
	}
	parabola := geometry.ParabolaWithOrigin{
		2, 2, 3, false,
	}
	iso_triangle := geometry.IsoscelesTriangleSides {
		10, 3,
	}
	equi_triangle := geometry.EquilateralTriangleSides {
		4,
	}
	right_triangle := geometry.RightTriangleSides {
		3,4,5,
	}
	scalene_triangle := geometry.ScaleneTriangleSides {
		8,5,7,
	}
	square := geometry.SquareSides {
		12,
	}
	rectangle := geometry.RectangleSides {
		2,5,
	}
	parallelogram := geometry.ParallelogramSides {
		10,4,3,
	}
	rhombus := geometry.RhombusDiagonals {
		4,6,
	}
	trapezium := geometry.TrapeziumSidesWithHeight {
		10,8,4,
	}
	scalene_quad := geometry.ScaleneQuadrilateralSides {
		10.3,4,5.7,6.2,
	}
	n_agon5 := geometry.CountAndLengthAndApothemOfAgonSide {
		5,6,4, //for pentagon
	}
	n_agon10 := geometry.CountAndLengthAndApothemOfAgonSide {
		10,12,5.67, //for decagon
	}
	fmt.Println(ellipse.GetEccentricity())
	fmt.Println(line.SlopeOfLine())
	fmt.Println(line.XIntercept())
	fmt.Println(line.YIntercept())
	fmt.Println(circle.AreaOfCircles())
	fmt.Println(parabola.AxisEquation())
	fmt.Println(parabola.DirectrixEquation())
	fmt.Println(iso_triangle.PerimeterOfIsoscelesTriangle())
	fmt.Println(equi_triangle.AreaOfEquilateralTriangle())
	fmt.Println(right_triangle.AltitudeOfRightTriangleBase())
	fmt.Println(scalene_triangle.AreaOfScaleneTriangle())
	fmt.Println(square.AreaOfSquare())
	fmt.Println(rectangle.DiagonalOfRectangle())
	fmt.Println(parallelogram.AreaOfParallelogram())
	fmt.Println(rhombus.AreaOfRhombus())
	fmt.Println(trapezium.AreaOfTrapezium())
	fmt.Println(scalene_quad.PerimeterOfScaleneQuadrilateral())
	fmt.Println(n_agon5.PerimeterOfAgon())
	fmt.Println(n_agon10.AreaOfAgon())												
}
