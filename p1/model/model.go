package model

import "database/sql"

func TimelinesAll(db *sql.DB) ([]Timeline, error) {
	rows, err := db.Query(`SELECT title FROM timelines`)
	if err != nil {
		return nil, err
	}
	return ScanTimelines(rows)
}

func (t *Timeline) Add(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	INSERT INTO atimelines (title)
	VALUES(?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Title)
}

func ArticlesAll(db *sql.DB) ([]Article, error) {
	rows, err := db.Query(`SELECT title,body FROM articles`)
	if err != nil {
		return nil, err
	}
	return ScanArticles(rows)
}

func (a *Article) Add(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	INSERT INTO articles (title, body)
	VALUES(?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body)
}
