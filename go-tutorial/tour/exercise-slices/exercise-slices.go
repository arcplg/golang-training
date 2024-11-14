package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) (out [][]uint8) {
	pic := make([][]uint8, dy, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8(x * x * y * y << 2)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}

func Bfield(xp, yp float64) (Bvalue float64) {
	const mu0 = 4e-7 * math.Pi     //Vacuum permeability
	const mu = 0.5 * mu0 / math.Pi //Multiplicative Factor
	const I1, I2 = 8e4, -8e4       //current in conductors 1 & 2
	const X1, Y1 = 128, 192        //abscissa, ordinate of conductor1
	const X2, Y2 = 128, 64         //abscissa, ordinate of conductor2
	Bxvalue1 := mu * I1 * (Y1 - yp) / (math.Pow(X1-xp, 2) + math.Pow(Y1-yp, 2))
	Byvalue1 := mu * I1 * (xp - X1) / (math.Pow(X1-xp, 2) + math.Pow(Y1-yp, 2))
	Bxvalue2 := mu * I2 * (Y2 - yp) / (math.Pow(X2-xp, 2) + math.Pow(Y2-yp, 2))
	Byvalue2 := mu * I2 * (xp - X2) / (math.Pow(X2-xp, 2) + math.Pow(Y2-yp, 2))
	Bvalue = (1e6 * (math.Pow(math.Pow(Bxvalue1+Bxvalue2, 2)+math.Pow(Byvalue1+Byvalue2, 2), 0.5)))
	return Bvalue
}
