// Created by luocy at 2024/03/27 22:01
// leetgo: 1.4.3
// https://leetcode.com/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	left, right, sum := 0, 0, 0
	for right < len(nums) {
		right++
		sum += nums[right-1]
		for sum >= target {
			if ans == 0 || right-left < ans {
				ans = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
