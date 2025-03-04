package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
	SidesCircle   intCustomType = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {

	var square float64
	if sidesNum == 3 {
		i := (sideLen + sideLen + sideLen) / 2
		square = math.Sqrt(i * (i - sideLen) * (i - sideLen) * (i - sideLen))
	} else if sidesNum == 4 {
		square = sideLen * sideLen
	} else if sidesNum == 0 {
		square = math.Pi * sideLen * sideLen
	} else {
		return 0
	}
	return square
}