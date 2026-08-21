package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)
	/* fmt.Println(numArray.arr)
	fmt.Println(numArray.px) */
	fmt.Println(numArray.SumRange(0, 2), 1)  // вернёт (-2) + 0 + 3 = 1
	fmt.Println(numArray.SumRange(2, 5), -1) // вернёт 3 + (-5) + 2 + (-1) = -1
	fmt.Println(numArray.SumRange(0, 5), -3) // вернёт (-2) + 0 + 3 + (-5) + 2 + (-1) = -3 */
}

type NumArray struct {
	arr []int
	px  []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	na.arr = make([]int, len(nums))
	copy(na.arr, nums)
	makePrefix(&na)
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.px[right+1] - this.px[left]
}

func makePrefix(na *NumArray) {
	na.px = make([]int, len(na.arr)+1)
	na.px[0] = 0
	for i := 0; i < len(na.arr); i++ {
		na.px[i+1] = na.px[i] + na.arr[i]
	}
}
