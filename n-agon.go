//for regular (n)gon

package gogeom

//number of sides and length of side in the agon shape
type CountAndLengthOfAgonSide struct {
	side_count int
	side_length float64
	side_apothem float64
}

//perimeter of the agon
func (n *CountAndLengthOfAgonSide) PerimeterOfAgon() float64 {
	return float64(n.side_count) * n.side_length
}

//area of the agon
func (n *CountAndLengthOfAgonSide) AreaOfAgon() float64 {
	return 0.5 * n.PerimeterOfAgon() * n.side_apothem
}
