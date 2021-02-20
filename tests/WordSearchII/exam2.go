package WordSearchII

import (
	"fmt"
	"time"
)

func findWords2(board [][]byte, words []string) []string {
	t := time.Now().Nanosecond()
	output := []string{}
	I := len(board)
	if I == 0{
		return output
	}
	J := len(board[0])
	charIndx := map[byte][]int{}

	for i:=0; i<I; i++{
		for j:=0; j<J; j++{
			index := i*J + j
			if v, ok := charIndx[board[i][j]]; ok{
				charIndx[board[i][j]] = append(v, index)
			}else{
				charIndx[board[i][j]] = []int{index}
			}
		}
	}

	for _, word := range words{
		if chIndex, ok := charIndx[word[0]]; ok{
			for _, idx := range chIndex{
				visited := make([]bool, I*J)
				visited[idx] = true
				if doDFS(word, charIndx, 1, idx, visited, I, J){
					output = append(output, word)
					break
				}
			}
		}
	}
	t2 := time.Now().Nanosecond()
	fmt.Printf("costTime %d:ns \n", t2 - t)
	return output
}

func doDFS(word string, charIndx map[byte][]int, wordIndex int, sqIndex int, visited []bool, I int, J int) bool{
	if wordIndex == len(word){
		return true
	}

	if chIndex, ok := charIndx[word[wordIndex]]; ok{
		for _, idx := range chIndex{
			//check if it is vertical or horizontal adjacent
			if !visited[idx] && isAdjacent(sqIndex, idx, I, J){
				visited[idx] = true
				success := doDFS(word, charIndx, wordIndex+1, idx, visited, I, J)
				if success{
					return true
				}else{
					visited[idx] = false
				}
			}
		}
	}
	return false
}

func isAdjacent(index int, target int, I int, J int)bool{
	i := index/J
	j := index%J

	iT := target /J
	jT := target % J

	if iT == i && (jT == j+1 || jT == j-1){
		return true
	}

	if jT == j && (iT == i-1 || iT == i+1){
		return true
	}
	return false
}
