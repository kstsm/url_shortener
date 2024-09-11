package models

type Link struct {
	ID        int    `json:"id"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}
