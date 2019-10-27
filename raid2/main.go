package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 9 {
		fmt.Println("Error")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}

		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				if ch != '.' {
					fmt.Println("Error")
					return
				}
			}
		}
	}

	arr := os.Args[1:]
	StrFromArr := make([]rune, 81)
	i := 0
	for _, value_arr := range arr {
		for k := 0; k <= 8; k++ {
			if value_arr[k] == '.' {
				StrFromArr[i] = '0'
				i++
			} else {
				StrFromArr[i] = rune(value_arr[k])
				i++
			}

		}
	}

	board := parseInput(StrFromArr)
	//	printBoard(board)

	if backtrack(&board) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(board)
	} else {
		fmt.Println("Error")
	}

}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return !hasEmptyCell(board)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input []rune) [9][9]int {
	board := [9][9]int{}
	i := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			i1 := Atoi(input[i])
			i++
			board[row][col] = i1
		}
	}
	return board
}

func Atoi(s rune) int {
	dec := 0
	//	for i, j := range s {
	ed := 0
	if s < '0' || s > '9' {
		return 0
	}
	for i := '1'; i <= s; i++ {
		ed = ed + 1
	}
	dec = dec*10 + ed
	//	}
	return dec
}
