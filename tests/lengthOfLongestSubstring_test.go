package tests

import (
	"fmt"
	"testing"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func lengthOfLongestSubstring(s string) int {
	chars := []byte(s)
	var l = 0
	var savedChars []byte
	for _, c := range chars {
		var f *int
		for index, i := range savedChars {
			if i == c {
				f = &index
				break
			}
		}
		savedChars = append(savedChars, c)
		if f != nil {
			savedChars = savedChars[*f + 1:]
			continue
		}
		_l := len(savedChars)
		if _l > l {
			l = _l
		}
	}
	return l
}

func TestLengthOfLongestSubstring(t *testing.T)  {
	fmt.Println(lengthOfLongestSubstring("aab") == 2)
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("dvdf") == 3)
	fmt.Println(lengthOfLongestSubstring("anviaj") == 5)

}
