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
	io := newFastIO()
	defer io.Flush()

	isPrime := make([]bool, 1000001)
	for i := 2; i < len(isPrime); i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < len(isPrime); i++ {
		if isPrime[i] {
			for j := i * i; j < len(isPrime); j += i {
				isPrime[j] = false
			}
		}
	}
	primesCnt := make([]int, len(isPrime))
	for i, cnt := 2, 0; i < len(isPrime); i++ {
		if isPrime[i] {
			cnt++
		}
		primesCnt[i] = cnt
	}

	var t int
	io.Read(&t)
	for caseID := 1; caseID <= t; caseID++ {
		var n, k int
		io.Read(&n, &k)
		primes := make([]int, k)
		for i := 0; i < k; i++ {
			io.Read(&primes[i])
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
		io.Write("Case %d: %d\n", caseID, ans)
	}
}

type fastIO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func newFastIO() fastIO {
	return fastIO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

func (io *fastIO) Read(args ...interface{}) {
	_, _ = fmt.Fscan(io.r, args...)
}

func (io *fastIO) Write(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(io.w, format, args...)
}

func (io *fastIO) Flush() {
	_ = io.w.Flush()
}
