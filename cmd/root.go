package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"url_shortener/cmd/telegram_server"
)

var cmd = &cobra.Command{}

func Execute() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cmd.AddCommand(telegram_server.Run)
	//cmd.AddCommand(api_server.Run)
}
