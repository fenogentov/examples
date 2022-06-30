package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		note := make(map[string][]string, 0)
		var n int
		fmt.Scan(&n)
		var name, numb string
		for j := 0; j < n; j++ {
			fmt.Scan(&name, &numb)
			a, ok := note[name]
			if !ok {
				note[name] = []string{numb}
				continue
			}
			tmp := a
			for idx, nbr := range a {
				if nbr == numb {
					tmp = append(a[:idx], a[idx+1:]...)
				}
			}

			if len(a) > 4 {
				tmp = a[:4]
			}
			a = append([]string{numb}, tmp...)
			note[name] = a
		}

		keys := make([]string, 0, len(note))
		for k := range note {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {

			fmt.Printf("%s: %d ", k, len(note[k]))
			for r, s := range note[k] {
				fmt.Printf("%s", s)
				if r < len(s)-3 {
					fmt.Printf(" ")
					continue
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
