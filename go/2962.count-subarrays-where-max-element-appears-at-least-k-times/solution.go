// Created by luocy at 2024/04/10 11:22
// leetgo: 1.4.5
// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countSubarrays(nums []int, k int) (ans int64) {
	n := len(nums)
	if n < k {
		return
	}
	mx := slices.Max(nums)
	left, counter := 0, 0
	for _, x := range nums {
		if x == mx {
			counter++
		}
		for counter == k {
			if nums[left] == mx {
				counter--
			}
			left++
		}
		ans += int64(left) // anchor the right
		/*
			for nums=[1, 3, 2, 3, 3, 1] k=2

			[3, 2, 3]
			[1, 3, 2, 3]
			left=2, right=3

			[3, 3]
			[2, 3, 3]
			[3, 2, 3, 3]
			[1, 3, 2, 3, 3]
			left=4, right=4

			[3, 3, 1]
			[2, 3, 3, 1]
			[3, 2, 3, 3, 1]
			[1, 3, 2, 3, 3, 1]
			left=4, right=5

			total = 2+4+4
		*/
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := countSubarrays(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
