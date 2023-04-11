// Created by luocy at 2023/03/30 09:35
// https://leetcode.com/problems/two-sum/

/*
1. Two Sum (Easy)
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such
that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same
element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]

```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]

```

**Constraints:**

- `2 <= nums.length <= 10⁴`
- `-10⁹ <= nums[i] <= 10⁹`
- `-10⁹ <= target <= 10⁹`
- **Only one valid answer exists.**

**Follow-up:** Can you come up with an algorithm that is less than `O(n²) ` time complexity?

*/

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
	fmt.Println("output: " + Serialize(ans))
}
