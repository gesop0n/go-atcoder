// https://atcoder.jp/contests/abc463/tasks/abc463_a
package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if 9*x == 16*y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
