package game

import (
	"fmt"
	"github.com/sambcox/go-connect-four/board"
	"github.com/sambcox/go-connect-four/player"
	"github.com/sambcox/go-connect-four/turn"
	"os"
)

type Game struct {
	Player1 *player.Player
	Player2 *player.Player
	Board   *board.Board
	Turn    *turn.Turn
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) MainMenu() {
	fmt.Println("Welcome to Connect Four!")
	fmt.Println("To play against PC, press c. To play with a friend, press p. To quit, press q.")

	for {
		wantToPlay := g.getUserInput()
		switch wantToPlay {
		case "c":
			g.Start()
			return
		case "p":
			g.TwoPlayerStart()
			return
		case "q":
			g.QuitGame()
			return
		default:
			fmt.Println("Invalid input, please press p or q")
		}
	}
}

func (g *Game) createPlayer(playerNumber int) error {
	fmt.Printf("Please enter player %d name\n", playerNumber)
	player, err := player.NewPlayer(g.getUserInput())
	if err != nil {
		return fmt.Errorf("error creating player %d: %w", playerNumber, err)
	}
	if playerNumber == 1 {
		g.Player1 = player
	} else {
		g.Player2 = player
	}
	return nil
}

func (g *Game) TwoPlayerStart() {
	if err := g.createPlayer(1); err != nil {
		fmt.Println(err)
		return
	}
	if err := g.createPlayer(2); err != nil {
		fmt.Println(err)
		return
	}

	board := board.NewBoard()
	g.Board = &board
	g.Turn = turn.NewTurn(g.Board)
	g.Board.PrintBoard()
	g.Player1TakeTurn()
}

func (g *Game) Start() {
	board := board.NewBoard()
	g.Board = &board
	g.Turn = turn.NewTurn(g.Board)
	g.Board.PrintBoard()
	g.GameUserTakeTurn()
}

func (g *Game) checkGameStatus() {
	if winner := g.Board.WinGame(); winner != "" {
		g.OverallWinGame(winner)
	} else if g.Board.EndGame() {
		g.DrawGame()
	}
	fmt.Println("--------------------------------")
}

func (g *Game) GameUserTakeTurn() {
	g.checkGameStatus()
	g.Turn.TakeTurn(false)
	fmt.Println("--------------------------------")
	g.Board.PrintBoard()
	g.GamePCTakeTurn()
}

func (g *Game) GamePCTakeTurn() {
	g.checkGameStatus()
	g.Turn.ComputerTakeTurn()
	g.Board.PrintBoard()
	g.GameUserTakeTurn()
}

func (g *Game) Player1TakeTurn() {
	g.checkGameStatus()
	fmt.Println("--------------------------------")
	fmt.Printf("%s, your turn\n", g.Player1.Name)
	fmt.Println("--------------------------------")
	g.Turn.TakeTurn(false)
	g.Board.PrintBoard()
	g.Player2TakeTurn()
}

func (g *Game) Player2TakeTurn() {
	g.checkGameStatus()
	fmt.Println("--------------------------------")
	fmt.Printf("%s, your turn\n", g.Player2.Name)
	fmt.Println("--------------------------------")
	g.Turn.TakeTurn(true)
	g.Board.PrintBoard()
	g.Player1TakeTurn()
}

func (g *Game) PlayAgain() {
	fmt.Println("To play against PC, press c. To play with a friend, press p. To quit, press q.")

	wantToPlay := g.getUserInput()
	switch wantToPlay {
	case "c":
		g.Start()
	case "p":
		g.TwoPlayerStart()
	case "q":
		g.QuitGame()
	default:
		fmt.Println("Invalid input, please press p or q")
		g.PlayAgain()
	}
}

func (g *Game) DrawGame() {
	fmt.Println("Thank you for playing! This game is a draw.")
	g.PlayAgain()
}

func (g *Game) OverallWinGame(winner string) {
	fmt.Println("--------------------------------")
	if winner == "X" {
		fmt.Println("You've won!")
	} else if winner == "O" {
		fmt.Println("You've lost!")
	}
	g.PlayAgain()
}

func (g *Game) PlayerWinGame() {
	var winner, loser *player.Player
	if winnerPiece := g.Board.WinGame(); winnerPiece == "X" {
		winner = g.Player1
		loser = g.Player2
	} else if winnerPiece == "O" {
		winner = g.Player2
		loser = g.Player1
	}

	fmt.Println("--------------------------------")
	fmt.Printf("Congratulations %s, you've won! Better luck next time, %s.\n", winner.Name, loser.Name)
	g.PlayAgain()
}

func (g *Game) QuitGame() {
	fmt.Println("--------------------------------")
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func (g *Game) getUserInput() string {
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
