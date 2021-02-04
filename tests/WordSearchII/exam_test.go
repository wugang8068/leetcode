package WordSearchII

import (
	"fmt"
	"testing"
)

func Test_findWords(t *testing.T) {
	b := [][]byte{[]byte("aaaaaaaaaaaa"), []byte("aaaaaaaaaaaa"), []byte("aaaaaaaaaaaa"), []byte("aaaaaaaaaaaa")}
	words := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	finds := findWords(b, words)
	fmt.Println(finds)
}
