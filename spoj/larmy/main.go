// (FIXME) Verdict: TLE.

package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 5001

var (
	row  = [maxN]int{}
	memo = [2][maxN]int{}
	cost = [maxN][maxN]int{}
)

// DP, Divide and Conquer Optimization
func rowDivideAndConquer(rowIdx, l, r, optL, optR int) {
	if l > r {
		return
	}

	var (
		mid = (l + r) >> 1
		ans = maxN * maxN
		opt = -1
	)
	for k := optL; k <= mid && k <= optR; k++ {
		t := memo[(rowIdx-1)%2][k] + cost[mid][k+1]
		if t < ans {
			ans = t
			opt = k
		}
	}

	memo[rowIdx%2][mid] = ans
	rowDivideAndConquer(rowIdx, l, mid-1, optL, opt)
	rowDivideAndConquer(rowIdx, mid+1, r, opt, optR)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	_, _ = fmt.Fscan(r, &n, &m)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(r, &row[i])
	}

	for i := n - 1; i >= 0; i-- {
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

	for i := 0; i < n; i++ {
		memo[0][i] = cost[i][0]
	}
	for i := 1; i < m; i++ {
		rowDivideAndConquer(i, 0, n-1, 0, n-1)
	}
	_, _ = fmt.Fprintf(w, "%d\n", memo[(m-1)%2][n-1])
}
