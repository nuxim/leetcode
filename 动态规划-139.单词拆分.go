/*
139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package leetcode

func wordBreak(s string, wordDict []string) bool {
	m := map[string]int{}
	for i, v := range wordDict {
		m[v] = i
	}
	//return doWordBreak(s, m)
	n := len(s)
	k := make([]int, n+1)
	ok := false
	k[0] = 1
	for i:=0;i<n;i++ {
		for j := 0; j <= i; j++ {
			if k[j] == 0 {
				continue
			}
			_, ok = m[s[j:i+1]]
			if ok {
				k[i+1] = 1
				if i == n - 1 {
					return true
				}
			}
		}
	}
	return false
}

func doWordBreak(s string, m map[string]int) bool {
	_, ok := m[s]
	if ok {
		return true
	}
	n := len(s)
	for i:=n-1;i>0;i--{
		t := s[i:n]
		_, ok = m[t]
		if ok && doWordBreak(s[:i], m) {
			return true
		}
	}
	return false
}