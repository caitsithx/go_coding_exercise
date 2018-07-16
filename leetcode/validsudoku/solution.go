package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rowChecked = true
	var columnStates [9][10]int
	var subBoardStates [9][10]int
	for i := range board {
		var rowStates [10]int

		for j := range board[i] {
			if board[i][j] != '.' {
				var ch = board[i][j] - 48
				if rowStates[ch] != 0 {
					rowChecked = false
					break
				}
				rowStates[ch] = 1
				if columnStates[j][ch] != 0 {
					rowChecked = false
					break
				}
				columnStates[j][ch] = 1

				subBoardIndex := i/3*3 + j/3
				if subBoardStates[subBoardIndex][ch] != 0 {
					rowChecked = false
					break
				}
				subBoardStates[subBoardIndex][ch] = 1
			}
		}

		if !rowChecked {
			break
		}
	}

	return rowChecked
}

func main() {
	var chars = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(chars))
}
