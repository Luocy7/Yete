// Created by luocy at 2024/04/08 10:33
// leetgo: 1.4.5
// https://leetcode.com/problems/3sum-closest/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSumClosest(nums []int, target int) (ans int) {
	n := len(nums)
	if n == 3 {
		ans = nums[0] + nums[1] + nums[2]
		return
	}
	sort.Ints(nums)
	minDiff := math.MaxInt

	for i, x := range nums[:n-2] { // i in [0, n-3]
		if i > 0 && x == nums[i-1] { // skip duplicate num
			continue
		}

		s := x + nums[i+1] + nums[i+2] // optimize1: when sum of the first three nums is greater than target, all other three nums will > target, just compare the diff and break the loop
		if s > target {
			if s-target < minDiff {
				ans = s
			}
			break
		}

		s = x + nums[n-2] + nums[n-1] // optimize2: when sum of x and the last two nums is less than target, this x with any other two nums will < target, so compare the diff and go to next x
		if s < target {
			if target-s < minDiff {
				minDiff = target - s
				ans = s
			}
			continue
		}

		j, k := i+1, n-1 // two points
		for j < k {      // [j, k]
			s = x + nums[j] + nums[k]
			if s == target {
				return target
			}

			if s > target {
				if s-target < minDiff {
					minDiff = s - target
					ans = s
				}
				k--
			} else {
				if target-s < minDiff {
					minDiff = target - s
					ans = s
				}
				j++
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := threeSumClosest(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
