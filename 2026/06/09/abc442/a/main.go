// https://atcoder.jp/contests/abc442/tasks/abc442_a
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := 0
	for _, c := range s {
		if c == 'i' || c == 'j' {
			ans += 1
		}
	}

	fmt.Println(ans)
}
