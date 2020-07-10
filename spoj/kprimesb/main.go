package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	const mxN = 1000001
	var isPrime = [mxN]bool{}
	for i := 2; i < mxN; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < mxN; i++ {
		if isPrime[i] {
			for j := i * i; j < mxN; j += i {
				isPrime[j] = false
			}
		}
	}
	primesCnt := [mxN]int{}
	for i, cnt := 2, 0; i < mxN; i++ {
		if isPrime[i] {
			cnt++
		}
		primesCnt[i] = cnt
	}

	var t int
	_, _ = fmt.Fscan(r, &t)
	for caseID := 1; caseID <= t; caseID++ {
		var n, k int
		_, _ = fmt.Fscan(r, &n, &k)
		primes := make([]int, k)
		for i := 0; i < k; i++ {
			_, _ = fmt.Fscan(r, &primes[i])
		}

		// The Inclusion-Exclusion Principle
		ans := n
		for i := 1; i < 1<<uint(k); i++ {
			popcount := 0
			l := 1
			for j, p := range primes {
				if i>>uint(j)&1 == 1 {
					popcount++
					l = lcm(l, p)
				}
			}
			if popcount&1 == 1 {
				l = -l
			}
			ans += n / l
		}
		if n >= 1 {
			ans = ans - 1 // 1 is not a composite number.
		}
		ans = ans - primesCnt[n]
		for i := 0; i < k; i++ {
			if primes[i] <= n { // this prime has been removed twice.
				ans++
			}
		}
		_, _ = fmt.Fprintf(w, "Case %d: %d\n", caseID, ans)
	}
}
