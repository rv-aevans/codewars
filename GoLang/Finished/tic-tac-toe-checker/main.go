package main

import (
	"fmt"
)

var input = [3][3]int{
	{2, 0, 0},
	{2, 1, 0},
	{1, 2, 1},
}

func main() {
	fmt.Println(IsSolved(input))
}

func IsSolved(board [3][3]int) int {
	empty := 0
	for i, row := range board {
		for _, d := range row {
			if d == 0 {
				empty = -1
			}
		}
		if row[0] == row[1] && row[1] == row[2] && row[0] != 0 {
			return row[0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != 0 {
			return board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != 0 {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[0][2] != 0 {
		return board[0][2]
	}
	return empty
}
