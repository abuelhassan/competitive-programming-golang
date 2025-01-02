// Given a stream of characters, find the first non-repeating character from stream.
// You need to tell the first non-repeating character in O(1) time at any moment.
package main

import (
	"container/list"
	"fmt"
	"strings"
)

func firstNonRepeating(str string) string {
	repeated := map[rune]struct{}{}
	ptr := map[rune]*list.Element{}
	dll := list.New()

	var bld strings.Builder
	for _, c := range str {
		el := ptr[c]
		if el == nil { // first occurrence
			ptr[c] = dll.PushBack(c)
		} else if _, ok := repeated[c]; !ok { // second occurrence
			repeated[c] = struct{}{}
			dll.Remove(el)
		}

		if fr := dll.Front(); fr != nil {
			bld.WriteRune(fr.Value.(rune))
		} else { // all repeated
			bld.WriteRune('#')
		}
	}
	return bld.String()
}

func main() {
	tt := []struct {
		in  string
		out string
	}{
		{
			in:  "abadbc",
			out: "aabbdd",
		},
		{
			in:  "abcabc",
			out: "aaabc#",
		},
		{
			in:  "aabcdbcdee",
			out: "a#bbbcd#e#",
		},
	}
	for _, t := range tt {
		if out := firstNonRepeating(t.in); out != t.out {
			panic(fmt.Sprintf("\ninput: %v\noutput: %v\nexpected: %v\n", t.in, out, t.out))
		}
	}
	fmt.Println("Todo bien!")
}
