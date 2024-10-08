package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"url_shortener/internal/cerrors"
	"url_shortener/internal/models"
	"url_shortener/internal/repository/queries"
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

func CreateLink(original, shortened, title string, user models.User) (*models.Link, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(queries.CreateLink, original, shortened)

	var link models.Link
	err = row.Scan(&link.ChatID, &link.Original, &link.Shortened)
	if err != nil {
		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	_, err = tx.Exec(queries.CreateTgLinks, link.ChatID, user.ChatID, title) //toDo: Upsert
	if err != nil {
		return nil, fmt.Errorf("tx.Exec: CreateTgLinks: %w", err)
	}

	return &link, tx.Commit()
}

func GetLinks() ([]*models.Link, error) {
	rows, err := db.Query(`SELECT id, original, shortened FROM links;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*models.Link
	for rows.Next() {
		var link models.Link

		err = rows.Scan(&link.ChatID, &link.Original, &link.Shortened)
		if err != nil {
			return nil, err
		}

		links = append(links, &link)
	}

	return links, nil
}

func GetLinkByID(linkID int) (*models.Link, error) {
	row := db.QueryRow(`SELECT id, original, shortened FROM links WHERE id = $1;`, linkID)

	var link models.Link

	err := row.Scan(&link.ChatID, &link.Original, &link.Shortened)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func DeleteLink(linkID int) error {
	tag, err := db.Exec(`DELETE FROM links WHERE id = $1;`, linkID)
	if err != nil {
		return err
	}

	rows, err := tag.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return cerrors.ErrNotFound
	}

	return nil
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

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	db = conn
}
