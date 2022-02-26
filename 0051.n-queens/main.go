package main

import (
	"fmt"
)

// TODO
func solveNQueens(n int) [][]string {
	var board [][]string

	reSetCoordinates := func(row int) (beforeRow, start int) {
		for i := row; i >= 0; i-- {
			for j := 0; j < n; j++ {
				if board[i][j] == "Q" && j != n-1 {
					return i, j + 1
				}
			}
		}
		return 0, 0
	}

	var place func(row, start int)
	place = func(row, start int) {
		if row > n-1 {
			return
		}
		for j := start; j < n; j++ {
			board[row][j] = "Q"
			if isAttack(board, n) {
				board[row][j] = "."
				if j == n-1 {
					resetRow, resetStart := reSetCoordinates(row - 1)
					if resetRow == 0 && resetStart == 0 {
						return
					}
					board = resetBoard(board, resetRow, n)
					place(resetRow, resetStart)
				} else {
					continue
				}
			} else {
				break
			}
		}
		place(row+1, 0)
	}

	skip := -1

	queens := make([][]string, 0)

	for i := 0; i < n; i++ {
		if i <= skip {
			continue
		}
		board = generateBoard(n)
		place(0, i)
		fmt.Println(board)
		str, valid, skp := validBoard(board, n)
		if valid {
			queens = append(queens, str)
			skip = skp
		}
	}
	return queens
}

func validBoard(board [][]string, n int) ([]string, bool, int) {
	str := make([]string, 0)
	skip := 0
	for i := 0; i < n; i++ {
		s := ""
		count := 0
		for j := 0; j < n; j++ {
			s += board[i][j]
			if board[i][j] == "Q" {
				count++
				if i == 0 {
					skip = j
				}
			}
		}
		str = append(str, s)
		if count == 0 {
			return str, false, 0
		}
	}
	return str, true, skip
}

func resetBoard(board [][]string, row, n int) [][]string {
	for i := row; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	return board
}

func isAttack(board [][]string, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == "Q" {
				for a := 0; a < n; a++ {
					if board[i][a] == "Q" && a != j {
						return true
					}
					if board[a][j] == "Q" && a != i {
						return true
					}
				}

				for a, b := i-1, j-1; a >= 0 && b >= 0; a, b = a-1, b-1 {
					if board[a][b] == "Q" {
						return true
					}
				}

				for a, b := i+1, j+1; a < n && b < n; a, b = a+1, b+1 {
					if board[a][b] == "Q" {
						return true
					}
				}

				for a, b := i+1, j-1; a < n && b >= 0; a, b = a+1, b-1 {
					if board[a][b] == "Q" {
						return true
					}
				}

				for a, b := i-1, j+1; a >= 0 && b < n; a, b = a-1, b+1 {
					if board[a][b] == "Q" {
						return true
					}
				}
			}
		}
	}
	return false
}

func generateBoard(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	return board
}

func main() {
	n := 9
	board := solveNQueens(n)
	fmt.Println(board)
}
