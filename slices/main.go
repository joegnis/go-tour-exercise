package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ay := make([][]uint8, dy)
	for j := range ay {
		ay[j] = make([]uint8, dx)
		for i := range ay[j] {
			ay[j][i] = uint8(i ^ j)
		}
	}
	return ay
}

func main() {
	pic.Show(Pic)
}
