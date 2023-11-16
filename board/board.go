package board

import (
	"fmt"
	"github.com/sambcox/go-connect-four/cell"
)

var columns = []string{"A", "B", "C", "D", "E", "F", "G"}

type Board struct {
	Columns map[string][]cell.Cell
}

func NewBoard() Board {
	b := Board{
		Columns: make(map[string][]cell.Cell),
	}
	b.initializeColumns()
	return b
}

func (b *Board) initializeColumns() {
	for _, column := range columns {
		var cells []cell.Cell
		for i := 0; i < 6; i++ {
			cells = append(cells, cell.NewCell())
		}
		b.Columns[column] = cells
	}
}

func (b *Board) PrintBoard() {
	fmt.Println("A B C D E F G")
	for row := 5; row >= 0; row-- {
		for _, column := range columns {
			fmt.Printf("%s ", b.Columns[column][row].Piece)
		}
		fmt.Println()
	}
}

func (b *Board) HorizontalWin() string {
	for row := 0; row < 6; row++ {
		consecutive := []string{}
		for _, column := range columns {
			consecutive = append(consecutive, b.Columns[column][row].Piece)
			if len(consecutive) >= 4 && containsWin(consecutive[len(consecutive)-4:]) {
				return consecutive[0]
			}
		}
	}
	return ""
}

func (b *Board) VerticalWin() string {
	for _, column := range columns {
		consecutive := []string{}
		for row := 0; row < 6; row++ {
			consecutive = append(consecutive, b.Columns[column][row].Piece)
			if len(consecutive) >= 4 && containsWin(consecutive[len(consecutive)-4:]) {
				return consecutive[0]
			}
		}
	}
	return ""
}

func (b *Board) DiagonalWin() string {
	for startRow := 0; startRow <= 2; startRow++ {
		for startCol := 0; startCol <= 3; startCol++ {
			consecutive := []string{}
			for i := 0; i < 4; i++ {
				consecutive = append(consecutive, b.Columns[columns[startCol+i]][startRow+i].Piece)
				if len(consecutive) == 4 && containsWin(consecutive) {
					return consecutive[0]
				}
			}
		}
	}

	for startRow := 0; startRow <= 2; startRow++ {
		for startCol := 3; startCol <= 6; startCol++ {
			consecutive := []string{}
			for i := 0; i < 4; i++ {
				consecutive = append(consecutive, b.Columns[columns[startCol-i]][startRow+i].Piece)
				if len(consecutive) == 4 && containsWin(consecutive) {
					return consecutive[0]
				}
			}
		}
	}

	return ""
}

func (b *Board) WinGame() string {
	if winner := b.HorizontalWin(); winner != "" {
		return winner
	} else if winner := b.VerticalWin(); winner != "" {
		return winner
	} else if winner := b.DiagonalWin(); winner != "" {
		return winner
	}
	return ""
}

func (b *Board) EndGame() bool {
	for _, column := range columns {
		if !b.Columns[column][5].IsEmpty() {
			return true
		}
	}
	return false
}

func containsWin(consecutive []string) bool {
	for i := 0; i <= len(consecutive)-4; i++ {
		if consecutive[i] == consecutive[i+1] &&
			consecutive[i] == consecutive[i+2] &&
			consecutive[i] == consecutive[i+3] &&
			consecutive[i] != "." {
			return true
		}
	}
	return false
}
