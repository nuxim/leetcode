/*
91. 解码方法
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if s[0] == byte('0') {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 1; i < n; i++ {
		if s[i-1] != byte('2') && s[i-1] != byte('1') && s[i] == byte('0') {
			return 0
		}
		if s[i-1] > byte('2') || s[i-1] == byte('0') || (s[i-1] == byte('2') && s[i] > byte('6')) {
			a, b = b, b
		} else if s[i] == byte('0') {
			a, b = b, a
		} else {
			a, b = b, a+b
		}
	}
	return b
}

func numDecodingsRecrusive(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		if s[0] > byte('2') || (s[0] == byte('2') && s[1] > byte('6')) {
			return 1
		}
		return 2
	}
	if s[n-2] > byte('2') || (s[n-2] == byte('2') && s[n-1] > byte('6')) {
		return numDecodings(s[:n-1])
	}
	return numDecodings(s[:n-2]) + numDecodings(s[:n-1])
}
