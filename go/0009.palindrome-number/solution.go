// Created by luocy at 2024/03/27 21:57
// leetgo: 1.4.3
// https://leetcode.com/problems/palindrome-number/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed, num := 0, x
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return x == reversed
}

// @lc code=end

// turn to string
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := isPalindrome(x)

	fmt.Println("\noutput:", Serialize(ans))
}
