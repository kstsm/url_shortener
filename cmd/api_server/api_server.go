package api_server

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"url_shortener/internal/handler"
)

var Run = &cobra.Command{
	Use: "api",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := handler.InitRoutes()
		fmt.Println("Starting server...")
		http.ListenAndServe(":8090", r)

		return nil
	},
}
