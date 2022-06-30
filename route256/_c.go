package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for g := 0; g < t; g++ {
		var n, m, el int
		fmt.Scan()
		fmt.Fscan(in, &n, &m)
		mtx := make([][]int, n)
		for f := 0; f < n; f++ {
			row := make([]int, m)
			for l := 0; l < m; l++ {
				fmt.Fscan(in, &el)
				row[l] = el
			}
			mtx[f] = row
		}

		var k, c int
		fmt.Fscan(in, &k)
		for h := 0; h < k; h++ {
			fmt.Fscan(in, &c)
			Bubblesort(mtx, c)
		}

		for _, st := range mtx {
			for u, v := range st {
				fmt.Printf("%d", v)
				if u == len(st)-1 {
					continue
				}
				fmt.Printf(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func Bubblesort(arr [][]int, column int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j][column-1] > arr[j+1][column-1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
