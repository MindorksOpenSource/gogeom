<p align="center"><img src="https://github.com/MindorksOpenSource/gogeomm/blob/master/example/gogeom.svg"></p>

# Gogeom - A Geometrical library for the Go programming language.
<!-- #### Refer [Wiki](https://github.com/MindorksOpenSource/gogeom/wiki/Circle) for detailed documentation -->
[![Mindorks](https://img.shields.io/badge/mindorks-opensource-blue.svg)](https://mindorks.com/open-source-projects) [![Mindorks Community](https://img.shields.io/badge/join-community-blue.svg)](https://mindorks.com/join-community) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

```
This is a Geometrical library for Go Language.
Which includes multiple Geometrical calculations like Circle, Lines etc in different forms.
```
## Installation

Installation is done using `go get`.
```
go get -u github.com/MindorksOpenSource/gogeom
```



#### These are following shape calculation which is supported by gogeom.
- [x] Circle
- [x] Line
- [x] Ellipse
- [ ] more to come..

# Circle
gogeom can handle two form of Circle.
- **Radius Form** 
  ``` 
  (x-a)^2 + (y-b)^2 = r^2 
  where,
  (a,b) is the center and "r" is the radius of the circle
  ``` 
- **General Form**
  ```
  x2 + y2 + Ax + By+ C = 0
  where,
  A,B and C are real numbers
  ```
#### Step to follow :


- import `geometry "github.com/MindorksOpenSource/gogeom"`
- Initalize two circles as (in Radius Form),
```
import geometry "github.com/MindorksOpenSource/gogeom"

func main() {
    // (x-a)^2 + (y-b)^2 = r^2 and (x-c)^2 + (y-d)^2 = s^2  are circle equation
    // r, s are radius of two circles
r := geometry.RadiusFormCircle{
        // a, b , r , c, d, s
        2, 3, 4, 5, 6, 7,
	}
}
```
-  Initalize two circles as (in General Form),
```
import `geometry "github.com/MindorksOpenSource/gogeom"

func main() {
// x2 + y2 + Ax + By+ C = 0 and  x2 + y2 + Dx + Ey + F = 0 are circle equation
g := geometry.GeneralFormOfCircle{
        // a, b , r , c, d, s
        2, 3, 4, 5, 6, 7,
	}
}

```
## Calculations for Circles
| Working                                   | Radius Form                               | General Form        |
| :-------------                         |:-------------                           | :-----            |
| **Distance** between two centers of circles| r.DistanceBetweenTwoCenters()    | g.DistanceBetweenTwoCenters()               |
| **Line  Equation** Connecting Two Centers| <small>r.LineEquationConnectingTwoCenters()</small>|   <small>g.LineEquationConnectingTwoCenters()</small>    |
| Check if both Circles **Intersect**| r.DoesCircleIntersect()||
| Check if both  Circles does not **Intersect**| r.AreTwoNonIntersectingCircle()||
|**Area** of Circle-1 and Circle-2 |r.AreaOfCircles()||
|**Circumference** of Circle-1 and Circle-2|r.CircumferenceOfCircles()||
|**∂** is the area of the triangle formed by the two circle centers and one of the intersection point. The sides of this triangle are S, r0 and R0 , the area is calculated by **Heron' s formula**.|r.CalculateDelta()||
|Calculates  **x1,y1,x2,y2**|r.CalculateXY()||
|Calculates **x1,y1**|r.CalculateAB()||
|Calculates  **x2,y2**|r.CalculateCD()||
|Calculates  **x1,x2**|r.CalculateXs()||
|Calculates  **y1,y2**|r.CalculateYs()||
|Calculates the **Line Equation** connecting both intersection points|<small>r.LineEquationConnectingTwoIntersectionPoint()</small>|<small>g.LineEquationConnectingTwoIntersectionPoint()</small>|
|Check if Both Circles Touch Each other as **tangent**|r.IsTangent()|g.IsTangent()|
|Check if two circles has **Outer Circle Tangency**|r.IsOuterCircleTangency()||
|Check if two circle has **Inner Circle Tangency**|r.IsInnerCircleTangency()||
|Tangential Points|r.TangentPoint()|g.TangentPoint()|
|Slope of Line of **Intersection Point**|<small>r.SlopeOfConnectingLineOfTwoIntersectionPoint()</small>|<small>g.SlopeOfConnectingLineOfTwoIntersectionPoint()</small>|
|Radius of Both Circls **R1,R2**||g.RadiusOfTwoCircle()|

# Line
gogeom can handle two form of Line.
- **General Form** 
  ``` 
  Ax+By+C =0
  where A, B and C are any real number and A and B are not both zero.``` 
- **Two Point Form**
  ```
  (x1,y1) and (x2,y2) form two lines passing through it
  ```

  #### Step to follow :


- import `geometry "github.com/MindorksOpenSource/gogeom"`
- Initalize two lines as (in **General** Form),
```
// initialise one line 
l := geometry.GeneralLine{
        // a, b, c
        2, 3, 4,
	}
}

---- or if you have to initialise multiple line ----
l := geometry.GeneralLines{
        // a, b, c, d, e, f
        2, 3, 4, 5, 6, 7,
	}
}
---- or if you have two point line ----
l := geometry.GeneralLines{
        // a, b, c, d, e, f
        2, 3, 4, 5, 
	}
}

```
- Initalize two lines as (in **Two Point Form Line** Form),

## Calculations for Line
| Working                                   | General Form (Single Line)  | General Form (Multiple Lines)                              | Two Point Form        |
| :-------------                         |:-------------                         | :-----              | :-----              |
|Slope of Line| l.SlopeOfLine()||l.SlopeOfLine()|
|Y-Intercept|l.YIntercept()|||
|X-Intercept|l.XIntercept()||
|Mid-Point of the line|||l.MidPoints()|
|Intersection of two lines **Ax + By + C = 0** and **Dx + Ey + F = 0**||l.IntersectionOfLines()||
|Point **(x, y)** which divides the line connecting two points **(x1 , y1)** and **(x2 , y2)** in the ratio **p:q**|||l.DividingPoints(p,q)|
|Point **(x, y)** which divides the line connecting two points **(x1 , y1)** and **(x2 , y2)** **externally** in the ratio **p:q**|||l.ExternalDividingPoints(p,q)|
|**Angle** Between Two Lines||l.AngleBetweenTwoLines()||
|**Line Eqn** passing two points|||l.LineThroughTwoPoint()|
|**EquiDistant Parallel Line**|l.EquiDistantParallelLine()|||
|Distance Between Two Points|||l.DistanceBetweenTwoPoints()|
|Distance Between Intercepts|||l.DistanceBetweenInterecepts()|
 
# Ellipse
Ellipse Equation,
```
x^2/a^2 + y^2/b^2 = 1,
(center at   x = 0   y = 0)
```       
 #### Step to follow :


- import `geometry "github.com/MindorksOpenSource/gogeom"`
- Initalize ellipse,
```
e := geometry.Ellipse {
        // a, b
        1,2
}
```
| Working                                   | General Form (Single Line) | Result
| :-------------                         |:-------------                         | :-----         
|Eccentricity of Ellipse|e.GetEccentricity()|float64|
|Shape Of Locus|e.GetShapeOfLocus()| Circle/Ellipse/Parabola/Hyerbola|
|Slope Of Tangent Line at **x1,y1**|e.GetSlopeOfTangentLine(x1,y1)| float64|
|Tangent Line Equation at **x1,y1**|e.GetTangentLineEquation(x1,y1)| string|
|Ramanujan Approx Circumference of ellipse|e.RamanujanApproxCircumference()| float64|

### TODO
* More features related to Geometrical Functions

## If this library helps you in anyway, show your love :heart: by putting a :star: on this project :v:

[Check out Mindorks awesome open source projects here](https://mindorks.com/open-source-projects)


#### Contributor
[Himanshu Singh](https://github.com/hi-manshu)


### License
```
   Copyright (C) 2018 MINDORKS NEXTGEN PRIVATE LIMITED

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```
