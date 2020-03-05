/*
152. 乘积最大子序列
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	start, end, cnt, max, a, b, c, d := 0, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			if cnt == 0 {
				a = i + 1
			}
			b = i
			cnt++
		}
		if i == n-1 {
			end = n
		}
		if nums[i] == 0 {
			end = i
		}
		if nums[i] == 0 || i == n-1 {
			c = nums[start]
			switch cnt % 2 {
			case 0:
				for j := start + 1; j < end; j++ {
					c *= nums[j]
				}
			case 1:
				for j := start + 1; j < b; j++ {
					c *= nums[j]
				}
				if a == n {
					break
				}
				d = nums[a]
				for j := a + 1; j < end; j++ {
					d *= nums[j]
				}
				if c < d {
					c = d
				}
			}
			if max < c {
				max = c
			}
			start, cnt, c, d = end+1, 0, 0, 0
		}
	}
	return max
}
