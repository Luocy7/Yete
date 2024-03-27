// Created by luocy at 2024/03/27 22:03
// leetgo: 1.4.3
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minOperations(nums []int, x int) (ans int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	left, right, cur, max := 0, 0, 0, 0
	for right < len(nums) {
		cur += nums[right]
		right++
		for cur > target && left < right {
			cur -= nums[left]
			left++
		}
		if cur == target {
			max = Max(max, right-left)
		}
	}
	if max == 0 {
		return -1
	}
	return len(nums) - max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := minOperations(nums, x)

	fmt.Println("\noutput:", Serialize(ans))
}
