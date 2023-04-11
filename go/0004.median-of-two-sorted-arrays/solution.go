// Created by luocy at 2023/03/30 11:33
// https://leetcode.com/problems/median-of-two-sorted-arrays/

/*
4. Median of Two Sorted Arrays (Hard)
Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median**
of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

**Example 1:**

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

```

**Example 2:**

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

```

**Constraints:**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10⁶ <= nums1[i], nums2[i] <= 10⁶`

*/

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
	fmt.Println("output: " + Serialize(ans))
}
