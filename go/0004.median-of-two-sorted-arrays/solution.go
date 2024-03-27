// Created by luocy at 2024/03/27 21:51
// leetgo: 1.4.3
// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		ans = float64(findKth(nums1, nums2, total/2+1))
	} else {
		ans = float64(findKth(nums1, nums2, total/2)+findKth(nums1, nums2, total/2+1)) / 2
	}
	return
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	i := min(len(nums1), k/2)
	j := k - i
	if nums1[i-1] < nums2[j-1] {
		return findKth(nums1[i:], nums2, j)
	} else if nums1[i-1] > nums2[j-1] {
		return findKth(nums1, nums2[j:], i)
	} else {
		return nums1[i-1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findMedianSortedArrays(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
