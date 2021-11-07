package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// returns true if play again
func play(user string) bool {
	rand.Seed(time.Now().UnixNano())
	// setting up the board
	board := initialState()

	fmt.Println("Let's start the game!")
	printBoard(board)

	// if the user chose to be O, then, let AI make the first move randomly
	if user == O {
		fmt.Println("\nAI Thinking...")
		move := Action{
			i: rand.Intn(3),
			j: rand.Intn(3),
		}
		board = result(board, move)
		printBoard(board)
	}

	for {
		action := promptAction(board)
		board = result(board, action)
		printBoard(board)

		if gameEnds(board, user) {
			break
		}

		fmt.Println("\nAI Thinking...")

		move := minimax(board)
		board = result(board, move)
		printBoard(board)
		if gameEnds(board, user) {
			break
		}
	}

	// asking whether if the user wants to play again.
	userResponse := ""

	for userResponse != "y" && userResponse != "n" && userResponse != "yes" && userResponse != "no" {
		userResponse = strings.ToLower(input("Play again? "))
	}

	return userResponse == "y" || userResponse == "yes"
}

func gameEnds(board Board, player string) bool {
	if terminal(board) {
		_winner := winner(board)
		fmt.Printf("Game ends: ")
		if _winner != EMPTY {
			if _winner == player {
				fmt.Println("\033[32mYou won!\033[0m")
			} else {
				fmt.Println("\033[31mYou lost!\033[0m")
			}
		} else {
			fmt.Println("Tie.")
		}
		return true
	}
	return false
}

func clearScreen() {
	clear := make(map[string]func())
	unixClear := func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = unixClear
	clear["darwin"] = unixClear
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearFunction, ok := clear[runtime.GOOS]
	if ok {
		clearFunction()
	}
}

func promptAction(board Board) Action {
	var response string
	var i, j int
	fmt.Println("\nYour turn now!")
	response = input("Row(0-2): ")
	ok := false
	for !ok {
		val, err := strconv.ParseUint(response, 10, 64)
		if err == nil {
			if val > 2 {
				fmt.Println("Input out of range, try again.")
				continue
			}
			i = int(val)
			ok = true
		} else {
			fmt.Println("Invalid input, try again.")
		}
	}
	ok = false
	response = input("Column(0-2): ")
	for !ok {
		val, err := strconv.ParseUint(response, 10, 64)
		if err == nil {
			if val > 2 {
				fmt.Println("Input out of range, try again.")
				continue
			}
			j = int(val)
			ok = true
		} else {
			fmt.Println("Invalid input, try again.")
		}
	}

	action := Action{
		i: i,
		j: j,
	}

	// if the prompted action tries points to a non-empty block
	if board[i][j] != EMPTY {
		fmt.Println("Cell unavailable, try somewhere else.")
		action = promptAction(board)
	}

	return action
}

func main() {
	var player string
	for {
		clearScreen()
		var userResponse string
		for userResponse != "O" && userResponse != "X" {
			userResponse = strings.ToUpper(input("Play as[X/O]: "))
		}
		player = userResponse

		fmt.Printf("You are now playing as %s\n", player)

		// plays the game and at the end, get a response that tells the program
		// whether or not to continue playing a new round
		again := play(player)

		// if response is no, quit
		if !again {
			fmt.Println("\nHave a nice day!")
			break
		}
		fmt.Println()
	}
}
