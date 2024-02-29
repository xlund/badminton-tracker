package game

import "strings"

type Player struct {
	name string
}

func (p Player) LowerName() string {
	return strings.ToLower(p.name)
}
