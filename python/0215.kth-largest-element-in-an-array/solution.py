# Created by luocy at 2024/05/05 21:15
# leetgo: 1.4.6
# https://leetcode.com/problems/kth-largest-element-in-an-array/

import random
from typing import *

from leetgo_py import *

# @lc code=begin


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        left, right, kth_index = 0, len(nums) - 1, len(nums) - k
        while True:
            pivot_index = random.randint(left, right)
            new_pivot_index = self.partition(nums, left, right, pivot_index)
            if new_pivot_index == kth_index:
                return nums[new_pivot_index]
            elif new_pivot_index > kth_index:
                right = new_pivot_index - 1
            else:
                left = new_pivot_index + 1

    def partition(self, nums, left, right, pivot_index):
        pivot = nums[pivot_index]
        nums[pivot_index], nums[right] = nums[right], nums[pivot_index]
        stored_index = left
        for i in range(left, right):
            if nums[i] < pivot:
                nums[i], nums[stored_index] = nums[stored_index], nums[i]
                stored_index += 1
        nums[right], nums[stored_index] = nums[stored_index], nums[right]
        return stored_index


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    k: int = deserialize("int", read_line())
    ans = Solution().findKthLargest(nums, k)
    print("\noutput:", serialize(ans, "integer"))
