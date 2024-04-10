// Created by luocy at 2024/04/10 10:00
// leetgo: 1.4.5
// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSubarrayLength(nums []int, k int) (ans int) {
	counter := make(map[int]int)
	i, j := 0, 0
	for i <= j && j < len(nums) {
		cu := nums[j]
		counter[cu]++
		for counter[cu] > k {
			counter[nums[i]]--
			i++
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
		j++
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := maxSubarrayLength(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
