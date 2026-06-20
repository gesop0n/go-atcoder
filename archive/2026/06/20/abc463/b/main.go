// https://atcoder.jp/contests/abc463/tasks/abc463_b
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var x string

	fmt.Scan(&n, &x)

	index := []rune(x)[0] - 65

	for range n {
		var s string
		fmt.Fscan(reader, &s)
		if s[index] == 'o' {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
