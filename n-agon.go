//for regular (n)gon

package gogeom

//number of sides and length of side in the agon shape
type CountAndLengthAndApothemOfAgonSide struct {
	Side_count int
	Side_length float64
	Side_apothem float64
}

//perimeter of the agon
func (n *CountAndLengthAndApothemOfAgonSide) PerimeterOfAgon() float64 {
	return float64(n.Side_count) * n.Side_length
}

//area of the agon
func (n *CountAndLengthAndApothemOfAgonSide) AreaOfAgon() float64 {
	return 0.5 * n.PerimeterOfAgon() * n.Side_apothem
}
