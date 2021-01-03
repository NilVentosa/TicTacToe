package main

import (
	"fmt"
)

func main() {
	play()
}

func play() {
	board := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	turn := "x"

	for {
		isGameOver, winner := isGameOver(board)
		if isGameOver {
			printResult(winner, board)
			break
		} else if turn == "x" {
			board = playerMove(board)
			turn = "o"
		} else {
			board = computerMove(board)
			turn = "x"
		}
	}
}

func findBestComputerMove(board [9]string) int {
	possibilities := getEmptySpots(board)
	best := 1
	var bestIndex int
	var currentBoard [9]string

	for _, item := range possibilities {
		currentBoard = board
		currentBoard[item] = "o"
		miniMaxValue := minimax(currentBoard, len(possibilities), true)
		if miniMaxValue < best {
			best = miniMaxValue
			bestIndex = item
		}
	}
	return bestIndex
}

func minimax(board [9]string, depth int, isMaximizingPlayer bool) int {
	isGameOver, _ := isGameOver(board)
	if isGameOver || depth == 0 || len(getEmptySpots(board)) == 0 {
		return evaluate(board)
	}

	var currentBoard [9]string
	if isMaximizingPlayer {
		maxEval := -1
		for _, item := range getEmptySpots(board) {
			currentBoard = board
			currentBoard[item] = "x"
			eval := minimax(currentBoard, depth-1, false)
			maxEval = max([]int{maxEval, eval})
		}
		return maxEval
	}

	minEval := 1
	for _, item := range getEmptySpots(board) {
		currentBoard = board
		currentBoard[item] = "o"
		eval := minimax(currentBoard, depth-1, true)
		minEval = min([]int{minEval, eval})
	}
	return minEval
}

func evaluate(board [9]string) int {
	isGameOver, winner := isGameOver(board)
	if isGameOver {
		if winner == "x" {
			return 1
		} else if winner == "o" {
			return -1
		}
	}
	return 0
}

func getEmptySpots(board [9]string) []int {
	var result []int

	for index, item := range board {
		if item == " " {
			result = append(result, index)
		}
	}
	return result
}

func computerMove(board [9]string) [9]string {
	printBoard(board)
	best := findBestComputerMove(board)
	board[best] = "o"
	return board
}

func playerMove(board [9]string) [9]string {
	fmt.Printf("Possibilities: %v\n", getEmptySpots(board))
	printBoard(board)

	fmt.Print("where to move? ")
	var index int
	fmt.Scan(&index)

	for {
		if index < 9 && index >= 0 && board[index] == " " {
			board[index] = "x"
			break
		} else {
			fmt.Print("where to move? ")
			fmt.Scan(&index)
		}
	}

	return board
}

func checkThree(a string, b string, c string) bool {
	if a == " " || b == " " || c == " " {
		return false
	}
	return a == b && b == c
}

func isGameOver(board [9]string) (bool, string) {
	if len(getEmptySpots(board)) == 0 {
		return true, " "
	}
	if checkThree(board[0], board[1], board[2]) {
		return true, board[0]
	}
	if checkThree(board[3], board[4], board[5]) {
		return true, board[3]
	}
	if checkThree(board[6], board[7], board[8]) {
		return true, board[6]
	}
	if checkThree(board[0], board[3], board[6]) {
		return true, board[0]
	}
	if checkThree(board[1], board[4], board[7]) {
		return true, board[1]
	}
	if checkThree(board[2], board[5], board[8]) {
		return true, board[2]
	}
	if checkThree(board[0], board[4], board[8]) {
		return true, board[0]
	}
	if checkThree(board[6], board[4], board[2]) {
		return true, board[6]
	}
	return false, " "
}

func printBoard(board [9]string) {
	fmt.Println(board[0] + " | " + board[1] + " | " + board[2])
	fmt.Println("---------")
	fmt.Println(board[3] + " | " + board[4] + " | " + board[5])
	fmt.Println("---------")
	fmt.Println(board[6] + " | " + board[7] + " | " + board[8])
}

func printResult(winner string, board [9]string) {
	printEmptyLines(2)
	if winner == " " {
		fmt.Printf("--== TIE ==--\n")
	} else {
		fmt.Printf("--== %v won the game ==--\n", winner)
	}
	printEmptyLines(2)
	printBoard(board)
	printEmptyLines(2)
}

func printEmptyLines(num int) {
	for i := 0; i < num; i++ {
		fmt.Println()
	}
}

func min(values []int) int {
	min := 0
	for i, e := range values {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}

func max(values []int) int {
	max := 0
	for i, e := range values {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}
