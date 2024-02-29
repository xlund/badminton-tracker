package game

import (
	"fmt"
	"time"
)

const (
	Singles GameType = "singles"
	Doubles GameType = "doubles"
	Home    Winner   = "home"
	Away    Winner   = "away"
	None    Winner   = "none"
)

type GameType string
type Winner string

type Game struct {
	id           int
	date         time.Time
	gameType     GameType
	home         Team
	away         Team
	score        Score
	walkover     bool
	isTournament bool
	winner       Winner
}

type Score struct {
	home   int
	away   int
	target int
}

func (g Game) HasResult() bool {
	return g.winner != None
}

func (g Game) Winner() Team {
	if g.winner == Home {
		return g.home
	} else if g.winner == Away {
		return g.away
	}
	return Team{}
}

func (g Game) Loser() Team {
	if g.winner == Home {
		return g.away
	} else if g.winner == Away {
		return g.home
	}
	return Team{}
}

func (g Game) WinningScore() int {
	if g.winner == Home {
		return g.score.home
	} else if g.winner == Away {
		return g.score.away
	}
	return 0
}

func (g Game) LosingScore() int {
	if g.winner == Home {
		return g.score.away
	} else if g.winner == Away {
		return g.score.home
	}
	return 0
}

func (g Game) ResultString() string {
	if !g.HasResult() {
		return "No result"
	}
	if g.gameType == "singles" {
		return fmt.Sprintf("------\n%s beat %s\nScore: %v—%v\n------\n\n", g.Winner().playerOne.name, g.Loser().playerOne.name, g.WinningScore(), g.LosingScore())
	}
	return fmt.Sprintf("------\n{%s} beat {%s}\nScore: %v—%v\n------\n\n", g.Winner(), g.Loser(), g.WinningScore(), g.LosingScore())
}

func (m Game) WithPlayer(player Player) bool {
	return m.home.playerOne == player || m.home.playerTwo == player || m.away.playerOne == player || m.away.playerTwo == player
}
