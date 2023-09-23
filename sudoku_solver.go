package main

import (
	"fmt"
	"strconv"
)

var (
	testBoarNineXNine = [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	testBoarFourXFour = [][]string{
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	boardSize  = len(testBoarNineXNine)
	countSquar = 3
	sizeSquar  = boardSize / countSquar
)

const (
	emptyCell string = "."
)

func main() {
	printBoard(testBoarNineXNine)
	if solveSudoku(testBoarNineXNine) {
		fmt.Println("Решил!")
		printBoard(testBoarNineXNine)
		return
	}
	fmt.Println("Не решается чета")
}

func solveSudoku(board [][]string) bool {
	currentPositionX, currentPositionY, flag := searchEmpty(board)
	if !flag {
		return true
	}

	for num := 1; num < boardSize+1; num++ {
		currentNum := strconv.Itoa(num)

		if checkValidate(board, currentNum, currentPositionX, currentPositionY) {
			board[currentPositionX][currentPositionY] = currentNum
			if solveSudoku(board) {
				return true
			}
			board[currentPositionX][currentPositionY] = emptyCell
		}
	}

	return false
}

func checkValidate(board [][]string, num string, positionX int, positionY int) bool {

	//check cols
	for x := 0; x < boardSize; x++ {
		if board[x][positionY] == num && x != positionX {
			return false
		}
	}

	//check rows
	for y := 0; y < boardSize; y++ {
		if board[positionX][y] == num && y != positionY {
			return false
		}
	}

	//check box
	boxPositionX, boxPositionY := positionitionSquar(positionX, positionY)
	for x := boxPositionX; x < boxPositionX+sizeSquar; x++ {
		for y := boxPositionY; y < boxPositionY+sizeSquar; y++ {
			if board[x][y] == num && x != positionX && y != positionY {
				return false
			}
		}
	}
	return true
}

func positionitionSquar(positionX int, positionY int) (int, int) {
	return (positionX / countSquar) * countSquar, (positionY / countSquar) * countSquar
}

func searchEmpty(board [][]string) (int, int, bool) {
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			if board[x][y] == "." {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

func printBoard(board [][]string) {
	fmt.Printf("\n-------------------\n")
	for x := 0; x < boardSize; x++ {
		fmt.Printf("|")
		for y := 0; y < boardSize; y++ {
			fmt.Printf(board[x][y])
			fmt.Printf("|")
		}
		fmt.Printf("\n-------------------\n")
	}
}
