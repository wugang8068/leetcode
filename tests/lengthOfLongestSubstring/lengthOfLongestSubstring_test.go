package lengthOfLongestSubstring

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	lastMaxLength := 0
	startIndex := 0
	chars := []byte(s)
	lastMaxLeft := 0
	lastMaxRight := 0
	for startIndex < len(chars){
		leftIndex  := startIndex
		o := leftIndex
		rightIndex := startIndex + 1
		for {
			if leftIndex < 0 || rightIndex > len(chars){
				if o == startIndex && startIndex > 0 {
					leftIndex = startIndex - 1
					o = leftIndex
					rightIndex = startIndex + 1
					if leftIndex <0 || rightIndex > len(chars) {
						startIndex++
						break
					}
				}else{
					startIndex++
					break
				}
			}
			if isAvailableChars(chars[leftIndex:rightIndex]){
				if lastMaxLength < rightIndex - leftIndex {
					lastMaxLength = rightIndex - leftIndex
					lastMaxLeft = leftIndex
					lastMaxRight = rightIndex
				}
			}
			leftIndex--
			rightIndex++
		}
	}
	return string(chars[lastMaxLeft:lastMaxRight])
}

func isAvailableChars(chars []byte) bool{
	if len(chars) == 0 {
		return false
	}
	l := 0
	r := len(chars) - 1
	for {
		if len(chars) % 2 == 0 {
			if r < l {
				break
			}
		}else{
			if r == l {
				break
			}
		}
		if chars[r] == chars[l]{
			r--
			l++
			continue
		}else{
			return false
		}
	}
	return true
}

func TestString(t *testing.T)  {
	fmt.Println(longestPalindrome("abcbbcba"))
}