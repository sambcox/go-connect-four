package board

import (
	"fmt"
	"github.com/sambcox/go-connect-four/cell"
)

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
	for _, column := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		var cells []cell.Cell
		for i := 0; i < 6; i++ {
			cells = append(cells, cell.NewCell())
		}
		b.Columns[column] = cells
	}
}

func (b *Board) PrintBoard() {
	fmt.Println("  A B C D E F G")
	for row := 5; row >= 0; row-- {
		for _, column := range []string{"A", "B", "C", "D", "E", "F", "G"} {
			fmt.Printf("%s ", b.Columns[column][row].Piece)
		}
		fmt.Println()
	}
}

func (b *Board) HorizontalWin() string {
	for row := 0; row < 6; row++ {
		consecutive := []string{}
		for _, column := range []string{"A", "B", "C", "D", "E", "F", "G"} {
			consecutive = append(consecutive, b.Columns[column][row].Piece)
		}
		if containsWin(consecutive) {
			return consecutive[0]
		}
	}
	return ""
}

func (b *Board) VerticalWin() string {
	for _, column := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		consecutive := []string{}
		for row := 0; row < 6; row++ {
			consecutive = append(consecutive, b.Columns[column][row].Piece)
		}
		if containsWin(consecutive) {
			return consecutive[0]
		}
	}
	return ""
}

func (b *Board) DiagonalWin() string {
	consecutiveDiagonals := [][]string{}

	for startRow := 0; startRow <= 2; startRow++ {
		for startCol := 0; startCol <= 3; startCol++ {
			var diagonal []string
			for i := 0; i < 4; i++ {
				diagonal = append(diagonal, b.Columns[getColumnKey(startCol+i)][startRow+i].Piece)
			}
			consecutiveDiagonals = append(consecutiveDiagonals, diagonal)
		}
	}

	for startRow := 0; startRow <= 2; startRow++ {
		for startCol := 3; startCol <= 6; startCol++ {
			var diagonal []string
			for i := 0; i < 4; i++ {
				diagonal = append(diagonal, b.Columns[getColumnKey(startCol-i)][startRow+i].Piece)
			}
			consecutiveDiagonals = append(consecutiveDiagonals, diagonal)
		}
	}

	for _, diagonal := range consecutiveDiagonals {
		if containsWin(diagonal) {
			return diagonal[0]
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
	for _, column := range b.Columns {
		if !column[5].IsEmpty() {
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

func getColumnKey(colIndex int) string {
	columns := []string{"A", "B", "C", "D", "E", "F", "G"}
	return columns[colIndex]
}
