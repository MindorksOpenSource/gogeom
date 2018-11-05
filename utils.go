package gogeo

import (
	"math"
	"strconv"
)

// to minimise the power function
func PowerFunction(a float64, b float64) float64 {
	return math.Pow(a, b)
}

// Converts FLoat to String
func FloatToString(inputnum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputnum, 'f', 2, 64)
}
