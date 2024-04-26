# Created by luocy at 2024/04/26 16:58
# leetgo: 1.4.5
# https://leetcode.com/problems/house-robber/

from typing import *

from leetgo_py import *


# @lc code=begin

def rob_normal(nums: list[int], i: int) -> int:
    if i == 0:
        return nums[0]
    if i == 1:
        return max(nums[0], nums[1])

    yes = rob_normal(nums, i - 1)
    no = rob_normal(nums, i - 2) + nums[i]
    return max(yes, no)


def rob_mem(nums: list[int], mem: list[int], i: int) -> int:
    if i == 0:
        return nums[0]
    if i == 1:
        return max(nums[0], nums[1])

    if mem[i]:
        return mem[i]

    yes = rob_mem(nums, mem, i - 1)
    no = rob_mem(nums, mem, i - 2) + nums[i]
    mem[i] = max(yes, no)
    return mem[i]


def rob_dp(nums: list[int]) -> int:
    n = len(nums)
    dp = [0] * n
    dp[0], dp[1] = nums[0], max(nums[0], nums[1])

    for i in range(2, n):
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
    return dp[n - 1]


def rob_dp_comp(nums: list[int]) -> int:
    n = len(nums)
    if n < 2:
        return nums[0]
    x, y = nums[0], max(nums[0], nums[1])

    for i in range(2, n):
        x, y = y, max(y, x + nums[i])
    return y


class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        pre, cur = 0, 0

        for i in range(n):
            pre, cur = cur, max(cur, pre + nums[i])
        return cur


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    ans = Solution().rob(nums)
    print("\noutput:", serialize(ans, "integer"))
