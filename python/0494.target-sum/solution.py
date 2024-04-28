# Created by luocy at 2024/04/27 15:30
# leetgo: 1.4.5
# https://leetcode.com/problems/target-sum/

from typing import *

from leetgo_py import *

# @lc code=begin

"""
set p is sum of num that add '+', s is sum of all nums

then t = (s - p) is sum of num that add '-'

p - t = p - (s - p) = target
2p - s = target
p = (s + target) / 2
so s + target must be even number

"""


class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        s = sum(nums) + target
        if s % 2 or s < 0:
            return 0
        s = s // 2
        # transfer to choice some of num from nums whose sum is s


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    target: int = deserialize("int", read_line())
    ans = Solution().findTargetSumWays(nums, target)
    print("\noutput:", serialize(ans, "integer"))
