// Created by luocy at 2024/03/27 21:23
// leetgo: 1.4.3
// https://leetcode.com/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	var mp = make(map[int]int)
	for i, num := range nums {
		if j, ok := mp[target-num]; ok {
			return []int{j, i}
		} else {
			mp[num] = i
		}
	}
	return nil
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
