// Created by luocy at 2024/03/27 21:59
// leetgo: 1.4.3
// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxArea(height []int) (ans int) {
	var i, j = 0, len(height) - 1
	for i < j {
		ans = Max(ans, Min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
	height := Deserialize[[]int](ReadLine(stdin))
	ans := maxArea(height)

	fmt.Println("\noutput:", Serialize(ans))
}
