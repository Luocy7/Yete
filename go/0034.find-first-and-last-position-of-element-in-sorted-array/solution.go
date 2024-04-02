// Created by luocy at 2024/04/01 15:59
// leetgo: 1.4.3
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchRange(nums []int, target int) (ans []int) {
	l := len(nums)
	if l < 1 || nums[0] > target || nums[l-1] < target {
		return []int{-1, -1}
	}
	lower := lowerBound(nums, target)
	if lower == l || nums[lower] != target {
		return []int{-1, -1}
	}
	higher := lowerBound(nums, target+1) - 1
	return []int{lower, higher}
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // closed interval
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // bound reduce to [mid+1, right]
		} else {
			right = mid - 1 // bound reduce to [left, mid-1]
		}
	}
	return left
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
