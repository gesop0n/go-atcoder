// https://atcoder.jp/contests/abc463/tasks/abc463_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	takahashi := make([][]int, n)
	for i := range n {
		takahashi[i] = make([]int, 2)
	}

	for i := range n {
		var h, l int
		fmt.Fscan(reader, &h, &l)
		takahashi[i][0] = h
		takahashi[i][1] = l
	}

	sort.Slice(takahashi, func(i, j int) bool { return takahashi[i][0] > takahashi[j][0] })
	fmt.Println(takahashi)

	var q int
	fmt.Fscan(reader, &q)

	head := 0
	for range q {
		var t int
		fmt.Fscan(reader, &t)

		for {
			if takahashi[head][1] > t {
				break
			} else {
				head += 1
			}
		}

		fmt.Println(takahashi[head][0])
	}
}
