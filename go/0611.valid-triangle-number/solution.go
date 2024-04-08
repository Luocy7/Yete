// Created by luocy at 2024/04/08 16:56
// leetgo: 1.4.5
// https://leetcode.com/problems/valid-triangle-number/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)

	for i := n - 1; i > 1; i-- {
		x := nums[i]
		if nums[i-1]+nums[i-2] < x {
			continue
		}
		if nums[0]+nums[1] > x {
			ans += (i - 1) * i / 2
			continue
		}

		j, k := 0, i-1
		for j < k {
			if nums[j]+nums[k] > x {
				ans += k - j
				k--
			} else {
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
	ans := triangleNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
