// Created by luocy at 2024/03/27 21:56
// leetgo: 1.4.3
// https://leetcode.com/problems/string-to-integer-atoi/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myAtoi(s string) (ans int) {
	var sign int // when we first meet a `-`, `+` or a digit, we can determine the sign, when sign is determined,
	// the next any symbol will break the loop
	for _, c := range s {
		if sign == 0 && c == ' ' { // not determine sign, meet space
			continue
		} else if sign == 0 && c == '-' { // meet first minus
			sign = -1
			continue
		} else if sign == 0 && c == '+' { // meet first plus
			sign = 1
			continue
		} else if c >= '0' && c <= '9' {
			if sign == 0 { // meet first digit
				sign = 1
			}
			ans = ans*10 + int(c-'0')
			if ans > 1<<31-1 { // limit the max to 2^31
				if sign > 0 {
					return 1<<31 - 1
				}
				return -1 << 31
			}
		} else {
			break
		}
	}
	return sign * ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := myAtoi(s)

	fmt.Println("\noutput:", Serialize(ans))
}
