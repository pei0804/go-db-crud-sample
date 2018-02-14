package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pei0804/go-db-crud-sample/p1/model"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	log.Println("start")
	migration := &migrate.FileMigrationSource{
		Dir: "../migration",
	}
	db, err := sql.Open("sqlite3", "dev.db?loc=auto")
	if err != nil {
	}
	n, err := migrate.Exec(db, "sqlite3", migration, migrate.Up)
	if err != nil {
	}
	log.Printf("Applied %d migrations!\n", n)

	article := &model.Article{
		Title: "title",
		Body:  "body",
	}
	timeline := &model.Timeline{
		Title: "title",
	}
	err = Transaction(db, func(tx *sql.Tx) error {
		_, err := article.Add(tx)
		if err != nil {
			return err
		}
		_, err = timeline.Add(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("---Articles---")
	articles, _ := model.ArticlesAll(db)
	for k, v := range articles {
		log.Println(k, v)
	}

	log.Println("---Timelines---")
	timelines, _ := model.TimelinesAll(db)
	for k, v := range timelines {
		log.Println(k, v)
	}
}

func Transaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
