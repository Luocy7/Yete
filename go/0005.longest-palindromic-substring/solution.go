// Created by Bob at 2023/03/30 11:48
// https://leetcode.com/problems/longest-palindromic-substring/

/*
5. Longest Palindromic Substring (Medium)
Given a string `s`, return the longestpalindromicsubstring in `s`.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// DP Time: O(n^2) Space: O(n^2)
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	// Keep track of the longest palindrome found so far; Like siding window
	start, maxLen := 0, 1
	// Check all substrings of length 2 or more
	for l := 2; l <= n; l++ { // l is the length of the substring
		for i := 0; i <= n-l; i++ { // i is the start index of the substring
			j := i + l - 1 // j is the end index of the substring
			if s[i] == s[j] && (l == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if l > maxLen {
					start = i
					maxLen = l
				}
			}
		}
	}

	// Return the longest palindrome found
	return s[start : start+maxLen]
}

// Center Expand Time: O(n^2) Space: O(1)
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)   // odd length
		res = maxPalindrome(s, i, i+1, res) // even length
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestPalindrome(s)
	fmt.Println("output: " + Serialize(ans))
}
