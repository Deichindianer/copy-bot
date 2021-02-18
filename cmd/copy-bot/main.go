package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Deichindianer/copy-bot/internal/copybot"
)

func main() {
	cp := copybot.NewCopyBot()
	if os.Getenv("COPYBOT_LISTEN_ADDR") == "" {
		log.Fatal("Set COPYBOT_LISTEN_ADDR to a proper listen address!")
	}
	log.Fatal(http.ListenAndServe(os.Getenv("COPYBOT_LISTEN_ADDR"), cp))
}
