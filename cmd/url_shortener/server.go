package url_shortener

import (
	"fmt"
	"net/http"
	"url_shortener/internal/handler"
)

func Run() {
	r := handler.InitRoutes()
	fmt.Println("Starting server...")
	http.ListenAndServe(":8090", r)
}
