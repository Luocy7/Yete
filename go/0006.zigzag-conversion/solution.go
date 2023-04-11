// Created by luocy at 2023/03/30 13:54
// https://leetcode.com/problems/zigzag-conversion/

/*
6. Zigzag Conversion (Medium)
The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R

```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);

```

**Example 1:**

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

```

**Example 2:**

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

```

**Example 3:**

```
Input: s = "A", numRows = 1
Output: "A"

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consists of English letters (lower-case and upper-case), `','` and `'.'`.
- `1 <= numRows <= 1000`

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	var rows = make([]string, numRows)
	index, step := 0, 1
	for _, c := range s {
		rows[index] += string(c)
		if index == 0 {
			step = 1
		} else if index == numRows-1 { // turn around at the bottom
			step = -1
		}
		index += step
	}

	var result string
	for _, row := range rows {
		result += row
	}
	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	numRows := Deserialize[int](ReadLine(stdin))
	ans := convert(s, numRows)
	fmt.Println("output: " + Serialize(ans))
}
