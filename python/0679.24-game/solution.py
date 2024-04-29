# Created by luocy at 2024/04/29 10:16
# leetgo: 1.4.5
# https://leetcode.com/problems/24-game/
import math
from typing import *

from leetgo_py import *

# @lc code=begin


class Solution:
    def judgePoint24(self, cards: List[int | float]) -> bool:
        # if len(cards) == 1:
        #     return math.isclose(cards[0], 24)
        #
        # for _ in range(len(cards)):
        #     a = cards.pop(0)  # 摸一张 (queue 操作)
        #     for _ in range(len(cards)):
        #         b = cards.pop(0)  # 再摸一张 (queue 操作)
        #         for value in [a + b, a - b, a * b, b and a / b]:  # 算一下
        #             cards.append(value)  # 记下来 (stack 操作)
        #             if self.judgePoint24(cards):
        #                 return True
        #             cards.pop()  # (stack 操作)
        #         cards.append(b)  # (queue 操作)
        #     cards.append(a)  # (queue 操作)
        # return False
        if len(cards) == 1:
            return math.isclose(cards[0], 24)
            # 每次把计算结果 和之前cards数组中其他元素组成的列表concate起来 作为新的cards 向下递归
        for i in range(len(cards) - 1):
            for j in range(i + 1, len(cards)):
                a, b = cards[i], cards[j]
                ext = cards[0:i] + cards[i + 1: j] + cards[j + 1:]
                if self.judgePoint24([a + b] + ext):
                    return True
                if self.judgePoint24([a * b] + ext):
                    return True
                if self.judgePoint24([a - b] + ext):
                    return True
                if a != b and self.judgePoint24([b - a] + ext):
                    return True
                if b != 0 and self.judgePoint24([a / b] + ext):
                    return True
                if a != 0 and a != b and self.judgePoint24([b / a] + ext):
                    return True
            # 走到这一步 说明之前都不行
        return False


# @lc code=end

if __name__ == "__main__":
    cards: List[int] = deserialize("List[int]", read_line())
    ans = Solution().judgePoint24(cards)
    print("\noutput:", serialize(ans, "boolean"))
