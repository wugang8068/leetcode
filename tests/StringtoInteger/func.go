package StringtoInteger

import (
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	fmt.Println(s)
	i, _ := strconv.Atoi(s)
	return i
}