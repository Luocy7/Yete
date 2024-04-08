// Created by luocy at 2024/04/08 15:49
// leetgo: 1.4.5
// https://leetcode.com/problems/4sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fourSum(nums []int, target int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)
	if n < 4 {
		return
	}
	for i, x := range nums[:n-3] {
		if i > 0 && nums[i-1] == x {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			y := nums[j]
			if j > i+1 && nums[j-1] == y {
				continue
			}

			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}

			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}

			a, b := j+1, n-1
			for a < b {
				s := x + y + nums[a] + nums[b]
				if s < target {
					a++
				} else if s > target {
					b--
				} else {
					ans = append(ans, []int{x, y, nums[a], nums[b]})
					for a++; a < b && nums[a-1] == nums[a]; a++ { // reduce left boundary to skip duplicate third num
					}
					for b--; a < b && nums[b+1] == nums[b]; b-- { // same as above to skip duplicate fourth num
					}
				}
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
	ans := fourSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
