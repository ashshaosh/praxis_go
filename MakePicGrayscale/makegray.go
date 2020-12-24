package main

import "golang.org/x/tour/pic"

// Pic func grayscale
func Pic(dx, dy int) [][]uint8 {
	slice2ret := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		sl := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			sl[x] = uint8((x + y) / 2) // x*y, x^y, x/33 * y ...
		}
		slice2ret[y] = sl
	}

	return slice2ret
}

func main() {
	pic.Show(Pic)
}
