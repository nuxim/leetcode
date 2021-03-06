/*
221. 最大正方形
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

测试用例：
[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
[["1","1","1","1"],["1","1","1","1"],["1","1","1","1"]]
[["1","0","1","0"],["1","0","1","1"],["1","0","1","1"],["1","1","1","1"]]
[["1","1","0","1"],["1","1","0","1"],["1","1","1","1"]]
[["0","0","0","1"],["1","1","0","1"],["1","1","1","1"],["0","1","1","1"],["0","1","1","1"]]
[["0","0","1","0"],["1","1","1","1"],["1","1","1","1"],["1","1","1","0"],["1","1","0","0"],["1","1","1","1"],["1","1","1","0"]]
[["1","0","1","0","0","1","1","1","0"],["1","1","1","0","0","0","0","0","1"],["0","0","1","1","0","0","0","1","1"],["0","1","1","0","0","1","0","0","1"],["1","1","0","1","1","0","0","1","0"],["0","1","1","1","1","1","1","0","1"],["1","0","1","1","1","0","0","1","0"],["1","1","1","0","1","0","0","0","1"],["0","1","1","1","1","0","0","1","0"],["1","0","0","1","1","1","0","0","0"]]
[["0","1","0","1","0","1","0","1","1","0","1","1","0","1","0"],["0","1","1","1","1","1","0","1","1","0","1","1","1","1","0"],["0","1","1","1","1","1","1","1","0","1","1","1","0","1","1"],["1","0","1","1","0","1","0","0","1","1","1","0","0","0","1"],["0","1","0","1","1","1","0","1","0","1","1","0","0","1","1"],["1","1","1","1","0","1","1","0","1","1","1","0","1","1","1"],["1","1","1","1","0","1","0","1","0","1","1","1","1","1","0"],["0","1","1","1","1","1","1","1","0","1","1","1","1","1","0"],["0","1","1","1","0","1","0","0","0","0","1","1","1","1","1"],["1","1","1","1","1","0","1","0","1","1","1","0","0","1","1"],["0","1","1","1","0","0","1","1","1","1","1","1","1","1","1"],["1","1","1","0","0","0","1","1","1","0","1","0","0","1","1"],["1","1","1","1","1","1","1","1","1","0","0","1","1","1","1"],["1","0","1","1","1","1","1","1","1","1","1","1","0","1","1"],["1","1","1","0","1","0","1","0","1","0","1","1","1","1","1"]]
*/
package leetcode

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	a := make([]int, m)
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == byte('0') {
				a[j] = 0
			} else {
				a[j] += 1
			}
		}
		for j := 0; j < m; j++ {
			if a[j] > 0 {
				if max == 0 {
					max = 1
				}
				t := a[j]
				for k := 1; k+j <= m; k++ {
					if t > a[j+k-1] {
						t = a[j+k-1]
					}
					if a[j+k-1] >= k && k > 1 && k <= t && max < k*k {
						max = k * k
						break
					}
					if a[j+k-1] <= k {
						break
					}
				}
			}
		}
	}
	return max
}
