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
	fmt.Println(ellipse.GetEccentricity())
	fmt.Println(line.SlopeOfLine())
	fmt.Println(line.XIntercept())
	fmt.Println(line.YIntercept())
	fmt.Println(circle.AreaOfCircles())
	fmt.Println(parabola.AxisEquation())

}
