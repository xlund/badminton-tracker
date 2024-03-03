package game

import "strings"

type Player struct {
	Name string
}

func (p Player) LowerName() string {
	return strings.ToLower(p.Name)
}
