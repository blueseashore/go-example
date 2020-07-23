package main

import "net/http"

func main() {
	http.HandleFunc("/say", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world"))
	})
	_ = http.ListenAndServe(":9090", nil)
}
