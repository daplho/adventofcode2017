// http://adventofcode.com/2017/day/3

package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := createSpiralPoint(289326)
	fmt.Printf("%d %d\n", int(x), int(y))
	fmt.Printf("Manhattan Distance: %d\n", int(math.Abs(x)+math.Abs(y)))
}

func createSpiralPoint(n float64) (float64, float64) {
	k := math.Ceil((math.Sqrt(n) - 1) / 2)
	t := 2*k + 1
	m := t * t
	t = t - 1

	if n >= m-t {
		return k - (m - n), -k
	} else {
		m = m - t
	}
	if n >= m-t {
		return -k, -k + (m - n)
	} else {
		m = m - t
	}
	if n >= m-t {
		return -k + (m - n), k
	} else {
		return k, k - (m - n - t)
	}
}
