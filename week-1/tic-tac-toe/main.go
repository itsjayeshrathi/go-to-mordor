package main

import (
	"fmt"
	"os"
)

type Value struct {
	X string
	O string
}

var Symbols = Value{X: "X", O: "O"}

var Grid = [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

var Players = map[string]string{"player1": Symbols.X, "player2": Symbols.O}

func Display() {
	for _, row := range Grid {
		for _, ele := range row {
			fmt.Print(ele, " ")
		}
		fmt.Println()
	}
}

func Insert(value string, key string) bool {
	for i, row := range Grid {
		for j, ele := range row {
			if value == ele {
				val := Players[key]
				Grid[i][j] = val
				return true
			}
		}
	}
	return false
}

func InsertVal(stopper int) {
	fmt.Println("Your postion should be from 1 to 9 from grid")
	value := ""
	if stopper%2 == 0 {

		for range 3 {
			fmt.Println("Player 1 choose the postion.")
			fmt.Scanln(&value)
			checker := Insert(value, "player1")
			if checker {
				return
			}
		}
	} else {

		for range 3 {
			fmt.Println("Player 2 choose the postion.")
			fmt.Scanln(&value)
			checker := Insert(value, "player2")
			if checker {
				return
			}
		}

	}
}

func CheckGrid(stopper int) {
	if stopper < 2 {
		return
	}

	check, val := false, ""

	// row wise check

	for _, row := range Grid {
		if row[0] == row[1] && row[1] == row[2] {
			check = true
			val = row[0]
			break
		}
	}

	// col wise check
	for j := range 3 {
		if Grid[0][j] == Grid[1][j] && Grid[1][j] == Grid[2][j] {
			check = true
			val = Grid[0][j]
			break
		}
	}

	// diagonal wise check
	if (Grid[0][0] == Grid[1][1] && Grid[1][1] == Grid[2][2]) || (Grid[0][2] == Grid[1][1] && Grid[1][1] == Grid[2][0]) {
		check = true
		val = Grid[1][1]
	}

	if check || stopper == 8 {
		if val == "X" {
			fmt.Println("player1 won!")
		} else if val == "O" {
			fmt.Println("player2 won!")
		} else if stopper == 8 {
			fmt.Println("Draw!")
		}
		os.Exit(0)
	}
}

func main() {
	stopper := 0

	fmt.Println("TIC-TAC-TOE")

	for stopper < 9 {
		InsertVal(stopper)
		Display()
		CheckGrid(stopper)
		stopper++
	}

}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// type TicTacToe struct {
// 	Grid    [][]string
// 	Players map[string]string
// 	Turn    int
// }

// func NewGame() *TicTacToe {
// 	return &TicTacToe{
// 		Grid: [][]string{
// 			{"1", "2", "3"},
// 			{"4", "5", "6"},
// 			{"7", "8", "9"},
// 		},
// 		Players: map[string]string{"player1": "X", "player2": "O"},
// 		Turn:    0,
// 	}
// }

// func (t *TicTacToe) Display() {
// 	fmt.Println("\nCurrent Grid:")
// 	for _, row := range t.Grid {
// 		fmt.Println(row)
// 	}
// 	fmt.Println()
// }

// func (t *TicTacToe) Insert(value string) bool {
// 	for i, row := range t.Grid {
// 		for j, ele := range row {
// 			if ele == value {
// 				t.Grid[i][j] = t.Players[t.CurrentPlayer()]
// 				return true
// 			}
// 		}
// 	}
// 	fmt.Println("Invalid move! Try again.")
// 	return false
// }

// func (t *TicTacToe) CurrentPlayer() string {
// 	if t.Turn%2 == 0 {
// 		return "player1"
// 	}
// 	return "player2"
// }

// func (t *TicTacToe) GetUserInput() {
// 	var value string
// 	for {
// 		fmt.Printf("%s, choose a position (1-9): ", t.CurrentPlayer())
// 		fmt.Scanln(&value)
// 		if _, err := strconv.Atoi(value); err == nil && t.Insert(value) {
// 			break
// 		}
// 		fmt.Println("Invalid input! Please enter a number between 1-9 that is not already taken.")
// 	}
// }

// func (t *TicTacToe) CheckWinner() {
// 	grid := t.Grid
// 	winPatterns := [][][2]int{
// 		{{0, 0}, {0, 1}, {0, 2}}, // Rows
// 		{{1, 0}, {1, 1}, {1, 2}},
// 		{{2, 0}, {2, 1}, {2, 2}},
// 		{{0, 0}, {1, 0}, {2, 0}}, // Columns
// 		{{0, 1}, {1, 1}, {2, 1}},
// 		{{0, 2}, {1, 2}, {2, 2}},
// 		{{0, 0}, {1, 1}, {2, 2}}, // Diagonals
// 		{{0, 2}, {1, 1}, {2, 0}},
// 	}

// 	for _, pattern := range winPatterns {
// 		if grid[pattern[0][0]][pattern[0][1]] == grid[pattern[1][0]][pattern[1][1]] &&
// 			grid[pattern[1][0]][pattern[1][1]] == grid[pattern[2][0]][pattern[2][1]] {
// 			fmt.Printf("%s won!\n", t.CurrentPlayer())
// 			os.Exit(0)
// 		}
// 	}

// 	if t.Turn == 8 {
// 		fmt.Println("It's a draw!")
// 		os.Exit(0)
// 	}
// }

// func main() {
// 	game := NewGame()
// 	fmt.Println("Welcome to Tic-Tac-Toe!")
// 	game.Display()

// 	for game.Turn < 9 {
// 		game.GetUserInput()
// 		game.Display()
// 		game.CheckWinner()
// 		game.Turn++
// 	}
// }
