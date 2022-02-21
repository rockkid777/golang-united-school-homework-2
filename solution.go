package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
  "math"
)

type Side int

const (
  SidesCircle Side = 0
  SidesTriangle    = 3
  SidesSquare      = 4
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
  switch sidesNum {
  case 0:
    return (sideLen * sideLen * math.Pi)
  case 3:
    return ((math.Sqrt(3) / 4) * (sideLen * sideLen))
  case 4:
    return (sideLen * sideLen)
  default:
    return 0
  }
}
