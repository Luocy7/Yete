// Created by luocy at 2024/03/27 21:54
// leetgo: 1.4.3
// https://leetcode.com/problems/zigzag-conversion/

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

	fmt.Println("\noutput:", Serialize(ans))
}
