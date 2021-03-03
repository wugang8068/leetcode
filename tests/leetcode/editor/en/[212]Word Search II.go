package test

import (
	"encoding/json"
	"fmt"
)

//Given an m x n board of characters and a list of strings words, return all wor
//ds on the board.
//
// Each word must be constructed from letters of sequentially adjacent cells, wh
//ere adjacent cells are horizontally or vertically neighboring. The same letter c
//ell may not be used more than once in a word.
//
//
// Example 1:
//
//
//Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f"
//,"l","v"]], words = ["oath","pea","eat","rain"]
//Output: ["eat","oath"]
//
//
// Example 2:
//
//
//Input: board = [["a","b"],["c","d"]], words = ["abcb"]
//Output: []
//
//
//
// Constraints:
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] is a lowercase English letter.
// 1 <= words.length <= 3 * 104
// 1 <= words[i].length <= 10
// words[i] consists of lowercase English letters.
// All the strings of words are unique.
//
// Related Topics Backtracking Trie
// 👍 3473 👎 143

//leetcode submit region begin(Prohibit modification and deletion)

type step struct {
	PreStep       *step
	Direction     string
	RotationCount int
	PositionX     int
	PositionY     int
	Byte          byte
	Bytes         []byte
	Board         [][]byte
}

func (p *step) rotationDirection() {
	p.Direction = moveDirections[p.Direction]
	p.RotationCount++
}

func (p step) m() map[string]interface{} {
	return map[string]interface{}{
		"pre":           nil,
		"direction":     p.Direction,
		"rotationCount": p.RotationCount,
		"px":            p.PositionX,
		"py":            p.PositionY,
		"b":             string(p.Byte),
		"bs":            string(p.Bytes),
	}
}

func (p step) json() string {
	m := p.m()
	if p.PreStep != nil {
		m["pre"] = p.PreStep.m()
	}
	b, _ := json.Marshal(m)
	return string(b)
}

func (p *step) moveUp() *step {
	cp := *p
	if p.PositionX == 0 || p.Board[p.PositionX-1][p.PositionY] == 0 {
		cp.rotationDirection()
		return moveToPosition(&cp)
	}
	cp.Byte = p.Board[p.PositionX-1][p.PositionY]
	cp.PositionX = p.PositionX - 1
	cp.Board = copyBoard(p.Board)
	cp.Board[p.PositionX-1][p.PositionY] = byte(0)
	cp.PreStep = p
	cp.RotationCount = 0
	return &cp
}

func (p *step) moveRight() *step {
	cp := *p
	if p.PositionY == len(p.Board[p.PositionX])-1 || p.Board[p.PositionX][p.PositionY+1] == 0 {
		cp.rotationDirection()
		return moveToPosition(&cp)
	}
	cp.Byte = p.Board[p.PositionX][p.PositionY+1]
	cp.PositionY = p.PositionY + 1
	cp.Board = copyBoard(p.Board)
	cp.Board[p.PositionX][p.PositionY+1] = byte(0)
	cp.PreStep = p
	cp.RotationCount = 0
	return &cp
}

func (p *step) moveDown() *step {
	cp := *p
	if p.PositionX == len(p.Board)-1 || p.Board[p.PositionX+1][p.PositionY] == 0 {
		cp.rotationDirection()
		return moveToPosition(&cp)
	}
	cp.Byte = p.Board[p.PositionX+1][p.PositionY]
	cp.PositionX = p.PositionX + 1
	cp.Board = copyBoard(p.Board)
	cp.Board[p.PositionX+1][p.PositionY] = byte(0)
	cp.PreStep = p
	cp.RotationCount = 0
	return &cp
}

func (p *step) moveLeft() *step {
	cp := *p
	if p.PositionY == 0 || p.Board[p.PositionX][p.PositionY-1] == 0 {
		cp.rotationDirection()
		return moveToPosition(&cp)
	}
	cp.Byte = p.Board[p.PositionX][p.PositionY-1]
	cp.PositionY = p.PositionY - 1
	cp.Board = copyBoard(p.Board)
	cp.Board[p.PositionX][p.PositionY-1] = byte(0)
	cp.PreStep = p
	cp.RotationCount = 0
	return &cp
}

var moveDirections = map[string]string{
	"up":    "right",
	"right": "down",
	"down":  "left",
	"left":  "up",
}

func copyBoard(board [][]byte) [][]byte {
	var cp [][]byte
	for _, r := range board {
		cr := make([]byte, len(r))
		copy(cr, r)
		cp = append(cp, cr)
	}
	return cp
}

func findWords(board [][]byte, words []string) []string {
	var finds []string
	loop := 0
	defer func() {
		fmt.Printf("loop count %d\n", loop)
	}()
	for positionX, rows := range board {
		for positionY, b := range rows {
			if len(words) == 0 {
				return finds
			}
			cp := copyBoard(board)
			cp[positionX][positionY] = byte(0)
			lastStep := &step{
				PreStep:   nil,
				Direction: "right",
				PositionX: positionX,
				PositionY: positionY,
				Byte:      b,
				Bytes:     []byte{b},
				Board:     cp,
			}
			_, find := findBytesIn(lastStep.Bytes, words)
			if len(find) > 0 {
				words = diff(words, find)
				finds = append(finds, find...)
			}
			for {
				loop++
				nextStep := moveToPosition(lastStep)
				//如果回到原点, 并且重复步骤, 说明当前点已经循环结束
				if nextStep == nil {
					break
				}
				//标志已经结束
				nextStep.Bytes = append(nextStep.Bytes, nextStep.Byte)
				possible, find := findBytesIn(nextStep.Bytes, words)
				//如果不可能找到或者
				if !possible {
					//回溯到上一步
					lastStep = nextStep.PreStep
					lastStep.rotationDirection()
					//fmt.Printf("possible:%v, cusBytes:%s, %v \n", possible, string(nextStep.Bytes), lastStep.json())
					continue
				} else {
					lastStep = nextStep
					if len(find) > 0 {
						words = diff(words, find)
						finds = append(finds, find...)
					}
					//fmt.Printf("possible:%v, cusBytes:%s, %v \n", possible, string(nextStep.Bytes), nextStep.json())
					continue
				}
			}
		}
	}
	return finds
}

func diff(word []string, inDiffWord []string) []string {
	f := make(map[string]bool)
	for _, w := range word {
		f[w] = true
	}
	for _, inf := range inDiffWord {
		if f[inf] {
			delete(f, inf)
		}
	}
	var df []string
	for kf, _ := range f {
		df = append(df, kf)
	}
	return df
}

func findBytesIn(b []byte, words []string) (possible bool, find []string) {
	possible = false
	for _, word := range words {
		bs := string(b)
		if bs == word {
			find = append(find, word)
		}
		wordsL := len(word)
		if wordsL >= len(b) {
			pb := true
			for i, _b := range b {
				if word[i] != _b {
					pb = false
				}
			}
			if pb == true {
				possible = true
			}
		}
	}
	return
}

func moveToPosition(p *step) *step {
	if p.RotationCount >= 4 {
		if p.PreStep != nil {
			p.PreStep.rotationDirection()
			return moveToPosition(p.PreStep)
		}
		return nil
	}
	switch p.Direction {
	case "up":
		return p.moveUp()
	case "right":
		return p.moveRight()
	case "down":
		return p.moveDown()
	case "left":
		return p.moveLeft()
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
