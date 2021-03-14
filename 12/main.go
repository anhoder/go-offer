package main

import (
	"fmt"
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

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if existChar(board, word, 0, i, j) {
				return true
			}

		}
	}

	return false
}

func existChar(board [][]byte, word string, index, i, j int) bool {

	if len(word) <= index {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}


	tmp := board[i][j]
	board[i][j] = 0

	res := existChar(board, word, index+1, i-1, j) ||
		existChar(board, word, index+1, i, j-1) ||
		existChar(board, word, index+1, i+1, j) ||
		existChar(board, word, index+1, i, j+1)

	board[i][j] = tmp

	return res
}