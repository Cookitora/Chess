package main

import (
	"fmt"
	"math"
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
func validInput(x string) bool {
	if x[0] >= 'a' && x[0] <= 'h' && x[1] <= '8' && x[1] >= '1' {
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
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
	return m[x]
}
func pieceHere(input string, board [8][8]string) bool {
	x, y := inputToCoords(input)
	//fmt.Println(x+1, 8-y, board[y][x])
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
		if len(x) != 2 || !validInput(x) || !pieceHere(x, board) || !rightPlayer(x, board, player) {
			fmt.Println("Error")
		} else {
			*input = x
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
func validDo(input string, do *string, board *[8][8]string, player int) {
	var x string
	for {
		fmt.Scan(&x)
		if x == "back" {
			*do = x
			unSelectPiece(input, &*board)
			return
		}
		if len(x) != 2 || !validInput(x) || !validMove(input, x, *board, player) {
			fmt.Println("Error")
		} else {
			*do = x
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
func validMove(input string, do string, board [8][8]string, player int) bool {
	x1, y1 := inputToCoords(input)
	x2, y2 := inputToCoords(do)
	p := board[y1][x1]
	switch p {
	case "bH<", "wH<": //конь
		{
			if ((math.Abs(float64(x1-x2)) == 2 && math.Abs(float64(y1-y2)) == 1) || (math.Abs(float64(x1-x2)) == 1 && math.Abs(float64(y1-y2)) == 2)) && !rightPlayer(do, board, player) { // нужно сделать функцию myPiece чтобы не есть свои, но есть чужие
				return true
			}
		}
	case "wP<":
		{

		}
	}
	return false
}
func movePiece(input string, do string, board *[8][8]string) {
	x1, y1 := inputToCoords(input)
	x2, y2 := inputToCoords(do)
	board[y2][x2] = board[y1][x1]
	board[y1][x1] = ""
	unSelectPiece(do, &*board)
}
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
		validDo(input, &do, &board, player)
		if do != "back" {
			movePiece(input, do, &board)
			switchPlayer(&player)
		}
	}
}

//TODO: хочу имплементировать ASCII символы шахмат шоб красивее было, но это чуть позже мне сегодня лень xD
