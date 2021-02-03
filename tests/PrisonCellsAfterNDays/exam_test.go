package PrisonCellsAfterNDays

import (
	"fmt"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	nR := prisonAfterNDays([]int{1,0,0,1,0,0,0,1}, 826)
	fmt.Println(nR)
}