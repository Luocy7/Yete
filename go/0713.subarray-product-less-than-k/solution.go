// Created by luocy at 2024/03/27 22:01
// leetgo: 1.4.3
// https://leetcode.com/problems/subarray-product-less-than-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	left, right, prod := 0, 0, 1
	for right < len(nums) {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
		right++
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := numSubarrayProductLessThanK(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
