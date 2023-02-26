package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created int
	Expires int
}
