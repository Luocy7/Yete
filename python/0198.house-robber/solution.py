# Created by luocy at 2024/04/26 16:58
# leetgo: 1.4.5
# https://leetcode.com/problems/house-robber/

from typing import *

from leetgo_py import *


# @lc code=begin

class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 2:
            return nums[0]
        x, y = nums[0], max(nums[0], nums[1])

        for i in range(2, n):
            x, y = y, max(y, x + nums[i])
        return y


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    ans = Solution().rob(nums)
    print("\noutput:", serialize(ans, "integer"))
