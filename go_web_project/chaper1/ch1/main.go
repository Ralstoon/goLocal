package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum2([]int{7, 2, 15, 11}, 9))

}
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
