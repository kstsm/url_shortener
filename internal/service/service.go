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

	return original, nil
}

func GetLinks() ([]*models.Link, error) {
	links, err := repository.GetLinks()
	if err != nil {
		return nil, err
	}

	prefix := "http://localhost:8090/"

	for _, link := range links {
		link.Shortened = prefix + link.Shortened
	}

	return links, nil
}

func GetLinkByID(linkID int) (*models.Link, error) {
	link, err := repository.GetLinkByID(linkID)
	if err != nil {
		return nil, err
	}

	prefix := "http://localhost:8090/"
	link.Shortened = prefix + link.Shortened

	return link, nil
}

func DeleteLink(linkID int) error {
	err := repository.DeleteLink(linkID)
	if err != nil {
		return err
	}

	return nil
}
