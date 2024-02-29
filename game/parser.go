package game

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func CsvParser(fp string) []Game {
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
	id := ParseInt(row[0])
	date, err := time.Parse("YYYY-MM-DD", row[2])
	if err != nil {
		println(err.Error())
	}
	home := ParseTeam(row[4], row[5])
	away := ParseTeam(row[6], row[7])
	gameType := ParseGameType(home, away)
	score := Score{
		home:   ParseInt(row[8]),
		away:   ParseInt(row[9]),
		target: ParseInt(row[3]),
	}
	walkover := ParseBool(row[11])
	isTournament := ParseBool(row[13])
	result := calculateWinner(score)

	return Game{
		id,
		date,
		gameType,
		home,
		away,
		score,
		walkover,
		isTournament,
		result}
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func ParseBool(b string) bool {
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

func ParseGameType(home Team, away Team) GameType {
	if home.playerTwo == (Player{}) && away.playerTwo == (Player{}) {
		return Singles
	}
	return Doubles
}

func calculateWinner(score Score) Winner {
	if score.home < score.target && score.away < score.target {
		return None
	}
	if score.home > score.away {
		return Home
	}
	return Away
}
