// Created by luocy at 2024/03/29 16:36
// leetgo: 1.4.3
// https://leetcode.com/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums) // QuickSort
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // when nums[i] == nums[i-1], skip duplicate num
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // x plus any 2 nums after nums[i+2] will > 0, so break
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // x plus biggest two nums will < 0, so add i
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
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
