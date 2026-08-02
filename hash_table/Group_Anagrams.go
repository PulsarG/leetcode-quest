// https://leetcode.com/problems/group-anagrams/description/?envType=problem-list-v2&envId=hash-table
// 49

package main

import (
	"fmt"
	//"strings"
)

func main() {
	fmt.Println(Run())
}

func Run() [][]string {
	data := make(map[[26]byte]int)
	res := make([][]string, 0)
	var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for w := range strs {
		list := [26]byte{}

		for letter := range strs[w] {
			list[strs[w][letter]-'a']++
		}

		if i, ok := data[list]; ok {
			res[i] = append(res[i], strs[w])
		} else {
			res = append(res, []string{strs[w]})

			data[list] = len(res) - 1
		}
	}

	return res
}
