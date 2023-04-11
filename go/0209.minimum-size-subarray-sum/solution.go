// Created by Bob at 2023/04/11 09:24
// https://leetcode.com/problems/minimum-size-subarray-sum/

/*
209. Minimum Size Subarray Sum (Medium)
Given an array of positive integers `nums` and a positive integer `target`, return the **minimal
length** of a subarray whose sum is greater than or equal to `target`. If there is no such subarray,
return `0` instead.

**Example 1:**

```
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

```

**Example 2:**

```
Input: target = 4, nums = [1,4,4]
Output: 1

```

**Example 3:**

```
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

```

**Constraints:**

- `1 <= target <= 10⁹`
- `1 <= nums.length <= 10⁵`
- `1 <= nums[i] <= 10⁴`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution of which the
time complexity is `O(n log(n))`.

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	left, right, sum := 0, 0, 0
	for right < len(nums) {
		right++
		sum += nums[right-1]
		for sum >= target {
			if ans == 0 || right-left < ans {
				ans = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)
	fmt.Println("output: " + Serialize(ans))
}
