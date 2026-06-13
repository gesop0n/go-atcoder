// https://atcoder.jp/contests/abc450/tasks/abc450_b
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	costs := make([][]int, n-1)
	for i := range n - 1 {
		costs[i] = make([]int, n)
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Fscan(reader, &costs[i][j])
		}
	}

	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			for c := b + 1; c < n; c++ {
				if costs[a][c] > costs[a][b]+costs[b][c] {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
}
