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

type position [2]int

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if board[i][j] == word[0] {
				var tmp byte
				tmp, board[i][j] = board[i][j], 0
				if existChar(board, word, 1, position{i, j}) {
					return true
				}
				board[i][j] = tmp
			}

		}
	}

	return false
}

func existChar(board [][]byte, word string, index int, lastPos position) bool {

	if len(word) <= index {
		return true
	}

	// left
	if lastPos[1] > 0 && checkPosition(board, word, index, position{lastPos[0],lastPos[1]-1}) {
		return true
	}

	// top
	if lastPos[0] > 0 && checkPosition(board, word, index, position{lastPos[0]-1,lastPos[1]}) {
		return true
	}

	// right
	if lastPos[1] < len(board[lastPos[0]])-1 && checkPosition(board, word, index, position{lastPos[0],lastPos[1]+1}) {
		return true
	}

	// bottom
	if lastPos[0] < len(board)-1 && checkPosition(board, word, index, position{lastPos[0]+1,lastPos[1]}) {
		return true
	}

	return false
}

func checkPosition(board [][]byte, word string, index int, pos position) bool {

	if board[pos[0]][pos[1]] != word[index] {
		return false
	}

	var tmp byte
	tmp, board[pos[0]][pos[1]] = board[pos[0]][pos[1]], 0
	if existChar(board, word, index+1, position{pos[0], pos[1]}) {
		return true
	}
	board[pos[0]][pos[1]] = tmp

	return false
}