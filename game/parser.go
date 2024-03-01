package game

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CsvParser(fp string) GameList {
	file, err := os.Open(fp)
	if err != nil {
		println(err.Error())
		return []Game{}
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	csv, err := reader.ReadAll()
	matches := []Game{}
	for _, row := range csv {
		match := FromCsvRow(row)
		matches = append(matches, match)
	}
	if err != nil {
		println(err.Error())
		return []Game{}
	}
	return matches
}

func FromCsvRow(row []string) Game {
	id := parseInt(row[0])
	date, _ := time.Parse("YYYY-MM-DD", row[2])
	// date := time.Now()

	t1 := parseTeam(row[4], row[5], parseInt(row[8]))
	t2 := parseTeam(row[6], row[7], parseInt(row[9]))
	target := parseInt(row[3])
	gameType := parseGameType(t1, t2)
	result := calculateResult(t1, t2, target)

	return Game{
		id,
		date,
		[2]Team{t1, t2},
		gameType,
		result,
	}

}
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func parseTeam(p1, p2 string, score int) Team {
	caser := cases.Title(language.Swedish)
	if p2 == "" {
		return Team{
			P1: Player{
				name: caser.String(p1),
			},
			P2:    Player{},
			Score: score,
		}
	}
	return Team{
		P1: Player{
			name: caser.String(p1),
		},
		P2: Player{
			name: caser.String(p2),
		},
		Score: score,
	}
}

func parseBool(b string) bool {
	if b == "YES" {
		return true
	} else if b == "NO" {
		return false
	}

	v, err := strconv.ParseBool(b)
	if err != nil {
		return false
	}
	return v
}

func parseGameType(home Team, away Team) GameType {
	if home.P2 == (Player{}) && away.P2 == (Player{}) {
		return Singles
	}
	return Doubles
}

func calculateResult(t1 Team, t2 Team, target int) Result {
	if t1.Score == target && t2.Score == target || t1.Score == t2.Score {
		return Result{
			Draw: true,
		}
	}
	if t1.Score == target {
		return Result{
			Winner: t1,
			Loser:  t2,
		}
	}
	return Result{
		Winner: t2,
		Loser:  t1,
	}
}
