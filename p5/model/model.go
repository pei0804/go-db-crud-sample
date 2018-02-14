package model

import (
	"database/sql"
	"log"
)

type DB struct {
	*sql.DB
}

type Tx struct {
	*sql.Tx
}

func NewDB(db *sql.DB) *DB {
	return &DB{db}
}

func (d *DB) Transaction(txFunc func(tx *Tx) error) (err error) {
	tx, err := d.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return
}

func (d *DB) Begin() (*Tx, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

func (d *DB) TimelinesAll() ([]Timeline, error) {
	rows, err := d.Query(`SELECT title FROM timelines`)
	if err != nil {
		return nil, err
	}
	return ScanTimelines(rows)
}

func (d *Tx) TimelinesAdd(title string) (sql.Result, error) {
	stmt, err := d.Prepare(`
	INSERT INTO timelines (title)
	VALUES(?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(title)
}

func (d *DB) ArticlesAll() ([]Article, error) {
	rows, err := d.Query(`SELECT title,body FROM articles`)
	if err != nil {
		return nil, err
	}
	return ScanArticles(rows)
}

func (d *Tx) ArticleAdd(title string, body string) (sql.Result, error) {
	stmt, err := d.Prepare(`
	INSERT INTO articles (title, body)
	VALUES(?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(title, body)
}
