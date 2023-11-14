package cell

type Cell struct {
	Piece string
}

func NewCell() Cell {
	return Cell{Piece: "."}
}

func (c *Cell) AddPiece() {
	c.Piece = "X"
}

func (c *Cell) ComputerAddPiece() {
	c.Piece = "O"
}

func (c *Cell) IsEmpty() bool {
	return c.Piece == "."
}
