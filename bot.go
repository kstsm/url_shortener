package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"url_shortener/internal/handler"
	"url_shortener/internal/repository"
	"url_shortener/internal/service"
)

func StartTelegramBot() {
	bot, err := tgbotapi.NewBotAPI("6365489606:AAFS-sQzKVlFr6Bhv5y4DP2R7K0zg9G8zUY")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Инициализация слоев
	repo, _ := repository.GetOriginalByShortened(serv)

	// Шаг 3: Инициализация сервиса с репозиторием
	serv, _ := service.GetLinks(hand)

	// Шаг 4: Инициализация обработчика с сервисом
	hand := handler.GetOriginalByShortened(serv)

	// Запуск бота
	hand.Start(bot)
}

func main() {
	StartTelegramBot()
}
