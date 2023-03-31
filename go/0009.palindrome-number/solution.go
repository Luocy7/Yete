// Created by Bob at 2023/03/30 20:39
// https://leetcode.com/problems/palindrome-number/

/*
9. Palindrome Number (Easy)
Given an integer `x`, return `true` if  `x` is a **palindrome**, and  `false` otherwise.

**Example 1:**

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

```

**Example 2:**

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is
not a palindrome.

```

**Example 3:**

```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

```

**Constraints:**

- `-2³¹ <= x <= 2³¹ - 1`

**Follow up:** Could you solve it without converting the integer to a string?

*/

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

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := isPalindrome(x)
	fmt.Println("output: " + Serialize(ans))
}
