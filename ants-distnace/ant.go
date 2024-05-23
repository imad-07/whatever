package main

import (
	"fmt"
)

func computeDistance(moves []int) int {
	x := 0
	y := 0

	for i := 0; i < len(moves); i++ {
		m := moves[i]

		switch m {
		case 0:
			{
				x++
				y--
			}
		case 1:
			{
				x++
				y++
			}
		case 2:
			{
				x--
				y++
			}
		case 3:
			{
				x--
				y--
			}
		}
	}
	fmt.Println(x, y)
	return Sqrt(NS(x*x + y*y))
}

func Sqrt(nb int) int {
	if nb == 1 || nb == 0 {
		return nb
	}
	for i := 0; i <= nb/2; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func NS(nb int) int {
	if nb == 1 || nb == 0 {
		return nb
	}
	for i := nb; i > 0; i-- {
		if Sqrt(nb) == 0 {
			nb--
		} else if Sqrt(nb) != 0 {
			return nb
		}
	}

	return nb
}

func main() {
	b := []int{0}
	l := computeDistance(b)
	fmt.Println(l)
}
