package db

import (
	"database/sql"
)

type SingleBook struct {
	Name   string
	Author string
}

type SingleHighlightNote struct {
	HightLight string
	Note       sql.NullString
	Style int
}

type SingleBookInList struct {
	Id     string
	Title  string
	Author string
	Number int
}
