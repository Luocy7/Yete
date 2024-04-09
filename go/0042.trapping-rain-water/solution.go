// Created by luocy at 2024/04/09 14:11
// leetgo: 1.4.5
// https://leetcode.com/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {
	return trap2(height)
}

func trap1(height []int) (ans int) {
	left, right := height[0], height[len(height)-1]
	leftBorder, rightBorder := []int{left}, make([]int, len(height))
	rightBorder[len(rightBorder)-1] = right

	for _, x := range height[1:] {
		if x > left {
			left = x
		}
		leftBorder = append(leftBorder, left)
	}

	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > right {
			right = height[i]
		}
		rightBorder[i] = right
	}

	for i := 0; i <= len(height)-1; i++ {
		ans += Min(leftBorder[i], rightBorder[i]) - height[i]
	}
	return
}

func trap2(height []int) (ans int) {
	var leftMax, rightMax int
	i, j := 0, len(height)-1
	for i < j {
		leftMax = Max(leftMax, height[i])
		rightMax = Max(rightMax, height[j])

		if leftMax < rightMax {
			ans += leftMax - height[i]
			i++
		} else {
			ans += rightMax - height[j]
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
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
