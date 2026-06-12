// https://atcoder.jp/contests/abc439/tasks/abc439_a
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := int(math.Pow(2, float64(n))) - 2*n
	fmt.Println(ans)
}
