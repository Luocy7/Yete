# Two Pointers

## 模板小结

- **首尾双指针**

  ![img](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/raw/master/_assets/TIM%E6%88%AA%E5%9B%BE20180928102534.png)

  - 一般用于寻找数组/双向链表中满足条件的**两个节点**；如果是寻找多个数，则先固定前 n-2 个数；
  - 为了不遗漏所有可能情况，可能要求数组**有序**；
  - **典型问题**：两数之和、三数之和三角形计数

- **同向双指针**

  [![img](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/raw/master/_assets/TIM%E6%88%AA%E5%9B%BE20180928102605.png)](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/blob/master/_assets/TIM截图20180928102605.png)

  - **数组**中，一般用于寻找满足某个条件的**连续区间**；
  - **链表**相关问题中经常会使用**快慢双指针**来寻找某个节点；
  - **典型问题**：最小覆盖子串、数组中的最长山脉（同向双指针）

- **反向双指针**

  [![img](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/raw/master/_assets/TIM%E6%88%AA%E5%9B%BE20180930171522.png)](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/blob/master/_assets/TIM截图20180930171522.png)

  - 先依次遍历都某个节点，然后使用双指针从该节点反向判断是否满足条件。
  - **典型问题**：最长回文子串、数组中的最长山脉（反向双指针）

- **分离双指针**

  [![img](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/raw/master/_assets/TIM%E6%88%AA%E5%9B%BE20180928103003.png)](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/blob/master/_assets/TIM截图20180928103003.png)

  - 输入是两个数组/链表，两个指针分别在两个容器中移动；
  - 根据问题的不同，初始位置可能都在头部，或者都在尾部，或一头一尾。

  [![img](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/raw/master/_assets/TIM%E6%88%AA%E5%9B%BE20180928103312.png)](https://github.com/seagle0128/Algorithm_Interview_Notes-Chinese/blob/master/_assets/TIM截图20180928103312.png)

  - **典型问题**：[两个数组的交集、合并两个有序数组]

### Index

- 首尾双指针
  - 两数之和
  - 三数之和
  - N 数之和
  - 最接近的三数之和
  - 两数之和 - 小于等于目标值的个数
  - 三数之和 - 小于等于目标值的个数
  - 三角形计数（Valid Triangle Number）
  - 接雨水（Trapping Rain Water）（一维）
  - 盛最多水的容器（Container With Most Water）
  - 反转字符串（Reverse String）
  - 颜色分类（Sort Colors）
- 同向双指针
  - 数组中的最长山脉（Longest Mountain in Array）（同向双指针）
  - 最小覆盖子串（Minimum Window Substring）
  - 长度最小的子数组（Minimum Size Subarray Sum）
  - 无重复字符的最长子串（Longest Substring Without Repeating Characters）
  - 水果成篮（Fruit Into Baskets）
- 反向双指针
  - 数组中的最长山脉（Longest Mountain in Array）（反向双指针）
  - 最长回文子串（Longest Palindromic Substring）
- 分离双指针
  - 两个数组的交集（Intersection of Two Arrays）
  - 合并两个有序数组（Merge Sorted Array）
- 链表相关
  - 分隔链表（Partition List）
  - 链表排序（Sort List）
    - 链表快排
    - 链表归并
    - 链表插入排序
    - 链表选择排序
    - 链表冒泡排序
  - 旋转链表（Rotate List）
- 其他
  - 最小区间（Smallest Range）
