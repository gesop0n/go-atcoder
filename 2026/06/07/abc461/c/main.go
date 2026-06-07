// https://atcoder.jp/contests/abc461/tasks/abc461_c
//
// 【問題】
//
//	N 個の宝石から「ちょうど K 個」を選ぶ。選んだ宝石の色が「M 種類以上」
//	という条件のもとで、価値の合計の最大値を求める。
//
// 【解法（公式 解法1）】
//
//	宝石を「色の条件を満たすための M 個（印つき）」と「純粋に価値をかせぐ
//	残り K-M 個（印なし）」の2つの役割に分けて考える。
//
//	1. 色ごとに最高価値の1個を「代表」とする。
//	   （ある色を1つ入れるなら、その色の中で最高価値を入れるのが当然得）
//	2. 代表を価値の降順に並べ、上位 M 個に印をつける。
//	   → この M 個は全部色が違うので、選んだ時点で「M 種類以上」を達成。
//	3. 残り K-M 個は色を気にせず、印なしの宝石から価値の高い順に取る。
//	   （印なし = 代表だが上位 M に入らなかったもの + 各色の2番手以降）
//
//	(★) この「上位 M 代表を全部選ぶ」形の中に最適解が必ず存在する、という
//	     主張を認めると、上の貪欲で答えが求まる。ボトルネックはソートで
//	     計算量は O(N log N)。
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	// t[c] = 色 c に属する宝石の価値リスト（色ごとに分類）
	t := make([][]int, n+1)
	for range n {
		var c, v int
		fmt.Fscan(reader, &c, &v)
		t[c] = append(t[c], v)
	}

	top := []int{}  // 各色の代表（その色の最高価値）を集める
	tail := []int{} // 印なし候補プール（各色の2番手以降）

	for _, r := range t {
		if len(r) == 0 {
			continue // この色の宝石は存在しない
		}

		// その色の中を価値の降順にソート → 先頭 r[0] がこの色の代表
		sort.Sort(sort.Reverse(sort.IntSlice(r)))

		top = append(top, r[0])       // 代表（手順1）
		tail = append(tail, r[1:]...) // 2番手以降は印なしプールへ
	}

	// 代表どうしを価値の降順に並べる（手順2）。
	// → top[:m] が「印つき M 個」、top[m:] は上位 M に入れなかった代表
	sort.Sort(sort.Reverse(sort.IntSlice(top)))

	// 印に入れなかった代表も「印なし候補」として価値かせぎプールに加える
	tail = append(tail, top[m:]...)

	// 印なしプールを価値の降順に並べ、上位から K-M 個を取れるようにする（手順3）
	sort.Sort(sort.Reverse(sort.IntSlice(tail)))

	ans := 0
	for _, v := range top[:m] { // 色の条件を満たす印つき M 個
		ans += v
	}
	for _, v := range tail[:k-m] { // 価値をかせぐ残り K-M 個
		ans += v
	}

	fmt.Println(ans)
}
