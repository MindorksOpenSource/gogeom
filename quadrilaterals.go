//for quardrelaterals
package gogeom

import (
	"math"
)

//sides of square
type SquareSides struct {
	Side float64
}

//diagonals of square
type SquareDiagonals struct {
	Diagonal float64
}

//sides of rectangle
type RectangleSides struct {
	Length, Breadth float64
}

//sides of parallelogram
type ParallelogramSides struct {
	Base, Slant, Height float64
}

//sides of trapezium with height
type TrapeziumSidesWithHeight struct {
	Base, Top, Height float64
}

//sides of trapezium slants
type TrapeziumSidesWithSlants struct {
	Base, Top, Slant1, Slant2 float64
}

//diagonals of rhombus
type RhombusDiagonals struct {
	Diagonal1, Diagonal2 float64
}

//sides and height of rhombus
type RhombusSides struct {
	Side, Height float64
}

//sides of scalene quadrilateral
type ScaleneQuadrilateralSides struct {
	Side1, Side2, Side3, Side4 float64
}

//perimeter of square when the sides are there
func (q *SquareSides) PerimeterOfSquare() float64 {
	return 4 * q.Side
}

//area of the square when the sides are there
func (q *SquareSides) AreaOfSquare() float64 {
	return math.Pow(q.Side,2)
}

//diagonal of the square when the side is there
func (q *SquareSides) DiagonalOfSquare() float64 {
	return math.Sqrt(2) * q.Side
}

//side of the square when diagonal is there
func (q *SquareDiagonals) SideOfSquare() float64 {
	return q.Diagonal / math.Sqrt(2)
}

//perimeter of the square when the diagonals are there
func (q *SquareDiagonals) PerimeterOfSquare() float64 {
	return 4 * q.SideOfSquare()
}

//area of the square when the diagonals are there
func (q *SquareDiagonals) AreaOfSquare() float64 {
	return math.Pow(q.SideOfSquare(), 2)
}

//perimeter of rectangle
func (q *RectangleSides) PerimeterOfRectangle() float64 {
	return 2 * (q.Length + q.Breadth)
}

//area of rectangle
func (q *RectangleSides) AreaOfRectangle() float64 {
	return q.Length * q.Breadth
}

//diagonal of rectangle
func (q *RectangleSides) DiagonalOfRectangle() float64 {
	return math.Sqrt(math.Pow(q.Length, 2) + math.Pow(q.Breadth, 2))
}

//perimeter of parallelogram
func (q *ParallelogramSides) PerimeterOfParallelogram() float64 {
	return 2 * (q.Base + q.Slant)
}

//area of parallelogram
func (q *ParallelogramSides) AreaOfParallelogram() float64 {
	return q.Base * q.Height
}

//area of trapezium
func (q *TrapeziumSidesWithHeight) AreaOfTrapezium() float64 {
	return 0.5 * (q.Base + q.Top) * q.Height
}

//perimeter of trapezium
func (q *TrapeziumSidesWithSlants) PerimeterOfTrapezium() float64 {
	return q.Base + q.Top + q.Slant1 + q.Slant2
}

//area of rhombus with diagonals given
func (q *RhombusDiagonals) AreaOfRhombus() float64 {
	return 0.5 * q.Diagonal1 * q.Diagonal2
}

//length of each side of rhombus with diagonals given
func (q *RhombusDiagonals) LengthOfRhombusSides() float64 {
	return 0.5 * math.Sqrt(math.Pow(q.Diagonal1, 2) + math.Pow(q.Diagonal2, 2))
}

//perimeter of rhombus with diagonals given
func (q *RhombusDiagonals) PerimeterOfRhombus() float64 {
	return 4 * q.LengthOfRhombusSides()
}

//area of rhombus with sides and height given
func (q *RhombusSides) AreaOfRhombus() float64 {
	return q.Side * q.Height
}

//perimeter of rhombus with sides and height given
func (q *RhombusSides) PerimeterOfRhombus() float64 {
	return 4 * q.Side
}

//perimeter of scalene quadrilateral
func (q *ScaleneQuadrilateralSides) PerimeterOfScaleneQuadrilateral() float64 {
	return q.Side1 + q.Side2 + q.Side3 + q.Side4
}
