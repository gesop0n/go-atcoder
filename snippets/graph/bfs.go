// Package graph contains copy-paste-ready graph algorithms for AtCoder.
package graph

// BFS は 0-indexed の隣接リストで与えられた無向/有向グラフについて、
// start から各頂点までの最短距離(辺の本数)を返す。
// start から到達できない頂点の距離は -1。
func BFS(adj [][]int, start int) []int {
	dist := make([]int, len(adj))
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, u := range adj[v] {
			if dist[u] == -1 {
				dist[u] = dist[v] + 1
				queue = append(queue, u)
			}
		}
	}
	return dist
}
