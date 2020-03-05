/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	a := make([]int, m)
	a[0] = grid[0][0]
	for i := 0; i < m-1; i++ {
		a[i+1] = a[i] + grid[0][i+1]
	}
	if n == 1 {
		return a[m-1]
	}
	for k := 1; k < n; k++ {
		a[0] += grid[k][0]
		for i := 0; i < m-1; i++ {
			if a[i+1] < a[i] {
				a[i+1] = a[i+1] + grid[k][i+1]
			} else {
				a[i+1] = a[i] + grid[k][i+1]
			}
		}
	}
	return a[m-1]
}

func minSum(grid [][]int, m, n int) int {
	sum := 0
	if n == 1 {
		for i := 0; i < m; i++ {
			sum += grid[0][i]
		}
		return sum
	}
	if m == 1 {
		for i := 0; i < n; i++ {
			sum += grid[i][0]
		}
		return sum
	}
	a := minSum(grid, m-1, n)
	b := minSum(grid, m, n-1)
	if a > b {
		return grid[n-1][m-1] + b
	}
	return grid[n-1][m-1] + a
}
