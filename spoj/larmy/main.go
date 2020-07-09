// (FIXME) Verdict: TLE.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// DP, Divide and Conquer Optimization
func rowDivideAndConquer(l, r, optL, optR int, dpBefore []int, dpCur []int, cost [][]int) {
	if l > r {
		return
	}

	var (
		mid = (l + r) >> 1
		ans = len(cost) * len(cost)
		opt = -1
	)
	for k := optL; k <= mid && k <= optR && k < len(cost)-1; k++ {
		t := dpBefore[k] + cost[mid][k+1]
		if t < ans {
			ans = t
			opt = k
		}
	}

	dpCur[mid] = ans
	rowDivideAndConquer(l, mid-1, optL, opt, dpBefore, dpCur, cost)
	rowDivideAndConquer(mid+1, r, opt, optR, dpBefore, dpCur, cost)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	_, _ = fmt.Fscan(r, &n, &m)
	row := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(r, &row[i])
	}

	cost := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		cost[i] = make([]int, n)
		for j := i - 1; j >= 0; j-- {
			cost[i][j] += cost[i][j+1]
			if row[i] < row[j] {
				cost[i][j]++
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			cost[i][j] += cost[i-1][j]
		}
	}

	memo := make([][]int, 2)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		memo[0][i] = cost[i][0]
	}
	for i := 1; i < m; i++ {
		rowDivideAndConquer(0, n-1, 0, n-1, memo[(i-1)%2], memo[i%2], cost)
	}
	_, _ = fmt.Fprintf(w, "%d\n", memo[(m-1)%2][n-1])
}
