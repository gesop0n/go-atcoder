// https://atcoder.jp/contests/abc461/tasks/abc461_a
package main

import "fmt"

func main() {
	var a, d int
	fmt.Scan(&a, &d)

	if a <= d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
