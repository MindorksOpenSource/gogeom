package gogeo

import (
	"fmt"
	"math"
)

// (x-a)*2 + (y-b)*2 = r*2
// (x-c)*2 + (y-d)*2 = r*2
type RadiusFormOfCircle struct {
	A, B, R0, C, D, R1 float64
}

// x2 + y2 + Ax + By+ C = 0
// x2 + y2 + Dx + Ey + F = 0
type GeneralFormOfCircle struct {
	A, B, C, D, E, F float64
}

// DistanceBetweenTwoCenters is used to calculate the distance between two circles
func (c *RadiusFormOfCircle) DistanceBetweenTwoCenters() float64 {
	distance := math.Sqrt(PowerFunction(c.A-c.C, 2) + PowerFunction(c.B-c.D, 2))
	return distance
}

// Does Circle Intersect tells you if both the circle intersect or not
func (c *RadiusFormOfCircle) DoesCircleIntersect() bool {
	isInterescting := false
	if (c.R0 - c.R1) > 0 {
		if c.DistanceBetweenTwoCenters() > (c.R0-c.R1) && c.R0+c.R1 > c.DistanceBetweenTwoCenters() {
			isInterescting = true
		}
	} else if (c.R0 - c.R1) < 0 {
		if c.DistanceBetweenTwoCenters() > -(c.R0-c.R1) && c.R0+c.R1 > c.DistanceBetweenTwoCenters() {
			isInterescting = true
		}

	}
	return isInterescting
}

//Area of circle with Radius R0 and R1
func (c *RadiusFormOfCircle) AreaOfCircles() (float64, float64) {
	areaOne := math.Pi * PowerFunction(c.R0, 2)
	areaTwo := math.Pi * PowerFunction(c.R1, 2)
	return areaOne, areaTwo

}

//circumference of circle with Radius R0 and R1
func (c *RadiusFormOfCircle) CircumferenceOfCircles() (float64, float64) {
	circumferenceOne := math.Pi * c.R0 * 2
	circumferenceTwo := math.Pi * c.R1 * 2
	return circumferenceOne, circumferenceTwo

}

//Calculates the minimum delta
// ∂ is the area of the triangle formed by the two circle centers and one of the intersection point.
// The sides of this triangle are S, r0 and R0 , the area is calculated by Heron' s formula.
func (c *RadiusFormOfCircle) CalculateDelta() float64 {
	delta := math.Sqrt((c.DistanceBetweenTwoCenters()+c.R0+c.R1)*(c.DistanceBetweenTwoCenters()+c.R0-c.R1)*(c.DistanceBetweenTwoCenters()-c.R0+c.R1)*(-c.DistanceBetweenTwoCenters()+c.R0+c.R1)) / 4
	return delta
}

//Calculates  A,B,C,D
// (x-a)*2 + (y-b)*2 = r*2 is one equation
//(x-c)*2 + (y-d)*2 = r*2 is second equation
func (c *RadiusFormOfCircle) CalculateXY() (float64, float64, float64, float64) {
	distance := c.DistanceBetweenTwoCenters()
	delta := c.CalculateDelta()
	A := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	C := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	B := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	D := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	return A, B, C, D
}

//Calculates  A,B
// (x-a)*2 + (y-b)*2 = r*2 is one equation

func (c *RadiusFormOfCircle) CalculateAB() (float64, float64) {
	distance := c.DistanceBetweenTwoCenters()
	delta := c.CalculateDelta()
	A := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	B := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	return A, B
}

//Calculates  C,D
//(x-c)*2 + (y-d)*2 = r*2 is second equation

func (c *RadiusFormOfCircle) CalculateCD() (float64, float64) {
	distance := c.DistanceBetweenTwoCenters()
	delta := c.CalculateDelta()
	C := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	D := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	return C, D
}

//Calculates  only X values (A,C)
// (x-a)*2 + (y-b)*2 = r*2 is one equation
//(x-c)*2 + (y-d)*2 = r*2 is second equation

func (c *RadiusFormOfCircle) CalculateXs() (float64, float64) {
	distance := c.DistanceBetweenTwoCenters()
	delta := c.CalculateDelta()
	A := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	C := ((c.A + c.C) / 2) + ((c.C - c.A) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.B - c.D) / PowerFunction(distance, 2)) * delta)
	return A, C
}

// Calculates only Y Values (B,D)
// (x-a)*2 + (y-b)*2 = r*2 is one equation
//(x-c)*2 + (y-d)*2 = r*2 is second equation
func (c *RadiusFormOfCircle) CalculateYs() (float64, float64) {
	distance := c.DistanceBetweenTwoCenters()
	delta := c.CalculateDelta()
	B := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) - (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	D := ((c.B + c.D) / 2) + ((c.D - c.B) * (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2)) / (2 * PowerFunction(distance, 2))) + (2 * ((c.A - c.C) / PowerFunction(distance, 2)) * delta)
	return B, D
}

// Calculates the Line Equation connecting both intersection points
func (circle *RadiusFormOfCircle) LineEquationConnectingTwoIntersectionPoint() string {
	var a, b, c, d = circle.CalculateXY()
	fmt.Println(a, b, c, d)
	var numeratorConst = (b - d)
	var denominatorConst = (a - c)
	var output string
	var slope = numeratorConst / denominatorConst
	var constVal = b - (slope * a)
	fmt.Println(numeratorConst, denominatorConst, slope, constVal)

	if circle.DoesCircleIntersect() == true {
		output = "y = " + FloatToString(slope) + " x + ( " + FloatToString(constVal) + " )"
	} else {
		output = "Circle doesn't Interesect"
	}

	return output
}

// Calculates the Line Equation connecting both centers
func (c *RadiusFormOfCircle) LineEquationConnectingTwoCenters() string {
	var numeratorConst = (c.D - c.B)
	var denominatorConst = (c.C - c.A)
	var output string
	var slope = numeratorConst / denominatorConst
	var constVal = c.B - (slope * c.A)
	if c.DoesCircleIntersect() == true {
		output = "y = " + FloatToString(slope) + " x + ( " + FloatToString(constVal) + " )"
	} else {
		output = "Circle doesn't Interesect"
	}
	return output
}

//Checks if the two circle are tangential or not
func (c *RadiusFormOfCircle) IsTangent() bool {
	isTangent := false
	if (PowerFunction((c.A-c.C), 2)+(PowerFunction((c.B-c.D), 2)) == PowerFunction((c.R0-c.R1), 2)) || (PowerFunction((c.A-c.C), 2)+(PowerFunction((c.B-c.D), 2)) == PowerFunction((c.R0+c.R1), 2)) {
		isTangent = true
	}
	return isTangent
}

//Check if the two circle has Outer Circle Tangency
func (c *RadiusFormOfCircle) IsOuterCircleTangency() bool {
	isOuterCircleTangency := false
	distanceBewtweenCircles := c.DistanceBetweenTwoCenters()
	if (c.R0 + c.R1) == distanceBewtweenCircles {
		isOuterCircleTangency = true
	}
	return isOuterCircleTangency
}

//Check if the two circle has Inner Circle Tangency
func (c *RadiusFormOfCircle) IsInnerCircleTangency() bool {
	isInnerCircleTangency := false
	distanceBewtweenCircles := c.DistanceBetweenTwoCenters()
	if (c.R0-c.R1) > 0 && (c.R0-c.R1) == distanceBewtweenCircles {
		isInnerCircleTangency = true
	} else if (c.R0-c.R1) < 0 && (-(c.R0 - c.R1)) == distanceBewtweenCircles {
		isInnerCircleTangency = true
	}
	return isInnerCircleTangency
}

// Point of tangency of two circles
func (c *RadiusFormOfCircle) TangentPoint() (float64, float64) {
	denominator := (2 * (PowerFunction((c.C-c.A), 2) + PowerFunction((c.D-c.B), 2)))
	radiusDifference := (PowerFunction(c.R0, 2) - PowerFunction(c.R1, 2))
	x := (((c.A - c.C) * radiusDifference) / denominator) - ((c.A + c.C) / 2)
	y := (((c.B - c.D) * radiusDifference) / denominator) - ((c.B + c.D) / 2)
	return x, y

}

/*
	This is for the general form
*/

// DistanceBetweenTwoCenters is used to calculate the distance between two circles
func (c *GeneralFormOfCircle) DistanceBetweenTwoCenters() float64 {
	distance := (math.Sqrt(PowerFunction(c.A-c.D, 2) + PowerFunction(c.B-c.E, 2))) / 2
	return distance
}

// Calculates the Line Equation connecting both intersection points
func (c *GeneralFormOfCircle) LineEquationConnectingTwoIntersectionPoint() string {
	var output string
	var denominator = (c.B - c.E)
	var numOne = (c.D - c.A)
	var numTwo = (c.F - c.C)
	output = "y =" + FloatToString(numOne/denominator) + " x + ( " + FloatToString(numTwo/denominator) + " )"
	return output
}

// Calculates the Line Equation connecting both intersection points
func (c *GeneralFormOfCircle) SlopeOfConnectingLineOfTwoIntersectionPoint() float64 {
	var output float64
	var denominator = (c.B - c.E)
	var numOne = (c.D - c.A)
	output = numOne / denominator
	return output
}

/// Calculates the Line Equation connecting both centers
func (c *GeneralFormOfCircle) LineEquationConnectingTwoCenters() string {
	var denominator = (c.D - c.A)
	var numOne = (c.E - c.B)
	output := "y = " + FloatToString(numOne/denominator) + " x" + FloatToString((((numOne/denominator)*c.A)-c.B)/2)
	return output
}

//Checks if the two circle are tangential or not
func (c *GeneralFormOfCircle) IsTangent() bool {
	isTangent := false
	firstSqrt := math.Sqrt((PowerFunction(c.A, 2) + PowerFunction(c.B, 2) - (4 * c.C)))
	secondSqrt := math.Sqrt((PowerFunction(c.D, 2) + PowerFunction(c.E, 2) - (4 * c.F)))
	if PowerFunction((c.A-c.D), 2)+PowerFunction((c.B-c.E), 2) == PowerFunction((firstSqrt+secondSqrt), 2) || PowerFunction((c.A-c.D), 2)+PowerFunction((c.B-c.E), 2) == PowerFunction((firstSqrt-secondSqrt), 2) {
		isTangent = true

	}
	return isTangent
}

// Point of tangency of two circles
func (c *GeneralFormOfCircle) TangentPoint() (float64, float64) {
	u := (c.F - c.C)
	v := (c.D - c.A)
	w := (c.B - c.E)
	denominator := 2 * (PowerFunction(v, 2) + PowerFunction(w, 2))
	x := (2*u*v + c.A*PowerFunction(w, 2) + c.B*u*w) / denominator
	y := (-2*u*w + c.B*PowerFunction(v, 2) + c.A*v*w) / denominator
	return x, y
}

// Generate radius of both circle
func (c *GeneralFormOfCircle) RadiusOfTwoCircle() (float64, float64) {
	sqrtOne := PowerFunction(c.A, 2) + PowerFunction(c.B, 2) - 4*c.C
	sqrtTwo := PowerFunction(c.D, 2) + PowerFunction(c.E, 2) - 4*c.F
	return sqrtOne, sqrtTwo

}

//Check weather two circle are not intersecting or touching eaching other
func (c *RadiusFormOfCircle) AreTwoNonIntersectingCircle() bool {
	isNotInterescting := false
	distance := c.DistanceBetweenTwoCenters()
	if distance > (c.R0 + c.R1) {
		isNotInterescting = true
	}
	return isNotInterescting
}
