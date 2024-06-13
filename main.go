package main

import (
	"log"
	"net/http"

	"github.com/RiteshHarihar/token-generator/handler"
)

func main() {
	authCode := make(chan string)
	th := handler.New(authCode)

	go func() {
		http.HandleFunc("/callback", th.CallbackHandler)

		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	th.GenerateToken(authCode)
}
