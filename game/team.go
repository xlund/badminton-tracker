package game

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Team struct {
	playerOne Player
	playerTwo Player
}

func (t Team) String() string {
	return fmt.Sprintf("%s, %s", t.playerOne.name, t.playerTwo.name)
}

func ParseTeam(playerOne string, playerTwo string) Team {
	caser := cases.Title(language.Swedish)
	if playerTwo == "" {
		return Team{
			playerOne: Player{
				name: caser.String(playerOne),
			},
			playerTwo: Player{},
		}
	}
	return Team{
		playerOne: Player{
			name: caser.String(playerOne),
		},
		playerTwo: Player{
			name: caser.String(playerTwo),
		},
	}
}
