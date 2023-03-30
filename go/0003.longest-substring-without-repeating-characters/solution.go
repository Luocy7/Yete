// Created by Bob at 2023/03/30 09:54
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
3. Longest Substring Without Repeating Characters (Medium)
Given a string `s`, find the length of the **longest** **substring** without repeating characters.

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

```

**Constraints:**

- `0 <= s.length <= 5 * 10⁴`
- `s` consists of English letters, digits, symbols and spaces.

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	return lengthOfLongestSubstring2(s)
}

// lengthOfLongestSubstring1 Brute Force
func lengthOfLongestSubstring1(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}
	return
}

// lengthOfLongestSubstring2 Sliding Window
func lengthOfLongestSubstring2(s string) (ans int) {
	n := len(s)
	m := make(map[byte]int)  // store the newest postion(start from 1) of a character
	i := 0                   // left bound of the sliding window
	for j := 0; j < n; j++ { // right bound of the sliding window
		if _, ok := m[s[j]]; ok {
			i = Max(m[s[j]], i)
		}
		ans = Max(ans, j-i+1) // `j-i+1` is the length of the sliding window
		m[s[j]] = j + 1       // `j + 1` means newest postion(start from 1) of a character
	}
	return
}

// allUnique returns true if all characters in s[start:end] are unique
func allUnique(s string, start, end int) bool {
	m := make(map[byte]bool)
	for i := start; i < end; i++ {
		if m[s[i]] {
			return false
		}
		m[s[i]] = true
	}
	return true
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)
	fmt.Println("output: " + Serialize(ans))
}
