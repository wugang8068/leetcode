package zigZagConversion

import (
	"math"
)

func convert(s string, numRows int) string {
	rows := make([][]int32, int(math.Min(float64(numRows), float64(len(s)))))
	if numRows == 1 {
		return s
	}
	goingDown := false
	curRow := 0
	for _, bt := range s {
		rows[curRow] = append(rows[curRow], bt)
		if curRow == 0 || curRow == numRows - 1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		}else{
			curRow -= 1
		}
	}
	out := ""
	for _, row := range rows {
		out += string(row)
	}
	return out
}
