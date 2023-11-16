package player

import "strings"

type Player struct {
	Name string
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: strings.Title(name),
	}
}
