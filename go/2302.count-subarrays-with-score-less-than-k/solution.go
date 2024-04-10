// Created by luocy at 2024/04/10 15:30
// leetgo: 1.4.5
// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countSubarrays(nums []int, k int64) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int64](ReadLine(stdin))
	ans := countSubarrays(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
