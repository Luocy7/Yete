// Created by luocy at 2024/03/27 21:55
// leetgo: 1.4.3
// https://leetcode.com/problems/reverse-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverse(x int) (ans int) {
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if x > 1<<31-1 || x < -1<<31 {
		return 0
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := reverse(x)

	fmt.Println("\noutput:", Serialize(ans))
}
