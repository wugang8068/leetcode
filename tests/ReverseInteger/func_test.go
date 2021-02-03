package ReverseInteger

import (
	"fmt"
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println(reverse(100))
	fmt.Println(reverse(2147483648), math.MaxInt32)
	fmt.Println(reverse(12345))
}