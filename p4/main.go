package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pei0804/go-db-crud-sample/p2/model"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	log.Println("start")
	migration := &migrate.FileMigrationSource{
		Dir: "./migration",
	}
	dbcon, err := sql.Open("sqlite3", "./dev.db?loc=auto")
	if err != nil {
	}
	n, err := migrate.Exec(dbcon, "sqlite3", migration, migrate.Up)
	if err != nil {
	}
	log.Printf("Applied %d migrations!\n", n)

	db := model.NewDB(dbcon)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	err = Commit(tx.Tx, func() error {
		_, err = tx.ArticleAdd("title", "body")
		if err != nil {
			return err
		}
		_, err = tx.TimelinesAdd("title")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("---Articles---")
	articles, _ := db.ArticlesAll()
	for k, v := range articles {
		log.Println(k, v)
	}

	log.Println("---Timelines---")
	timelines, _ := db.TimelinesAll()
	for k, v := range timelines {
		log.Println(k, v)
	}
}

func Commit(tx *sql.Tx, txFunc func() error) (err error) {
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
	err = txFunc()
	return
}
