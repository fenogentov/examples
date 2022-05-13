package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type nums [3]int

type tr []nums

func (ms tr) Len() int {
	return len(ms)
}

func (ms tr) Less(i, j int) bool {
	if ms[i][0] < ms[j][0] {
		return true
	}
	if ms[i][1] < ms[j][1] {
		return true
	}
	if ms[i][2] < ms[j][2] {
		return true
	}

	return false
}

func (ms tr) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	var a string
	var err error

	fmt.Scan(&a)
	buf := strings.Split(a, ",")
	f := make([]int, len(buf))

	for i, c := range buf {
		f[i], err = strconv.Atoi(c)
		if err != nil {
			return
		}
	}

	var tree tr
	for w := 0; w < len(f)-3; w++ {
		for v := w + 1; v < len(f)-2; v++ {
			for q := w + 2; q < len(f)-1; q++ {
				if f[w]+f[v]+f[q] == 0 {
					var tmp []int = []int{f[w], f[v], f[q]}
					sort.Ints(tmp)
					h := false
					for _, t := range tree {
						if t[0] == tmp[0] && t[1] == tmp[1] && t[2] == tmp[2] {
							h = true
						}
					}
					if !h {
						tree = append(tree, nums{tmp[0], tmp[1], tmp[2]})
					}
				}
			}
		}
	}
	sort.Sort(tr(tree))

	for k, s := range tree {
		fmt.Printf("%d,%d,%d", s[0], s[1], s[2])
		if k < len(tree)-1 {
			fmt.Printf(" ")
		}
	}

	fmt.Println()
}
