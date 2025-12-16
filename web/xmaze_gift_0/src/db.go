package xmazegift

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE elves (name TEXT PRIMARY KEY, birth_date TEXT, death_date TEXT, move TEXT)")
	if err != nil {
		panic(err)
	}

	for _, elf := range elves {
		err = insertElf(db, &elf)
		if err != nil {
			panic(err)
		}
	}

	slog.Info("DB is ready!")
	return db, nil
}

func insertElf(db *sql.DB, elf *Elf) error {
	_, err := db.Exec("INSERT INTO elves (name, birth_date, death_date, move) VALUES (?, ?, ?, ?)", elf.Name, elf.BirthDate, elf.DeathDate, "")
	return err
}

func killElf(db *sql.DB, name string) error {
	query := `
	UPDATE elves
	SET death_date = (?)
	WHERE name = (?);`
	_, err := db.Exec(query, time.Now().UTC().String(), name)
	return err
}

func getElfByName(db *sql.DB, name string) (Elf, error) {
	query := `SELECT * FROM elves WHERE name = ?`
	row := db.QueryRow(query, name)

	var elf Elf
	err := row.Scan(&elf.Name, &elf.BirthDate, &elf.DeathDate, &elf.Moves)
	if err != nil {
		slog.Error(err.Error())
		return Elf{}, err
	}

	return elf, nil
}

func getAncestors(db *sql.DB, name string) ([]Elf, error) {
	query := `SELECT *
		FROM elves
		WHERE death_date IS NOT NULL AND death_date <> ''
		AND birth_date < (SELECT birth_date FROM elves WHERE name = '%s')
		AND birth_date >= '2025-11-30'
		ORDER BY birth_date DESC
		LIMIT 1;`
	rows, err := db.Query(fmt.Sprintf(query, name))
	if err != nil {
		slog.Error(err.Error())
		return []Elf{}, err
	}
	defer rows.Close()

	var elves []Elf
	for rows.Next() {
		var dbElf Elf
		if err := rows.Scan(&dbElf.Name, &dbElf.BirthDate, &dbElf.DeathDate, &dbElf.Moves); err != nil {
			slog.Error(err.Error())
			return []Elf{}, err
		}
		elves = append(elves, dbElf)
	}

	if err = rows.Err(); err != nil {
		slog.Error(err.Error())
		return []Elf{}, err
	}
	return elves, err
}

func moveElf(db *sql.DB, elfName string, move string) (string, error) {
	_, err := db.Exec("UPDATE elves SET move = CONCAT(move, ?) WHERE name = ?;", move, elfName)
	if err != nil {
		return "", fmt.Errorf("you broke my memory")
	}

	var moves string
	err = db.QueryRow(
		"SELECT move FROM elves WHERE name = ?;",
		elfName,
	).Scan(&moves)

	if err != nil {
		return "", fmt.Errorf("you never moved")
	}
	return moves, err
}
