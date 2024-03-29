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

type GameList []Game

func (g GameList) Filter(f func(Game) bool) GameList {
	var games GameList
	for _, game := range g {
		if f(game) {
			games = append(games, game)
		}
	}
	return games
}

type Game struct {
	ID       int
	Date     time.Time
	Teams    [2]Team
	GameType GameType
	Result   Result
	Target   int
}

type Result struct {
	Winner   Team
	Loser    Team
	Walkover bool
	Draw     bool
}

func (g Game) HasResult() bool {
	return g.Result.Winner != Team{}
}

func (g Game) IsTiebreak() bool {
	return g.Result.Winner.Score >= g.Target && g.Result.Loser.Score >= g.Target
}

func (g Game) ResultString() string {
	if !g.HasResult() {
		return "No result"
	}
	result := g.Result
	if g.GameType == "singles" {
		return fmt.Sprintf("------\n%s beat %s\nScore: %v—%v\n------\n\n", result.Winner.P1.Name, result.Loser.P1.Name, result.Winner.Score, result.Loser.Score)
	}
	return fmt.Sprintf("------\n{%s} beat {%s}\nScore: %v—%v\n------\n\n", result.Winner, result.Loser, result.Winner.Score, result.Loser.Score)
}
