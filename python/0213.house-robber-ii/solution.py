# Created by luocy at 2024/04/27 15:21
# leetgo: 1.4.5
# https://leetcode.com/problems/house-robber-ii/

from typing import *

from leetgo_py import *

# @lc code=begin

# dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
# like 198 rob-i  transfer to
# - rob the first house which means nums[0] + rob(nums[2:-1])
# - not rob the first house which means rob(nums[1:])
# get the bigger one is the answer


class Solution:
    def rob(self, nums: List[int]) -> int:
        pass


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    ans = Solution().rob(nums)
    print("\noutput:", serialize(ans, "integer"))
