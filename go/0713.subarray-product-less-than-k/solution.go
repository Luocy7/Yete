// Created by luocy at 2023/04/11 09:43
// https://leetcode.com/problems/subarray-product-less-than-k/

/*
713. Subarray Product Less Than K (Medium)
Given an array of integers `nums` and an integer `k`, return the number of contiguous subarrays
where the product of all the elements in the subarray is strictly less than  `k`.

**Example 1:**

```
Input: nums = [10,5,2,6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

```

**Example 2:**

```
Input: nums = [1,2,3], k = 0
Output: 0

```

**Constraints:**

- `1 <= nums.length <= 3 * 10⁴`
- `1 <= nums[i] <= 1000`
- `0 <= k <= 10⁶`

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	left, right, prod := 0, 0, 1
	for right < len(nums) {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
		right++
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := numSubarrayProductLessThanK(nums, k)
	fmt.Println("output: " + Serialize(ans))
}
