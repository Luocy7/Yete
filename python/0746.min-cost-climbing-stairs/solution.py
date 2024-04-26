# Created by luocy at 2024/04/26 21:17
# leetgo: 1.4.5
# https://leetcode.com/problems/min-cost-climbing-stairs/

from typing import *

from leetgo_py import *

# @lc code=begin

# dp[i] = min(dp[i - 2] + cost[i - 2], dp[i - 1] + cost[i - 1])
# dp[0] = dp[1] = 0


class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n = len(cost)
        pre, cur = 0, 0
        for i in range(2, n + 1):
            pre, cur = (
                cur,
                min(
                    pre + cost[i - 2],
                    cur + cost[i - 1],
                ),
            )
        return cur

    def minCostClimbingStairs_dp(self, cost: List[int]) -> int:
        n = len(cost)
        dp = [0] * (n + 1)
        dp[0] = dp[1] = 0
        for i in range(2, n + 1):
            dp[i] = min(
                dp[i - 2] + cost[i - 2],
                dp[i - 1] + cost[i - 1],
            )
        return dp[n]


# @lc code=end

if __name__ == "__main__":
    cost: List[int] = deserialize("List[int]", read_line())
    ans = Solution().minCostClimbingStairs(cost)
    print("\noutput:", serialize(ans, "integer"))
