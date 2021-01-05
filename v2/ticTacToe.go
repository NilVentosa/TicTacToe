package main

import (
	"fmt"
	"strconv"
)

var positions = map[int16]int16 {
	0:		int16(0b000000001), 1:		int16(0b000000010), 2:		int16(0b000000100),
	3: 		int16(0b000001000), 4:		int16(0b000010000), 5:		int16(0b000100000),
	6:		int16(0b001000000), 7:		int16(0b010000000), 8:		int16(0b100000000),
}

var winningPositions = map[int16]struct{}{
	0b000000111: {}, 0b000111000: {}, 0b111000000: {}, 0b001001001: {},
	0b010010010: {}, 0b100100100: {}, 0b100010001: {}, 0b101000100: {},
}

func main() {
	play()
}

func play() {
	isPlayerTurn := true
	var playerBoard int16 = 0
	var computerBoard int16 = 0
	for {
		if isGameOver(playerBoard, computerBoard) {
			printResult(playerBoard, computerBoard)
			break
		} else {
			if isPlayerTurn {
				printBoard(playerBoard, computerBoard)
				cell := getPlayerInput(playerBoard | computerBoard)
				playerBoard = playerBoard | positions[cell]
				isPlayerTurn = !isPlayerTurn
			} else {
				computerBoard = computerMove(playerBoard, computerBoard)
				isPlayerTurn = !isPlayerTurn
			}
		}
	}
}

func computerMove(playerBoard int16, computerBoard int16) int16 {
	var best = 1
	var bestIndex int16
	emptySpots := getEmptySpots(playerBoard | computerBoard)

	for _, item := range emptySpots {
		updatedBoard := computerBoard | item
		miniMaxValue := minimax(playerBoard, updatedBoard, len(emptySpots), true)
		if miniMaxValue < best {
			best = miniMaxValue
			bestIndex = item
		}
	}
	return positions[bestIndex] | computerBoard
}

func minimax(playerBoard int16, computerBoard int16, depth int, isMaximizer bool) int {
	if isGameOver(playerBoard, computerBoard) || depth == 0 {
		return evaluate(playerBoard, computerBoard)
	}

	var updatedBoard int16
	if isMaximizer {
		maxEval := -1
		for _, item := range getEmptySpots(playerBoard | computerBoard) {
			updatedBoard = playerBoard
			updatedBoard = updatedBoard | positions[item]
			eval := minimax(updatedBoard, computerBoard, depth-1, false)
			if eval > maxEval {
				maxEval = eval
			}
		}
		return maxEval
	}
	minEval := 1
	for _, item := range getEmptySpots(playerBoard | computerBoard) {
		updatedBoard = computerBoard
		updatedBoard = computerBoard | positions[item]
		eval := minimax(playerBoard, updatedBoard, depth-1, true)
		if eval < minEval {
			minEval = eval
		}
	}
	return minEval
}

func evaluate(playerBoard int16, computerBoard int16) int {
	if isBoardWinner(playerBoard) {
		return 1
	} else if isBoardWinner(computerBoard) {
		return -1
	}
	return 0
}

func getEmptySpots(board int16) []int16 {
	var emptySpots []int16
	for key, value := range positions {
		if board & value != value {
			emptySpots = append(emptySpots, key)
		}
	}
	return emptySpots
}

func printResult(playerBoard int16, computerBoard int16) {
	if isBoardWinner(playerBoard) {
		fmt.Println("--==The player won==--")
		printBoard(playerBoard, computerBoard)
	} else if isBoardWinner(computerBoard) {
		fmt.Println("--==The computer won==--")
		printBoard(playerBoard, computerBoard)
	} else if isBoardTied(playerBoard | computerBoard) {
		fmt.Println("--==It is a tie==--")
		printBoard(playerBoard, computerBoard)
	}
}

func printBoard(playerBoard int16, computerBoard int16) {
	playerText := "000000000" + strconv.FormatInt(int64(playerBoard), 2)
	playerText = playerText[len(playerText)-9:]

	computerText := "000000000" + strconv.FormatInt(int64(computerBoard), 2)
	computerText = computerText[len(computerText)-9:]

	for i := 8; i >= 0; i-- {
		if (i+1) % 3 == 0 && i != 8 {
			fmt.Println("\n---------")
		}
		if (i+1) % 3 != 0  {
			fmt.Print(" | ")
		}
		if string(computerText[i]) == "1" {
			fmt.Print("o")
		} else if string(playerText[i]) == "1" {
			fmt.Print("x")
		} else {
			fmt.Print("·")
		}
	}
	fmt.Println()
}

func getPlayerInput(board int16) int16 {
	for {
		fmt.Print("where to move? ")
		var cell int16
		_, err := fmt.Scan(&cell)
		_, ok := positions[cell]
		isCellFree := board & positions[cell] != positions[cell]
		if ok {
			if err == nil && isCellFree {
				return cell
			}
		}
	}
}

func isGameOver(playerBoard int16, computerBoard int16) bool {
	if isBoardWinner(playerBoard) || isBoardWinner(computerBoard) || isBoardTied(playerBoard | computerBoard) {
		return true
	}
	return false
}

func isBoardTied(board int16) bool {
	return board == 0b111111111
}

func isBoardWinner(board int16) bool {
	for winnerBoard, _ := range winningPositions {
		if board & winnerBoard == winnerBoard {
			return true
		}
	}
	return false
}
