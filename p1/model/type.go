//go:generate scaneo $GOFILE

package model

type Article struct {
	Title string `db:"title"`
	Body  string `db:"body"`
}

type Timeline struct {
	Title string `db:"title"`
}
