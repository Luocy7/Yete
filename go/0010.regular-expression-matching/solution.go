// Created by luocy at 2024/03/27 21:58
// leetgo: 1.4.3
// https://leetcode.com/problems/regular-expression-matching/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	p := Deserialize[string](ReadLine(stdin))
	ans := isMatch(s, p)

	fmt.Println("\noutput:", Serialize(ans))
}
