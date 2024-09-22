package models

type Link struct {
	ChatID    int    `json:"chatID"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

type CreateLinkRequest struct {
	Message string `json:"message"`
	/*	Original string `json:"original"`
		Title    string `json:"title"`*/
}

type CreateLinkResponse struct {
	ChatID    int    `json:"chatID"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

type User struct {
	ChatID int `json:"chatID"`
}
