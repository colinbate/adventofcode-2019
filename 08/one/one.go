package one

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
		pixel   int
		pval    int
		layer   int
		total   int
		best    [3]int
		counter [3]int
	)
	fmt.Println("Part One")
	best[0] = layerSize + 1
	for {
		n, err := fmt.Scanf("%c", &pixel)
		if n == 0 {
			fmt.Println(err)
			break
		}
		pval, _ = strconv.Atoi(string(pixel))
		counter[pval]++
		total++
		if total/layerSize != layer {
			if counter[0] < best[0] {
				best[0] = counter[0]
				best[1] = counter[1]
				best[2] = counter[2]
			}
			counter[0] = 0
			counter[1] = 0
			counter[2] = 0
			layer = total / layerSize
		}
	}
	fmt.Println("Best:", best[1]*best[2])
}
