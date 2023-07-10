package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	_, _ = fmt.Fscan(r, &n, &m)

	lg := make([]int, n+1)
	for i := 2; i <= n; i++ {
		lg[i] = lg[i/2] + 1
	}

	st := make([][]int, lg[n]+1)
	st[0] = make([]int, n)
	for j := 0; j < n; j++ {
		_, _ = fmt.Fscan(r, &st[0][j])
	}
	for i := 1; i <= lg[n]; i++ {
		st[i] = make([]int, n)
		for j := 0; j+(1<<i) <= n; j++ {
			st[i][j] = max(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	ans := 0
	for k := 0; k < m; k++ {
		var i, j int
		_, _ = fmt.Fscan(r, &i, &j)
		j--
		pwr := lg[j]
		if max(st[pwr][i-1], st[pwr][j-(1<<pwr)]) == st[0][i-1] {
			ans++
		}
	}

	_, _ = fmt.Fprintf(w, "%d\n", ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
