// https://atcoder.jp/contests/abc455/tasks/abc455_b
package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	ss := make([]string, h)
	for i := range h {
		fmt.Scan(&ss[i])
	}

	ans := 0
	// 上端 h1・下端 h2・左端 w1・右端 w2 を固定して長方形領域を全探索する
	for h1 := 0; h1 < h; h1++ {
		for h2 := h1 + 1; h2 <= h; h2++ {
			for w1 := 0; w1 < w; w1++ {
				for w2 := w1 + 1; w2 <= w; w2++ {
					// 点対称かどうかを判定する
					flag := true
					for i := h1; i < h2 && flag; i++ {
						for j := w1; j < w2; j++ {
							if ss[i][j] != ss[h1+h2-i-1][w1+w2-j-1] {
								flag = false
								break
							}
						}
					}
					if flag {
						ans++
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
