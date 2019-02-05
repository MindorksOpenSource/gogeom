//for quardrelaterals
/*
square->
sides(a) | side(diagonal)
perimeter, area, diagonal

rectangle->
sides(length,breadth) | side(length,diagonal) | side(breadth,diagonal)
perimeter, area, diagonal

parallelogram->
sides(base,slant,height)
perimeter, area

trapezium->
sides(base, upper, height) | sides(base, upper, slant1, slant2)
area, perimeter

rhombus->
sides(diagonal1, diagonal2) | sides(base, height)
area, each_side, perimeter

kite->
sides(diagonal1, diagonal2) | sides(smaller, larger)
area, perimeter

scaled->
sides(side1, side2, side3 side4)
perimeter
*/

//package gogeom
package main

import (
	"fmt"
	"math"
)

//sides of square
type SquareSides struct {
	side float64
}

//diagonals of square
type SquareDiagonals struct {
	diagonal float64
}

//sides of rectangle
type RectangleSides struct {
	length, breadth float64
}

//perimeter of square when the sides are there
func (q *SquareSides) PerimeterOfSquareWithSide() float64 {
	return 4 * q.side
}

//area of the square when the sides are there
func (q *SquareSides) AreaOfSquareWithSide() float64 {
	return math.Pow(q.side,2)
}

//diagonal of the square when the side is there
func (q *SquareSides) DiagonalOfSquare() float64 {
	return math.Sqrt(2) * q.side
}

//side of the square when diagonal is there
func (q *SquareDiagonals) SideOfSquare() float64 {
	return q.diagonal / math.Sqrt(2)
}

//perimeter of the square when the diagonals are there
func (q *SquareDiagonals) PerimeterOfSquareWithDiagonal() float64 {
	return 4 * q.SideOfSquare()
}

//area of the square when the diagonals are there
func (q *SquareDiagonals) AreaOfSquareWithDiagonal() float64 {
	return math.Pow(q.SideOfSquare(), 2)
}

//perimeter of rectangle
func (q *RectangleSides) PerimeterOfRectangle() float64 {
	return 2 * (q.length + q.breadth)
}

//area of rectangle
func (q *RectangleSides) AreaOfRectangle() float64 {
	return q.length * q.breadth
}

//diagonal of rectangle
func (q *RectangleSides) DiagonalOfRectangle() float64 {
	return math.Sqrt(math.Pow(q.length, 2) + math.Pow(q.breadth, 2))
}

func main() {
	input := RectangleSides {
		3,4,
	}
	fmt.Println(input.PerimeterOfRectangle())
	fmt.Println(input.AreaOfRectangle())
	fmt.Println(input.DiagonalOfRectangle())
}
