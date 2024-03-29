// Created by luocy at 2024/03/29 16:36
// leetgo: 1.4.3
// https://leetcode.com/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	if n < 3 {
		return
	}

	QuickSort(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return
}

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := right
	counter := left
	for i := left; i < right; i++ {
		if nums[i] < nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[pivot], nums[counter] = nums[counter], nums[pivot]
	return counter
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
