package main

import (
	"fmt"
	"strconv"
	"strings"
)

func display(s [8][8]string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if s[i][j] == "" {
				fmt.Print("-  ")
			} else {
				fmt.Print(s[i][j])
			}
		}
		fmt.Println(8 - i)
	}
	fmt.Println("a ", "b ", "c ", "d ", "e ", "f ", "g ", "h ")
}
func rightLetter(x string) bool {
	if x >= "a" && x <= "h" {
		return true
	}
	return false
}
func rightDigit(x string) bool {
	y, _ := strconv.Atoi(x)
	if y > 0 && y < 9 {
		return true
	}
	return false
}
func rightPlayer(input string, board [8][8]string, player int) bool {
	x, y := inputToCoords(input)
	if (player == 1 && strings.Contains(board[y][x], "w")) || (player == 2 && strings.Contains(board[y][x], "b")) {
		return true
	}
	return false
}
func letterToNumber(x string) int {
	var y int
	switch x {
	case "a":
		y = 1
	case "b":
		y = 2
	case "c":
		y = 3
	case "d":
		y = 4
	case "e":
		y = 5
	case "f":
		y = 6
	case "g":
		y = 7
	case "h":
		y = 8
	}
	return y
}
func pieceHere(input string, board [8][8]string) bool {
	x, y := inputToCoords(input)
	fmt.Println(x+1, 8-y, board[y][x])
	if board[y][x] != "" {
		return true
	}
	return false
}
func inputToCoords(input string) (int, int) {
	x := letterToNumber(string(input[0])) - 1
	y, _ := strconv.Atoi(string(input[1]))
	y = 8 - y
	return x, y
}
func validPiece(input *string, board [8][8]string, player int) {
	var x string
	for {
		fmt.Scan(&x)
		if len(x) != 2 || !rightLetter(string(x[0])) || !rightDigit(string(x[1])) || !pieceHere(x, board) || !rightPlayer(x, board, player) {
			fmt.Println("Error")
		} else {
			*input = x
			fmt.Println("right")
			return
		}
	}
}

func selectPiece(input string, board *[8][8]string) {
	x, y := inputToCoords(input)
	board[y][x] = strings.Replace(board[y][x], " ", "<", 1)

}
func unSelectPiece(input string, board *[8][8]string) {
	x, y := inputToCoords(input)
	board[y][x] = strings.Replace(board[y][x], "<", " ", 1)
}
func validDo(input string, do *string, board *[8][8]string) {
	var x string
	for {
		fmt.Scan(&x)
		if x == "back" {
			unSelectPiece(input, &*board)
			return
		}
		if len(x) != 2 || !rightLetter(string(x[0])) || !rightDigit(string(x[1])) {
			fmt.Println("Error")
		} else {
			*do = x
			fmt.Println("right")
			return
		}
	}
}
func switchPlayer(player *int) {
	if *player == 1 {
		*player = 2
		return
	}
	*player = 1
	return
}

// func validMove()
func main() {
	var input, do string
	var player = 1
	var board [8][8]string
	board[0] = [8]string{"bR ", "bH ", "bB ", "bQ ", "bK ", "bB ", "bH ", "bR "}
	board[1] = [8]string{"bP ", "bP ", "bP ", "bP ", "bP ", "bP ", "bP ", "bP "}
	board[6] = [8]string{"wP ", "wP ", "wP ", "wP ", "wP ", "wP ", "wP ", "wP "}
	board[7] = [8]string{"wR ", "wH ", "wB ", "wQ ", "wK ", "wB ", "wH ", "wR "}
	for {
		display(board)
		validPiece(&input, board, player)
		selectPiece(input, &board)
		display(board)
		validDo(input, &do, &board)
		if do != "back" {
			switchPlayer(&player)
		}
	}
}
