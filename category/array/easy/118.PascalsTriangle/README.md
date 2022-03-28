118. Pascal's Triangle
     题目
     Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.



Note: In Pascal's triangle, each number is the sum of the two numbers directly above it.

![img.png](img.png)

Example:

```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```
题目大意
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。在杨辉三角中，每个数是它左上方和右上方的数的和。

解题思路
给定一个 n，要求打印杨辉三角的前 n 行。
简单题。按照杨辉三角的生成规则循环打印即可。