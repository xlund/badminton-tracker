package repos

import (
	"database/sql"
)

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlayerRepository interface {
	Migrate() error
	Create(player Player) (*Player, error)
	All() ([]Player, error)
	Find(id int) (*Player, error)
	Update(player Player) (*Player, error)
	Delete(id int) error
}

type PlayerRepo struct {
	db *sql.DB
}

func NewPlayerRepo(db *sql.DB) *PlayerRepo {
	return &PlayerRepo{db: db}
}

func (pr *PlayerRepo) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS players (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`
	_, err := pr.db.Exec(query)
	return err
}

func (pr *PlayerRepo) Create(player Player) (*Player, error) {
	res, err := pr.db.Exec("INSERT INTO players (name) VALUES (?)", player.Name)
	if err != nil {
		// If unique constraint on field
		// is added in the future
		// it should be handled here
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	player.ID = int(id)
	return &player, nil
}

func (pr *PlayerRepo) All() ([]Player, error) {
	rows, err := pr.db.Query("SELECT id, name FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.ID, &player.Name)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func (pr *PlayerRepo) Find(id int) (*Player, error) {
	row := pr.db.QueryRow("SELECT id, name FROM players WHERE id = ?", id)
	var player Player
	err := row.Scan(&player.ID, &player.Name)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *PlayerRepo) Update(player Player) (*Player, error) {
	_, err := pr.db.Exec("UPDATE players SET name = ? WHERE id = ?", player.Name, player.ID)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *PlayerRepo) Delete(id int) error {
	_, err := pr.db.Exec("DELETE FROM players WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
