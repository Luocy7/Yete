# Created by luocy at 2024/04/26 21:38
# leetgo: 1.4.5
# https://leetcode.com/problems/combination-sum-iv/
from functools import cache
from typing import *

from leetgo_py import *

# @lc code=begin


# dp[i] = sum(dp[i - num] for num in nums if num <= i)
# dp[0] = 1 select none of nums


def combinationSum4_mem(nums: List[int], target: int) -> int:
    @cache
    def dfs(i: int) -> int:
        if i == 0:
            return 1
        return sum(dfs(i - x) for x in nums if x <= i)

    return dfs(target)


def combinationSum4_dp(nums: List[int], target: int) -> int:
    dp = [1] + [0] * target
    for i in range(1, target + 1):
        dp[i] = sum(dp[i - x] for x in nums if x <= i)
    return dp[target]


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        return combinationSum4_dp(nums, target)


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    target: int = deserialize("int", read_line())
    ans = Solution().combinationSum4(nums, target)
    print("\noutput:", serialize(ans, "integer"))
