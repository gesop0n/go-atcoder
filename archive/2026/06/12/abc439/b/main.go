// https://atcoder.jp/contests/abc439/tasks/abc439_b
package main

import (
	"fmt"
)

// f は n の各桁の2乗の和を返す。
func f(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

func isHappyNumber(n int) bool {
	seen := map[int]bool{}
	for n != 1 {
		if seen[n] {
			// 一度通った数に戻ってきた = ループ
			return false
		}
		seen[n] = true
		n = f(n)
	}
	return true
}

// 別解: Floyd の循環検出（うさぎとかめ）。
// n -> f(n) の関数グラフ上の閉路検出。map 不要でメモリ O(1)。
// slow と fast は必ず閉路内で合流し、合流点が 1（自己ループ）なら Yes。
//
// func isHappyNumber(n int) bool {
// 	slow, fast := n, f(n)
// 	for slow != fast {
// 		slow = f(slow)    // 1歩ずつ
// 		fast = f(f(fast)) // 2歩ずつ
// 	}
// 	return slow == 1
// }

func main() {
	var n int
	fmt.Scan(&n)

	if isHappyNumber(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
