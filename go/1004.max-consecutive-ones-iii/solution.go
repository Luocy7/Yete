// Created by luocy at 2023/04/11 10:13
// https://leetcode.com/problems/max-consecutive-ones-iii/

/*
1004. Max Consecutive Ones III (Medium)
Given a binary array `nums` and an integer `k`, return the maximum number of consecutive  `1`'s in
the array if you can flip at most `k` `0`'s.

**Example 1:**

```
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
```

**Example 2:**

```
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

```

**Constraints:**

- `1 <= nums.length <= 10⁵`
- `nums[i]` is either `0` or `1`.
- `0 <= k <= nums.length`

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestOnes(nums []int, k int) (ans int) {
	left, right := 0, 0
	var stack []int
	for right < len(nums) {
		if nums[right] == 0 {
			stack = append(stack, right)
		}
		if len(stack) > k {
			left = stack[0] + 1
			stack = stack[1:]
		}
		ans = Max(ans, right-left+1)
		right++
	}
	return
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
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := longestOnes(nums, k)
	fmt.Println("output: " + Serialize(ans))
}
