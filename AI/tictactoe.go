package AI

import (
	"fmt"
	"log"
	"math"
)

var X = "X"
var O = "O"

func Player(board map[string]string) string {
	/*
		Returns player who has the next turn on a board
	*/

	var xCount, oCount int
	for _, value := range board {
		if value == X {
			xCount++
		} else if value == O {
			oCount++
		}
	}

	if xCount > oCount {
		return O
	} else {
		return X
	}
}

func actions(board map[string]string) map[string]bool {
	/*
		Returns set of all possible actions available on the board
	*/
	res := make(map[string]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[cell(i, j)] == "" {
				res[cell(i, j)] = true
			}
		}
	}
	return res
}

func Result(board map[string]string, action string) map[string]string {
	/*
		Returns the board that results from making move action on the board.
	*/
	if board[action] != "" {
		log.Fatalln("Invalid Move")
	}

	res := make(map[string]string)
	for k, v := range board {
		res[k] = v
	}

	res[action] = Player(board)

	return res
}

func Winner(board map[string]string) string {
	/*
		Returns the winner of the game, if there is one
	*/
	for i := 0; i < 3; i++ {
		if (board[cell(i, 0)] == board[cell(i, 1)]) && (board[cell(i, 1)] == board[cell(i, 2)]) && (board[cell(i, 0)] != "") {
			return board[cell(i, 0)]
		}
	}

	for i := 0; i < 3; i++ {
		if (board[cell(0, i)] == board[cell(1, i)]) && (board[cell(1, i)] == board[cell(2, i)]) && (board[cell(0, i)] != "") {
			return board[cell(0, i)]
		}
	}

	if (board[cell(0, 0)] == board[cell(1, 1)]) && (board[cell(1, 1)] == board[cell(2, 2)]) && (board[cell(0, 0)] != "") {
		return board[cell(0, 0)]
	}

	if (board[cell(0, 2)] == board[cell(1, 1)]) && (board[cell(1, 1)] == board[cell(2, 0)]) && (board[cell(2, 0)] != "") {
		return board[cell(2, 0)]
	}

	return ""
}

func Terminal(board map[string]string) bool {
	/*
		Returns true if game is over, false otherwise.
	*/
	if win := Winner(board); win != "" {
		return true
	}

	for _, value := range board {
		if value == "" {
			return false
		}
	}

	return true
}

func utility(board map[string]string) int {
	if win := Winner(board); win == X {
		return 1
	} else if win == O {
		return -1
	}
	return 0
}

func Minimax(board map[string]string) string {
	if Terminal(board) {
		return ""
	}

	if Player(board) == X {
		maxActionValue, maxAction := -1, ""
		for action := range actions(board) {
			if v := minValue(Result(board, action)); v > maxActionValue {
				maxAction = action
				maxActionValue = v
			}
		}
		return maxAction
	} else {
		minActionValue, minAction := 1, ""
		for action := range actions(board) {
			if v := maxValue(Result(board, action)); v < minActionValue {
				minAction = action
				minActionValue = v
			}
		}
		return minAction
	}
}

func maxValue(board map[string]string) int {
	if Terminal(board) {
		return utility(board)
	}
	v := -math.MaxInt
	for action := range actions(board) {
		v = max(v, minValue(Result(board, action)))
	}
	return v
}

func minValue(board map[string]string) int {
	if Terminal(board) {
		return utility(board)
	}
	v := math.MaxInt
	for action := range actions(board) {
		v = min(v, maxValue(Result(board, action)))
	}
	return v
}

func cell(i, j int) string {
	return fmt.Sprintf("c%v", 3*i+j)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
