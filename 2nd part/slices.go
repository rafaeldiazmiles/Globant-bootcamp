package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slidePadre := make([][]uint8, dy)
	// en el ejemplo que vi usan y en vez de i y x en vez de j
	for i := 0; i < dy; i++ {
		slidePadre[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slidePadre[i][j] = uint8(j ^ i)
		}
	}
	return slidePadre
}

func main() {
	pic.Show(Pic)
}
