package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m)(^[-])|([^0-9A-Za-z_-])`)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		users := make(map[string]struct{}, 0)
		var n int
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			var log string
			fmt.Fscan(in, &log)
			if len(log) < 2 || len(log) > 24 {
				fmt.Fprintln(out, "NO")
				continue
			}
			if re.MatchString(log) {
				fmt.Fprintln(out, "NO")
				continue
			}
			l := strings.ToLower(log)
			if _, ok := users[l]; ok {
				fmt.Fprintln(out, "NO")
				continue
			}
			users[l] = struct{}{}
			fmt.Fprintln(out, "YES")
		}
		fmt.Fprintln(out)
	}
}
