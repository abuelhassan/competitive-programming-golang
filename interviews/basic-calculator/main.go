package main

import (
	"fmt"
	"unicode"
)

type stack []interface{}

// exp is always valid, it doesn't have negative numbers (e.g."-3")
// and it consists of digits, *, /, +, -, ( and/or ).
func calculate(exp string) int {
	ans := 0
	numsSt, opsSt := make(stack, 0), make(stack, 0)
	for i := 0; i < len(exp); i++ {
		switch exp[i] {
		case '*', '/', '+', '-', '(':
			opsSt.push(exp[i])
			continue
		case ')':
			num := 0
			for opsSt.top().(byte) != '(' {
				switch opsSt.pop().(byte) {
				case '+':
					num += numsSt.pop().(int)
				case '-':
					num -= numsSt.pop().(int)
				}
			}
			num += numsSt.pop().(int)
			numsSt.push(num)
			opsSt.pop()
		default: // digit
			num := 0
			for i < len(exp) && unicode.IsDigit(rune(exp[i])) {
				num = num*10 + int(exp[i]-'0')
				i++
			}
			numsSt.push(num)
			i--
		}
		if len(opsSt) == 0 || len(numsSt) < 2 {
			continue
		}
		switch opsSt.top().(byte) {
		case '+', '-':
			continue
		case '*':
			opsSt.pop()
			numsSt.push(numsSt.pop().(int) * numsSt.pop().(int))
		case '/':
			opsSt.pop()
			b, a := numsSt.pop().(int), numsSt.pop().(int)
			numsSt.push(a / b)
		}
	}
	for len(opsSt) > 0 {
		switch opsSt.pop().(byte) {
		case '+':
			ans += numsSt.pop().(int)
		case '-':
			ans -= numsSt.pop().(int)
		}
	}
	ans += numsSt.pop().(int)
	return ans
}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}
func (s *stack) pop() interface{} {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}
func (s *stack) top() interface{} {
	return (*s)[len(*s)-1]
}

func main() {
	tt := []struct {
		in  string
		out int
	}{
		{
			in:  "0",
			out: 0,
		},
		{
			in:  "1+1",
			out: 2,
		},
		{
			in:  "6-4/2*2",
			out: 2,
		},
		{
			in:  "2*(5+5*2)/3+(6/2+8)",
			out: 21,
		},
		{
			in:  "(2+6*3+5-(3*14/7+2)*5)+3",
			out: -12,
		},
	}
	for _, t := range tt {
		if out := calculate(t.in); out != t.out {
			panic(fmt.Sprintf("\ninput: %v\noutput: %v\nexpected: %v\n", t.in, out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
