// Created by luocy at 2024/04/08 10:02
// leetgo: 1.4.5
// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countPairs(nums []int, target int) (ans int) {
	// `0 <= i < j < n`and `nums[i] + nums[j] < target` means no duplicate pairs like (3,4)&(4,3), so sort nums does not matter
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := countPairs(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
