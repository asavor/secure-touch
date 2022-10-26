package mact

import (
	"github.com/brianvoe/gofakeit"
	"math"
	"strconv"
	"strings"
)

type gen1 struct {
	polynomial float64
}

func MakeGen1(a ...float64) *gen1 {
	switch len(a) {
	case 0:
		return &gen1{
			polynomial: 9,
		}

	default:
		// Max polynomial degree is 19 for a function with max x & y of 1
		if a[0] > 19 {
			a[0] = 19
		}

		return &gen1{
			polynomial: a[0],
		}
	}
}

func (g *gen1) y(x float64) float64 {
	// -\left(x-0.01\right)^{19}+\sin^{-1}\left(x-0.01\right)\
	// -x^19 + sin^-1(x)
	// -(x-0.01)^19 + sin^-1(x-0.01)

	// translate to 0.01 right
	x = x - 0.01

	// polynomial
	poly := -math.Pow(x, g.polynomial)

	// inverse sin
	sin := math.Asin(x)

	function := poly + sin

	return function
}

type formula interface {
	y(x float64) float64
}

type calculator struct {
	formula
	x []float64

	xc []float64 // x compliment
	yc []float64 // y compliment
}

func MakeCalculator(f formula, xvalslen int) *calculator {
	xvals := GenX(xvalslen)
	return &calculator{f, xvals, make([]float64, len(xvals)), make([]float64, len(xvals))}
}

func (c *calculator) Rotate(deg float64) {
	for i, x := range c.x {
		y := c.formula.y(x)
		xsin := x * math.Sin(deg*math.Pi/180)
		xcos := x * math.Cos(deg*math.Pi/180)
		ysin := y * math.Sin(deg*math.Pi/180)
		ycos := y * math.Cos(deg*math.Pi/180)

		xc := xcos - ysin
		yc := xsin + ycos

		c.xc[i] = xc
		c.yc[i] = yc
	}
}

func (c *calculator) Scale(xScale, yScale float64) {
	for i, _ := range c.x {
		c.xc[i] = c.xc[i] * xScale
		c.yc[i] = c.yc[i] * yScale
	}
}

func (c *calculator) Values() (x, y []float64) {
	return c.xc, c.yc
}

func GenX(i int) []float64 {
	x := make([]float64, i)
	pow := 3.0

	for i := 0; i < len(x); i++ {
		v := float64(i) / 100.0
		noise := gofakeit.Float64Range(0, 0.005)
		v += noise

		// -x^3 + 1
		//x[i] = math.Pow(-v, pow) + 1

		// x^(1/3)
		x[i] = math.Pow(v, 1/pow)

		// x^3
		//x[i] = math.Pow(v, pow)

		//  sin(x)
		//x[i] = math.Sin(v)
	}

	return x
}

// Example
func makeMact(x, y float64) string {
	poly := gofakeit.Float64Range(1, 20)
	g1 := MakeGen1(poly)
	c := MakeCalculator(g1, 100)
	deg := gofakeit.Float64Range(0, 30)
	c.Rotate(deg)
	c.Scale(x, y)

	xSlice, ySlice := c.Values()

	str := strings.Builder{}
	ti := 0
	for i, _ := range xSlice {
		str.WriteString(strconv.Itoa(i) + ",1," + strconv.Itoa(ti) + "," + strconv.Itoa(int(xSlice[i])) + "," + strconv.Itoa(int(ySlice[i])) + ";")
		ti += gofakeit.Number(1, 100)
	}

	return str.String()
}

func TfylGenerateMact(mactCount, height, width int16) ([]float64, []float64) {

	poly := gofakeit.Float64Range(1, 20)
	g1 := MakeGen1(poly)
	c := MakeCalculator(g1, int(mactCount))
	c.Rotate(0)
	c.Scale(float64(width), float64(height))

	return c.Values()
}
