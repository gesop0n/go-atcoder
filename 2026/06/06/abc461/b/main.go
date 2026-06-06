// https://atcoder.jp/contests/abc461/tasks/abc461_b
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	as := make([]int, n)
	bs := make([]int, n)
	for i := range n {
		fmt.Scan(&as[i])
	}
	for i := range n {
		fmt.Scan(&bs[i])
	}


	ans := "Yes"
	for i := range n {
		if bs[as[i]-1] != i+1 {
			ans = "No"
		} 
	}

	fmt.Println(ans)
}
