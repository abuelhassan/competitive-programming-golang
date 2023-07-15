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

	var n int
	_, _ = fmt.Fscan(r, &n)

	logTable := make([]int, n+1)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i>>1] + 1
	}

	spTable := make([][]int, logTable[n]+1)
	spTable[0] = make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(r, &spTable[0][i])
	}
	for i := 1; i <= logTable[n]; i++ {
		spTable[i] = make([]int, n)
		for j := 0; j+(1<<i) <= n; j++ {
			spTable[i][j] = gcd(spTable[i-1][j], spTable[i-1][j+(1<<(i-1))])
		}
	}

	mp := make(map[int]int64)
	for i := 0; i < n; i++ {
		ptr, cGCD := i, spTable[0][i]
		for ptr < n {
			lwr, upr := ptr, n
			for lwr < upr {
				mid := lwr + ((upr - lwr) >> 1)
				pwr := logTable[mid-i]
				if gcd(spTable[pwr][i], spTable[pwr][mid-(1<<pwr)+1]) == cGCD {
					lwr = mid + 1
				} else {
					upr = mid
				}
			}
			mp[cGCD] += int64(lwr - ptr)
			ptr = lwr
			if lwr < n {
				cGCD = gcd(cGCD, spTable[0][lwr])
			}
		}
	}

	var q int
	_, _ = fmt.Fscan(r, &q)
	for i := 0; i < q; i++ {
		var x int
		_, _ = fmt.Fscan(r, &x)
		_, _ = fmt.Fprintf(w, "%d\n", mp[x])
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
