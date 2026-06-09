// https://atcoder.jp/contests/abc442/tasks/abc442_b
package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	volume := 0
	isPlaying := false
	for range q {
		var a int
		fmt.Scan(&a)

		if a == 1 {
			volume += 1
		} else if a == 2 {
			if volume >= 1 {
				volume -= 1
			}
		} else {
			isPlaying = !isPlaying
		}

		if volume >= 3 && isPlaying {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
