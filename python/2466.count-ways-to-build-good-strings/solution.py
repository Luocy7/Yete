# Created by luocy at 2024/04/27 14:31
# leetgo: 1.4.5
# https://leetcode.com/problems/count-ways-to-build-good-strings/

from leetgo_py import *

# @lc code=begin

# construct a i lenght good string equals to

# dp[i] = (dp[i - zero] if i > zero else 0) + (dp[i - one] if i > one else 0)
# dp[0] = 1


class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        pass


# @lc code=end

if __name__ == "__main__":
    low: int = deserialize("int", read_line())
    high: int = deserialize("int", read_line())
    zero: int = deserialize("int", read_line())
    one: int = deserialize("int", read_line())
    ans = Solution().countGoodStrings(low, high, zero, one)
    print("\noutput:", serialize(ans, "integer"))
