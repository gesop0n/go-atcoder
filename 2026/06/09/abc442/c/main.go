// https://atcoder.jp/contests/abc442/tasks/abc442_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// c[i] := 研究者 i と利害関係に「ない」研究者の人数
	// 初期値は自分以外の全員 n-1。利害関係のある相手が現れるたびに減らす。
	c := make([]int, n)
	for i := range c {
		c[i] = n - 1
	}
	for range m {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		c[a-1]--
		c[b-1]--
	}

	// 求める答えは C(c[i], 3) = c(c-1)(c-2)/6
	for i := range n {
		ci := c[i]
		ans := ci * (ci - 1) * (ci - 2) / 6
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(ans))
	}
	writer.WriteByte('\n')
}
