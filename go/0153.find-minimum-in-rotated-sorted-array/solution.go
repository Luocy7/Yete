// Created by luocy at 2024/04/07 10:00
// leetgo: 1.4.3
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMin(nums []int) (ans int) {
	return findMin1(nums)
}

func findMin1(nums []int) (ans int) {
	// if nums[0] > nums[-1] the nums is must be rotated

	left, right := -1, len(nums)-1 // open interval (-1, n-1); close interval [0, n-2]
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] { // nums[mid] is k or in the right of k
			right = mid
		} else { // think: when does the left num bigger then the right num?
			left = mid
		}
	}
	return nums[right]
}

func findMin2(nums []int) (ans int) {
	left, right, l := 0, len(nums)-2, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[l] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findMin(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
