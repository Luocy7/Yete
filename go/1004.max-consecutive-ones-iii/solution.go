// Created by luocy at 2024/03/27 22:02
// leetgo: 1.4.3
// https://leetcode.com/problems/max-consecutive-ones-iii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestOnes(nums []int, k int) (ans int) {
	left, right := 0, 0
	var stack []int
	for right < len(nums) {
		if nums[right] == 0 {
			stack = append(stack, right)
		}
		if len(stack) > k {
			left = stack[0] + 1
			stack = stack[1:]
		}
		ans = Max(ans, right-left+1)
		right++
	}
	return
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
	k := Deserialize[int](ReadLine(stdin))
	ans := longestOnes(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
