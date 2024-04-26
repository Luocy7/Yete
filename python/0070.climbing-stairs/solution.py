# Created by luocy at 2024/04/26 20:57
# leetgo: 1.4.5
# https://leetcode.com/problems/climbing-stairs/

from leetgo_py import *

# @lc code=begin

# dp[i] = dp[i - 1] + dp[i - 2]


class Solution:
    def climbStairs(self, n: int) -> int:
        pre, cur = 0, 1
        for _ in range(n):
            pre, cur = cur, pre + cur
        return cur


# @lc code=end

if __name__ == "__main__":
    n: int = deserialize("int", read_line())
    ans = Solution().climbStairs(n)
    print("\noutput:", serialize(ans, "integer"))
