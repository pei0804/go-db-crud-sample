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
		Dir: "../migration",
	}
	dbcon, err := sql.Open("sqlite3", "dev.db?loc=auto")
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

	// bad pattern
	// becouse this source doesn't execute rollback in panic
	_, err = tx.ArticleAdd("title", "body")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = tx.TimelinesAdd("title")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

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
