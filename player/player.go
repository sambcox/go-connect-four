// Package player defines the Player type and related functions.
package player

import (
	"errors"
	"strings"
)

// Player represents a player in the game.
type Player struct {
	Name string
}

// NewPlayer creates a new Player with the given name.
// The name is converted to title case.
// Returns an error if the name is empty.
func NewPlayer(name string) (*Player, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &Player{
		Name: strings.Title(name),
	}, nil
}
