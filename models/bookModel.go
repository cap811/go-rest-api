package models

import "time"

type CreateBookInput struct {
	Title  string
	Author string
	Year   string
}

type UpdateBookInput struct {
	ID     uint
	Title  string
	Author string
	Year   string
}

type Book struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Author    string    `json:"author,omitempty"`
	Year      string    `json:"year,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
