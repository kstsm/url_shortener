package service

import (
	"fmt"
	"url_shortener/internal/models"
	"url_shortener/internal/repository"
)

func GetOriginalByShortened(shortened string) (string, error) {
	original, err := repository.GetOriginalByShortened(shortened)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// ToDo: добавить префикс

	return original, nil
}

func GetLinks() ([]models.Link, error) {
	links, err := repository.GetLinks()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// ToDo: добавить префикс

	return links, nil
}

func GetLinkByID(linkID int) (models.Link, error) {
	link, err := repository.GetLinkByID(linkID)
	if err != nil {
		fmt.Println(err)
		return models.Link{}, err
	}

	// ToDo: добавить префикс
	prefix := "http://localhost:8090/"
	link.Shortened = prefix + link.Shortened

	return link, nil
}
