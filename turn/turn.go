package turn

import (
	"fmt"
	"github.com/sambcox/go-connect-four/board"
	"math/rand"
	"strings"
	"time"
)

type Turn struct {
	Board *board.Board
}

func NewTurn(board *board.Board) *Turn {
	return &Turn{
		Board: board,
	}
}

func (t *Turn) UserTakeTurn() {
	fmt.Print("Please enter a letter between A and G: ")
	userInput := getUserInput()
	if strings.Contains("ABCDEFG", strings.ToUpper(userInput)) {
		t.UserPlacePiece(strings.ToUpper(userInput))
	} else {
		fmt.Println("That is an invalid input! Please select a letter between A and G.")
		t.UserTakeTurn()
	}
}

func (t *Turn) TwoPlayerTakeTurn() {
	fmt.Print("Please enter a letter between A and G: ")
	userInput := getUserInput()
	if strings.Contains("ABCDEFG", strings.ToUpper(userInput)) {
		t.PlacePiece(strings.ToUpper(userInput))
	} else {
		fmt.Println("That is an invalid input! Please select a letter between A and G.")
		t.TwoPlayerTakeTurn()
	}
}

func (t *Turn) ComputerTakeTurn() {
	rand.Seed(time.Now().UnixNano())
	columns := []string{"A", "B", "C", "D", "E", "F", "G"}
	shuffledColumns := shuffle(columns)
	computerInput := shuffledColumns[0]
	t.PCPlacePiece(computerInput)
}

func (t *Turn) UserPlacePiece(columnInputted string) {
	if !t.Board.Columns[columnInputted][5].IsEmpty() {
		fmt.Println("That column is full! Please select another.")
		t.TwoPlayerTakeTurn()
	} else {
		for i, cell := range t.Board.Columns[columnInputted] {
			if cell.IsEmpty() {
				t.Board.Columns[columnInputted][i].AddPiece()
				break
			}
		}
	}
}

func (t *Turn) PCPlacePiece(columnInputted string) {
	if !t.Board.Columns[columnInputted][5].IsEmpty() {
		t.ComputerTakeTurn()
	} else {
		for i, cell := range t.Board.Columns[columnInputted] {
			if cell.IsEmpty() {
				t.Board.Columns[columnInputted][i].ComputerAddPiece()
				break
			}
		}
	}
}

func (t *Turn) PlacePiece(columnInputted string) {
	if !t.Board.Columns[columnInputted][5].IsEmpty() {
		fmt.Println("That column is full! Please select another.")
		t.UserTakeTurn()
	} else {
		for i, cell := range t.Board.Columns[columnInputted] {
			if cell.IsEmpty() {
				t.Board.Columns[columnInputted][i].ComputerAddPiece()
				break
			}
		}
	}
}

func getUserInput() string {
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

func shuffle(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
