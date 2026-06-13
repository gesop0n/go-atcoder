// https://atcoder.jp/contests/abc436/tasks/abc436_a
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	forward := strings.Repeat("o", n-len(s))
	fmt.Println(forward + s)
}
