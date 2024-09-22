package telegram_server

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"url_shortener/internal/models"
	"url_shortener/internal/service"

	"github.com/go-telegram/bot"
	tgModels "github.com/go-telegram/bot/models"
)

var Run = &cobra.Command{
	Use: "telegram",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		fmt.Println("Запуск телеги")

		opts := []bot.Option{
			bot.WithDefaultHandler(handler),
		}

		b, err := bot.New("6365489606:AAFS-sQzKVlFr6Bhv5y4DP2R7K0zg9G8zUY", opts...)
		if err != nil {
			panic(err)
		}
		b.Start(ctx)

		return nil
	},
}

func handler(ctx context.Context, b *bot.Bot, update *tgModels.Update) {
	request := models.CreateLinkRequest{
		Message: update.Message.Text,
	}

	user := models.User{
		ChatID: int(update.Message.Chat.ID),
	}

	link, err := service.CreateLink(request, user)
	if err != nil {
		fmt.Println(err.Error())
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Ашибка",
		})
		return
	}

	if link != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Ваша ссылка создана \n \n %s: %s", request.Message, link.Shortened),
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: user.ChatID,
			Text:   "Теперь отправтье название ссылки",
		})
	}

}
