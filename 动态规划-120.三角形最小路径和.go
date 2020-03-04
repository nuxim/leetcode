/*
120. 三角形最小路径和
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package leetcode

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}
	a := make([]int, n)
	a[0] = triangle[0][0]
	for i:= 1; i<n;i++ {
		a[i] = a[i-1] + triangle[i][i]
		for j:=i-1;j>0;j-- {
			if a[j] < a[j-1] {
				a[j] += triangle[i][j]
			} else {
				a[j] = a[j-1] + triangle[i][j]
			}
		}
		a[0] += triangle[i][0]
	}
	for i:=0;i<n;i++ {
		if a[0] > a[i] {
			a[0] = a[i]
		}
	}
	return a[0]
}