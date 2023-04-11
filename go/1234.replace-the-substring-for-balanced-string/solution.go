// Created by luocy at 2023/04/11 10:36
// https://leetcode.com/problems/replace-the-substring-for-balanced-string/

/*
1234. Replace the Substring for Balanced String (Medium)
You are given a string s of length `n` containing only four kinds of characters: `'Q'`, `'W'`,
`'E'`, and `'R'`.

A string is said to be **balanced** if each of its characters appears `n / 4` times where `n` is the
length of the string.

Return the minimum length of the substring that can be replaced with **any** other string of the
same length to make  `s`**balanced**. If s is already **balanced**, return `0`.

**Example 1:**

```
Input: s = "QWER"
Output: 0
Explanation: s is already balanced.

```

**Example 2:**

```
Input: s = "QQWE"
Output: 1
Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.

```

**Example 3:**

```
Input: s = "QQQW"
Output: 2
Explanation: We can replace the first "QQ" to "ER".

```

**Constraints:**

- `n == s.length`
- `4 <= n <= 10⁵`
- `n` is a multiple of `4`.
- `s` contains only `'Q'`, `'W'`, `'E'`, and `'R'`.

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func balancedString(s string) (ans int) {
	n, e := len(s), len(s)/4
	count := make(map[byte]int, 4)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}
	left, right, ans := 0, 0, n
	fmt.Println(count)
	for right < n {
		count[s[right]]--
		right++
		for left < n && count['Q'] <= e && count['W'] <= e && count['E'] <= e && count['R'] <= e {
			ans = Min(ans, right-left)
			count[s[left]]++
			left++
		}
	}
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := balancedString(s)
	fmt.Println("output: " + Serialize(ans))
}
