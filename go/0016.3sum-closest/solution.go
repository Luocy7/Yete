// Created by luocy at 2023/04/13 10:59
// https://leetcode.com/problems/3sum-closest/

/*
16. 3Sum Closest (Medium)
Given an integer array `nums` of length `n` and an integer `target`, find three integers in `nums`
such that the sum is closest to `target`.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

**Example 1:**

```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

```

**Example 2:**

```
Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

```

**Constraints:**

- `3 <= nums.length <= 500`
- `-1000 <= nums[i] <= 1000`
- `-10⁴ <= target <= 10⁴`

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSumClosest(nums []int, target int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := threeSumClosest(nums, target)
	fmt.Println("output: " + Serialize(ans))
}
