package main

import (
	"log"
	"net/http"

	"github.com/Deichindianer/copy-bot/internal/copybot"
)

func main() {
	cp := copybot.NewCopyBot()
	log.Fatal(http.ListenAndServe(":8080", cp))
}
