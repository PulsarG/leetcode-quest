// https://leetcode.com/problems/find-the-highest-altitude/submissions/2127118716/
// 1732

package main

//import "fmt"

func main() {
	Run()
}

func Run() int {
	g := []int{-4, -3, -2, -1, 4, 3, 2}
	return largestAltitude(g)
}

func largestAltitude(g []int) int {
	res := make([]int, len(g)+1)
	res[0] = 0
	var max int
	for i, v := range g {
		res[i+1] = res[i] + v
		if res[i+1] > max {
			max = res[i+1]
		}
	}
	return max
}
