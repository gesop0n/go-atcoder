// https://atcoder.jp/contests/abc455/tasks/abc455_c
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	as := map[int]int{}
	for range n {
		var a int
		fmt.Scan(&a)
		as[a] += a
	}

	values := make([]int, 0, len(as))
	for _, v := range as {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// k番目以降の和
	sum := 0
	for i := k; i < len(values); i++ {
		sum += values[i]
	}
	fmt.Println(sum)
}
