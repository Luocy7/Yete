// Created by luocy at 2024/03/27 22:00
// leetgo: 1.4.3
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(numbers []int, target int) (ans []int) {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			ans = []int{l + 1, r + 1}
			return
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numbers := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(numbers, target)

	fmt.Println("\noutput:", Serialize(ans))
}
