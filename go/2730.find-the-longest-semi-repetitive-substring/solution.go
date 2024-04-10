// Created by luocy at 2024/04/10 10:43
// leetgo: 1.4.5
// https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestSemiRepetitiveSubstring(s string) (ans int) {
	return longestSemiRepetitiveSubstring1(s)
}

func longestSemiRepetitiveSubstring1(s string) (ans int) {
	left, duplicate := 0, 0
	for right, x := range s {
		if right > 0 && x == rune(s[right-1]) {
			if duplicate > 0 {
				left = duplicate
			}
			duplicate = right
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++
	}
	return
}

func longestSemiRepetitiveSubstring2(s string) (ans int) {
	ans++
	left, duplicate := 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			duplicate++
			if duplicate > 1 {
				left++
				for s[left] != s[left-1] {
					left++
				}
				duplicate = 1
			}
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestSemiRepetitiveSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
