/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}
	if n == 2 {
		if s[0] == s[1] {
			return s
		}
		return string(s[0])
	}
	a, b, start, end, max := 0, 0, 0, 0, 0
	for i := 1; i < n; i++ {
		if i < n-1 && s[i-1] == s[i+1] {
			a, b = i-1, i+1
			for a > 0 && b < n-1 && s[a-1] == s[b+1] {
				a--
				b++
			}
			if max < b-a+1 {
				max = b - a + 1
				start, end = a, b
			}
		}
		if s[i-1] == s[i] {
			a, b = i-1, i
			for b < n-1 && s[b+1] == s[i] {
				b++
			}
			if max < b-a+1 {
				max = b - a + 1
				start, end = a, b
			}
		}
		if s[i-1] == s[i] {
			a, b = i-1, i
			for a > 0 && b < n-1 && s[a-1] == s[b+1] {
				a--
				b++
			}
			if max < b-a+1 {
				max = b - a + 1
				start, end = a, b
			}
		}
	}
	return s[start : end+1]
}
