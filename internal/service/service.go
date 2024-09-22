package service

import (
	"fmt"
	"net/url"
	"url_shortener/internal/helper"
	"url_shortener/internal/models"
	"url_shortener/internal/repository"
)

type UserState struct {
	Original string
	Title    string
	State    int
}

var userStates = make(map[int]*UserState)

func GetOriginalByShortened(shortened string) (string, error) {
	original, err := repository.GetOriginalByShortened(shortened)
	if err != nil {
		return "", err
	}

	return original, nil
}

func CreateLink(request models.CreateLinkRequest, user models.User) (*models.CreateLinkResponse, error) {
	_, exists := userStates[user.ChatID]
	if !exists {
		userStates[user.ChatID] = &UserState{State: 0}
	}

	userState := userStates[user.ChatID]

	switch userState.State {
	case 0:
		userState.Original = request.Message
		userState.State = 1

		return nil, nil
	case 1:
		userState.Title = request.Message
		defer delete(userStates, user.ChatID)
	}

	if userState.State == 0 {
		_, err := url.ParseRequestURI(request.Message)
		if err != nil {
			return nil, fmt.Errorf("ParseRequestURI: %w", err)
		}
	}

	shortened := helper.GenerateRandomString(6)

	link, err := repository.CreateLink(userState.Original, shortened, userState.Title, user)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateLink: %w", err)
	}

	prefix := "http://localhost:8090/"

	response := &models.CreateLinkResponse{
		ChatID:    link.ChatID,
		Original:  link.Original,
		Shortened: prefix + link.Shortened,
	}

	return response, nil
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
