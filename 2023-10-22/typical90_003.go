package main

import "fmt"

// Graph グラフを表すデータ型
type Graph [][]int

// 頂点 s から DFS (ここではスタックを使う)
func dfs(G Graph, s int) []int {
	N := len(G)

	// 頂点 s からの距離
	dist := make([]int, N)
	for i := range dist {
		dist[i] = -1
	}
	dist[s] = 0

	// スタックで DFS
	st := []int{s}
	for len(st) > 0 {
		v := st[len(st)-1]
		st = st[:len(st)-1] // pop
		for _, nv := range G[v] {
			if dist[nv] == -1 {
				st = append(st, nv) // push
				dist[nv] = dist[v] + 1
			}
		}
	}

	// リターン
	return dist
}

func getAnswerByDFS() {
	// 入力
	var N int
	fmt.Scan(&N)

	// N 頂点のグラフを作る
	G := make(Graph, N)
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	// 頂点 0 から DFS
	dist0 := dfs(G, 0)

	// 距離最大の点を求める
	mx := -1
	mv := -1
	for v := 0; v < N; v++ {
		if mx < dist0[v] {
			mx = dist0[v]
			mv = v
		}
	}

	// 頂点 mv から DFS
	distmv := dfs(G, mv)

	// その最大値を求める
	mx = -1
	for v := 0; v < N; v++ {
		if mx < distmv[v] {
			mx = distmv[v]
		}
	}
	fmt.Println(mx + 1)
}

// 頂点 s から BFS
func bfs(G Graph, s int) []int {
	N := len(G)

	// 頂点 s からの距離
	dist := make([]int, N)
	for i := range dist {
		dist[i] = -1
	}
	dist[s] = 0

	// キューで BFS
	queue := []int{s}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:] // dequeue
		for _, nv := range G[v] {
			if dist[nv] == -1 {
				queue = append(queue, nv) // enqueue
				dist[nv] = dist[v] + 1
			}
		}
	}

	// リターン
	return dist
}

func getByBFS() {
	// 入力
	var N int
	fmt.Scan(&N)

	G := make(Graph, N)
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	// 頂点 0 から BFS
	dist0 := bfs(G, 0)

	// 距離最大の点を求める
	mx := -1
	mv := -1
	for v := 0; v < N; v++ {
		if mx < dist0[v] {
			mx = dist0[v]
			mv = v
		}
	}

	// 頂点 mv から BFS
	distmv := bfs(G, mv)

	// その最大値を求める
	mx = -1
	for v := 0; v < N; v++ {
		if mx < distmv[v] {
			mx = distmv[v]
		}
	}
	fmt.Println(mx + 1)
}
