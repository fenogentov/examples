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
	for i := 0; i < t; i++ {
		nabor := make(map[int]int, 0)
		var n int
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			var p int
			fmt.Fscan(in, &p)
			nabor[p]++
		}
		summ := 0
		for c, k := range nabor {
			summ += (k - k/3) * c
		}
		fmt.Fprintln(out, summ)
	}
}
