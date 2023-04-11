// Created by Bob at 2023/03/31 18:57
// https://leetcode.com/problems/container-with-most-water/

/*
11. Container With Most Water (Medium)
You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that
the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the
most water.

Return the maximum amount of water a container can store.

**Notice** that you may not slant the container.

**Example 1:**

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
the max area of water (blue section) the container can contain is 49.

```

**Example 2:**

```
Input: height = [1,1]
Output: 1

```

**Constraints:**

- `n == height.length`
- `2 <= n <= 10⁵`
- `0 <= height[i] <= 10⁴`

*/

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
	fmt.Println("output: " + Serialize(ans))
}
