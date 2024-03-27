// Created by luocy at 2024/03/27 22:02
// leetgo: 1.4.3
// https://leetcode.com/problems/replace-the-substring-for-balanced-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func balancedString(s string) (ans int) {
	n, e := len(s), len(s)/4
	count := make(map[byte]int, 4)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}
	left, right, ans := 0, 0, n
	fmt.Println(count)
	for right < n {
		count[s[right]]--
		right++
		for left < n && count['Q'] <= e && count['W'] <= e && count['E'] <= e && count['R'] <= e {
			ans = Min(ans, right-left)
			count[s[left]]++
			left++
		}
	}
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := balancedString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
