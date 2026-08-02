// https://leetcode.com/problems/valid-anagram/description/?envType=problem-list-v2&envId=hash-table
// 242

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Run())
}

func Run() bool {
	s := "anagram"
	t := "nagaram"

	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s {
		_, ok := m[v]
		if ok {
			continue
		} else {
			count := strings.Count(s, string(v))
			m[v] = count
		}
	}

	for _, v := range t {
		_, ok := m2[v]
		if ok {
			continue
		} else {
			count := strings.Count(t, string(v))
			m2[v] = count
		}
	}

	for k, z := range m {
		fmt.Println(k, z)
		v, ok := m2[k]
		fmt.Println(v)
		if !ok {
			return false
		} else {
			fmt.Println(z, v)
			if z != v {
				return false
			}
		}
	}

	/* for _, v := range t {
		_, ok := m[v]
		if !ok {
			return false
		} else {
			count := strings.Count(t, string(v))
			if m[v] != count {
				return false
			}
		}
	} */

	return true
}
