// Created by luocy at 2024/03/27 21:47
// leetgo: 1.4.3
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

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
	m := make(map[byte]int)  // store the newest position(start from 1) of a character
	i := 0                   // left bound of the sliding window
	for j := 0; j < n; j++ { // right bound of the sliding window
		if _, ok := m[s[j]]; ok {
			i = Max(m[s[j]], i)
		}
		ans = Max(ans, j-i+1) // `j-i+1` is the length of the sliding window
		m[s[j]] = j + 1       // `j + 1` means newest position(start from 1) of a character
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

	fmt.Println("\noutput:", Serialize(ans))
}
