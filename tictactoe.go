package main

import (
	"fmt"
	"math"
)

const X string = "X"
const O string = "O"
const EMPTY string = " "
const BOARD_SIZE int = 3

type Board [][]string
type Action struct {
	i, j int
}
type Packet struct {
	action Action
	value  int
}

func minimax(board Board) Action {
	currPlayer := player(board)
	var buffer Packet
	if currPlayer == X {
		buffer = maxValue(board)
	} else {
		buffer = minValue(board)
	}
	return buffer.action
}

// returns all of the available actions
func actions(board Board) (actions []Action) {
	for i, row := range board {
		for j, cell := range row {
			if cell == EMPTY {
				actions = append(actions, Action{
					i: i,
					j: j,
				})
			}
		}
	}
	return
}

func result(board Board, action Action) Board {
	duplicate := make([][]string, BOARD_SIZE)
	for i := range board {
		duplicate[i] = make([]string, BOARD_SIZE)
		copy(duplicate[i], board[i])
	}
	duplicate[action.i][action.j] = player(board)
	return duplicate
}

func player(board Board) string {
	// set up counters for x and o
	var xCounter, oCounter uint8 = 0, 0

	for _, row := range board {
		for _, cell := range row {
			if cell == X {
				xCounter++
			} else if cell == O {
				oCounter++
			}
		}
	}

	if xCounter > oCounter {
		return O
	} else {
		return X
	}
}

func initialState() (board Board) {
	for i := 0; i < BOARD_SIZE; i++ {
		row := []string{}
		for j := 0; j < BOARD_SIZE; j++ {
			row = append(row, EMPTY)
		}
		board = append(board, row)
	}
	return
}

func printBoard(board Board) {
	const cyan, reset = "\033[36m", "\033[0m"
	output := cyan

	borderTop := "+"
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			borderTop += "-"
		}
		borderTop += "+"
	}
	output += borderTop
	output += "\n"

	for _, row := range board {
		var rowStr string = "| "
		for _, cell := range row {
			if cell == X {
				rowStr += "\033[92m"
			} else if cell == O {
				rowStr += "\033[93m"
			}
			rowStr += cell + reset + cyan + " | "
		}
		output += rowStr + "\n"
		borderBottom := "+"
		for i := 0; i < BOARD_SIZE; i++ {
			for j := 0; j < BOARD_SIZE; j++ {
				borderBottom += "-"
			}
			borderBottom += "+"
		}

		output += borderBottom + "\n"
	}
	fmt.Println(output + reset)
}

func checkHorizontally(board Board) string {
	for i := 0; i < BOARD_SIZE; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != EMPTY {
			return board[i][0]
		}
	}
	return EMPTY
}
func checkVertically(board Board) string {
	for i := 0; i < BOARD_SIZE; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != EMPTY {
			return board[0][i]
		}
	}
	return EMPTY
}

func checkDiagonally(board Board) string {
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != EMPTY {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != EMPTY {
		return board[0][2]
	}
	return EMPTY
}

// Returns the winner of the game, if there is one, and empty otherwise.
func winner(board Board) string {
	h := checkHorizontally(board)
	v := checkVertically(board)
	d := checkDiagonally(board)
	if v != EMPTY || d != EMPTY || h != EMPTY {
		if h == X || v == X || d == X {
			return X
		} else {
			return O
		}
	} else {
		return EMPTY
	}
}

func full(board Board) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == EMPTY {
				return false
			}
		}
	}
	return true
}

func terminal(board Board) bool {
	if winner(board) != EMPTY || full(board) {
		return true
	}
	return false
}

// Returns 1 if X has won the game, -1 if 0 has won the game, 0 otherwise
func utility(board Board) int {
	if terminal(board) {
		if winner(board) == X {
			return 1
		} else if winner(board) == O {
			return -1
		} else {
			return 0
		}
	}
	return 0
}

func maxValue(board Board) (res Packet) {
	if terminal(board) {
		res.value = utility(board)
		return
	}
	allActions := actions(board)

	res.value = math.MinInt
	res.action = allActions[0]

	for _, action := range allActions {
		opponentBest := minValue(result(board, action))
		v := opponentBest.value
		if v > res.value {
			res.value = v
			res.action = action
			if v == 1 {
				return
			}
		}
	}

	return
}

func minValue(board Board) (res Packet) {
	if terminal(board) {
		res.value = utility(board)
		return
	}
	allActions := actions(board)

	res.value = math.MaxInt
	res.action = allActions[0]

	for _, action := range allActions {
		opponentBest := maxValue(result(board, action))
		v := opponentBest.value
		if v < res.value {
			res.value = v
			res.action = action
			if v == -1 {
				return
			}
		}
	}

	return
}
