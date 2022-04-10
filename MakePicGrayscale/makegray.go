package main

import "golang.org/x/tour/pic"

// MakeGrey func grayscale
func MakeGrey(dotx, doty int) [][]uint8 {
	slice2ret := make([][]uint8, doty)

	for y := 0; y < doty; y++ {
		sl := make([]uint8, dotx)
		for x := 0; x < dotx; x++ {
			sl[x] = uint8((x + y) / 2) // x*y, x^y, x/33 * y ...
		}
		slice2ret[y] = sl
	}

	return slice2ret
}

func main() {
	pic.Show(MakeGrey)
}
