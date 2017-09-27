package main

import (
	"fmt"
	"sort"
)

type Numbers []int

func (a Numbers) Len() int           { return len(a) }
func (a Numbers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Numbers) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return result
	}
	sort.Sort(Numbers(nums))

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] > nums[i-1] {
			j := i + 1
			k := len(nums) - 1

			for j < k {
				if nums[i]+nums[j]+nums[k] == 0 {
					l := make([]int, 0)
					l = append(l, nums[i])
					l = append(l, nums[j])
					l = append(l, nums[k])
					result = append(result, l)
					j++
					k--

					for j < k && nums[j] == nums[j-1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if nums[i]+nums[j]+nums[k] < 0 {
					j++
				} else {
					k--
				}
			}
		}
	}

	return result
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(input)
	fmt.Println(result)
}
