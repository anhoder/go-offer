package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"

	//arr := [][]byte{{'a', 'b'}, {'c', 'd'}}
	//word := "abcd"

	arr := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	word := "AAB"


	fmt.Println(exist(arr, word))
}

type position [2]int

var boardMap map[byte][]position

func exist(board [][]byte, word string) bool {

	boardMap = map[byte][]position{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			boardMap[board[i][j]] = append(boardMap[board[i][j]], position{i, j})
		}
	}


	return existChar(board, word, 0, position{-1, -1})
}

func existChar(board [][]byte, word string, index int, lastPos position) bool {

	if len(word) <= index {
		return true
	}

	if _, exist := boardMap[word[index]]; !exist {
		return false
	}

	for _, p := range boardMap[word[index]] {

		if board[p[0]][p[1]] == 0 {
			continue
		}

		if lastPos[0] >= 0 &&
			lastPos[1] >= 0 &&
			math.Abs(float64(p[0] - lastPos[0])) + math.Abs(float64(p[1] - lastPos[1])) != 1 {
			continue
		}

		tmp := board[p[0]][p[1]]
		board[p[0]][p[1]] = 0

		if existChar(board, word, index + 1, p) {
			board[p[0]][p[1]] = tmp
			return true
		}

		board[p[0]][p[1]] = tmp

	}


	return false
}