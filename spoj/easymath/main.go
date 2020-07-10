package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int64) int64 {
	return a * (b / gcd(a, b))
}

func solve(n int64, factors []int64) int64 {
	// The Inclusion-Exclusion Principle
	ans := n
	for i := 1; i < 1<<uint(len(factors)); i++ {
		popcount := 0
		l := int64(1)
		for j, f := range factors {
			if i>>uint(j)&1 == 1 {
				popcount++
				l = lcm(l, f)
			}
		}
		if popcount&1 == 1 {
			l = -l
		}
		ans += n / l
	}
	return ans
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t int
	_, _ = fmt.Fscan(r, &t)
	for c := 0; c < t; c++ {
		var n, m, a, d int64
		_, _ = fmt.Fscan(r, &n, &m, &a, &d)
		factors := []int64{a, a + d, a + (d * 2), a + (d * 3), a + (d * 4)}
		_, _ = fmt.Println(solve(m, factors) - solve(n-1, factors))
	}
}
