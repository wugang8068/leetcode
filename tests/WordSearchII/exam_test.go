package WordSearchII

import (
	"fmt"
	"testing"
)

func Test_findWords(t *testing.T) {
	b := [][]byte{[]byte("oaan"), []byte("etae"), []byte("ihkr"), []byte("iflv")}
	words := []string{"oath","pea","eat","rain","hklf","hf"}
	finds := findWords3(b, words)
	fmt.Println(finds)
}
