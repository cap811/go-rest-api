package models

type Book struct {
	ID     uint
	Title  string
	Author *string
	Year   uint8
}
