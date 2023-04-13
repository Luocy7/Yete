// Created by luocy at 2023/04/11 14:30
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

/*
1658. Minimum Operations to Reduce X to Zero (Medium)
You are given an integer array `nums` and an integer `x`. In one operation, you can either remove
the leftmost or the rightmost element from the array `nums` and subtract its value from `x`. Note
that this **modifies** the array for future operations.

Return the **minimum number** of operations to reduce  `x`to **exactly** `0`if it is possible,
otherwise, return  `-1`.

**Example 1:**

```
Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.

```

**Example 2:**

```
Input: nums = [5,6,7,8,9], x = 4
Output: -1

```

**Example 3:**

```
Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5
operations in total) to reduce x to zero.

```

**Constraints:**

- `1 <= nums.length <= 10⁵`
- `1 <= nums[i] <= 10⁴`
- `1 <= x <= 10⁹`

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minOperations(nums []int, x int) (ans int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	left, right, cur, max := 0, 0, 0, 0
	for right < len(nums) {
		cur += nums[right]
		right++
		for cur > target && left < right {
			cur -= nums[left]
			left++
		}
		if cur == target {
			max = Max(max, right-left)
		}
	}
	if max == 0 {
		return -1
	}
	return len(nums) - max
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
	x := Deserialize[int](ReadLine(stdin))
	ans := minOperations(nums, x)
	fmt.Println("output: " + Serialize(ans))
}
