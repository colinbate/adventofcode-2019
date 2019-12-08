package two

import (
	"fmt"
	"strconv"
)

const width int = 25
const height int = 6
const layerSize int = width * height

// Run is the entry point for this solution.
func Run() {
	var (
		image []string
		pixel int
		pval  int
		total int
		index int
	)
	fmt.Println("Part One")
	image = make([]string, layerSize)
	for {
		n, err := fmt.Scanf("%c", &pixel)
		if n == 0 {
			fmt.Println(err)
			break
		}
		pval, _ = strconv.Atoi(string(pixel))
		index = total % layerSize

		if image[index] == "" {
			if pval == 0 {
				image[index] = "."
			} else if pval == 1 {
				image[index] = "#"
			}
		}
		total++
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(image[y*width+x])
		}
		fmt.Println()
	}
}
