package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"url_shortener/internal/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "url_shortener"
)

var db *sql.DB

func GetOriginalByShortened(shortened string) (string, error) {
	row := db.QueryRow(`SELECT original FROM links WHERE shortened = $1;`, shortened)

	var original string

	err := row.Scan(&original)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return original, nil
}

func GetLinks() ([]models.Link, error) {
	rows, err := db.Query(`SELECT id, original, shortened FROM links;`)
	if err != nil {
		fmt.Println(err)
		return []models.Link{}, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var link models.Link

		err = rows.Scan(&link.ID, &link.Original, &link.Shortened)
		if err != nil {
			fmt.Println(err)
			return []models.Link{}, err
		}

		links = append(links, link)
	}

	return links, nil
}

func GetLinkByID(linkID int) (models.Link, error) {
	row := db.QueryRow(`SELECT id, original, shortened FROM links WHERE id = $1;`, linkID)

	var link models.Link

	err := row.Scan(&link.ID, &link.Original, &link.Shortened)
	if err != nil {
		fmt.Println(err)
		return models.Link{}, err
	}

	return link, nil
}

func init() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	db = conn
}
