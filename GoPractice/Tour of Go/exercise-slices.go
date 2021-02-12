package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//initialize columns
	pic := make([][]uint8, dy)
	for i := range pic {
		//initialize row
		pic[i] = make([]uint8, dx)
		
		//fill bits
		for j := range pic[i] {
			val := i^j//i*j//(i+j)/2
			pic[i][j] = uint8(val)
		}
	}

	return pic
}

func main() {
	Pic(2,2)
	pic.Show(Pic)
}
