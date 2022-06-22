package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n int
	fmt.Scan(&n)

	var t int64
	alice := ""
	voce := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		voce[i] = t
		if i == 0 {
			alice = alice + string(byte(bits.Len(uint(t))+96))
			continue
		}
		b := t - voce[i-1]
		if b < 0 {
			b = -b
		}
		sim := bits.Len(uint(b))
		if sim > 25 {
			alice = alice + " "
			continue
		}
		alice = alice + string(byte(sim+96))
	}
	fmt.Println(alice)
}
