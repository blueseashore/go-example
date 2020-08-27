package server

import (
	"github.com/micro/go-micro/v2/web"
	"log"
	"net/http"
)
func main() {
	service := web.NewService(
		web.Name("uckendo.com"),
	)
	service.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(request.Host))
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
