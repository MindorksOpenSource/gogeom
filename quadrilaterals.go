//for quardrelaterals
package gogeom

import (
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

//sides of parallelogram
type ParallelogramSides struct {
	base, slant, height float64
}

//sides of trapezium with height
type TrapeziumSidesWithHeight struct {
	base, top, height float64
}

//sides of trapezium slants
type TrapeziumSidesWithSlants struct {
	base, top, slant1, slant2 float64
}

//diagonals of rhombus
type RhombusDiagonals struct {
	diagonal1, diagonal2 float64
}

//sides and height of rhombus
type RhombusSides struct {
	side, height float64
}

//sides of scalene quadrilateral
type ScaleneQuadrilateralSides struct {
	side1, side2, side3, side4 float64
}

//perimeter of square when the sides are there
func (q *SquareSides) PerimeterOfSquare() float64 {
	return 4 * q.side
}

//area of the square when the sides are there
func (q *SquareSides) AreaOfSquare() float64 {
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
func (q *SquareDiagonals) PerimeterOfSquare() float64 {
	return 4 * q.SideOfSquare()
}

//area of the square when the diagonals are there
func (q *SquareDiagonals) AreaOfSquare() float64 {
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

//perimeter of parallelogram
func (q *ParallelogramSides) PerimeterOfParallelogram() float64 {
	return 2 * (q.base + q.slant)
}

//area of parallelogram
func (q *ParallelogramSides) AreaOfParallelogram() float64 {
	return q.base * q.height
}

//area of trapezium
func (q *TrapeziumSidesWithHeight) AreaOfTrapezium() float64 {
	return 0.5 * (q.base + q.top) * q.height
}

//perimeter of trapezium
func (q *TrapeziumSidesWithSlants) PerimeterOfTrapezium() float64 {
	return q.base + q.top + q.slant1 + q.slant2
}

//area of rhombus with diagonals given
func (q *RhombusDiagonals) AreaOfRhombus() float64 {
	return 0.5 * q.diagonal1 * q.diagonal2
}

//length of each side of rhombus with diagonals given
func (q *RhombusDiagonals) LengthOfRhombusSides() float64 {
	return 0.5 * math.Sqrt(math.Pow(q.diagonal1, 2) + math.Pow(q.diagonal2, 2))
}

//perimeter of rhombus with diagonals given
func (q *RhombusDiagonals) PerimeterOfRhombus() float64 {
	return 4 * q.LengthOfRhombusSides()
}

//area of rhombus with sides and height given
func (q *RhombusSides) AreaOfRhombus() float64 {
	return q.side * q.height
}

//perimeter of rhombus with sides and height given
func (q *RhombusSides) PerimeterOfRhombus() float64 {
	return 4 * q.side
}

//perimeter of scalene quadrilateral
func (q *ScaleneQuadrilateralSides) PerimeterOfScaleneQuadrilateral() float64 {
	return q.side1 + q.side2 + q.side3 + q.side4
}
