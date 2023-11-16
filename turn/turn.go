package turn

import (
	"fmt"
	"github.com/sambcox/go-connect-four/board"
	"math/rand"
	"strings"
	"time"
)

const validColumns = "ABCDEFG"

type Turn struct {
	Board *board.Board
}

func NewTurn(board *board.Board) *Turn {
	return &Turn{
		Board: board,
	}
}

func (t *Turn) TakeTurn(isPlayerTwo bool) {
	for {
		fmt.Print("Please enter a letter between A and G: ")
		userInput := getUserInput()
		if strings.Contains(validColumns, strings.ToUpper(userInput)) {
			if isPlayerTwo {
				t.PlacePiece(strings.ToUpper(userInput), "user2")
			} else {
				t.PlacePiece(strings.ToUpper(userInput), "user")
			}
			break
		} else {
			fmt.Println("That is an invalid input! Please select a letter between A and G.")
		}
	}
}

func (t *Turn) ComputerTakeTurn() {
	rand.Seed(time.Now().UnixNano())
	columns := []string{"A", "B", "C", "D", "E", "F", "G"}
	shuffledColumns := shuffle(columns)
	computerInput := shuffledColumns[0]
	t.PlacePiece(computerInput, "computer")
}

func (t *Turn) PlacePiece(columnInputted string, pieceType string) {
	if !t.Board.Columns[columnInputted][5].IsEmpty() {
		fmt.Println("That column is full! Please select another.")
		if pieceType == "user" {
			t.TakeTurn(false)
		} else if pieceType == "computer" {
			t.ComputerTakeTurn()
		} else if pieceType == "user2" {
			t.TakeTurn(true)
		}
	} else {
		for i, cell := range t.Board.Columns[columnInputted] {
			if cell.IsEmpty() {
				if pieceType == "user" {
					t.Board.Columns[columnInputted][i].AddPiece()
				} else if pieceType == "computer" || pieceType == "user2" {
					t.Board.Columns[columnInputted][i].ComputerAddPiece()
				}
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
