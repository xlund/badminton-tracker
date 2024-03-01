package game

import (
	"fmt"
)

type Team struct {
	P1    Player
	P2    Player
	Score int
}

func (t Team) String() string {
	return fmt.Sprintf("%s, %s", t.P1.name, t.P2.name)
}
