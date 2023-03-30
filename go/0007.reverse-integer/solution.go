// Created by Bob at 2023/03/30 15:04
// https://leetcode.com/problems/reverse-integer/

/*
7. Reverse Integer (Medium)
Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the
value to go outside the signed 32-bit integer range `[-2³¹, 2³¹ - 1]`, then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

**Example 1:**

```
Input: x = 123
Output: 321

```

**Example 2:**

```
Input: x = -123
Output: -321

```

**Example 3:**

```
Input: x = 120
Output: 21

```

**Constraints:**

- `-2³¹ <= x <= 2³¹ - 1`

*/

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
	fmt.Println("output: " + Serialize(ans))
}
