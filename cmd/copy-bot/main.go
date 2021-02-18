package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Deichindianer/copy-bot/internal/copybot"
)

func main() {
	cp := copybot.New()
	listenAddr, ok := os.LookupEnv("COPYBOT_LISTEN_ADDR")
	if !ok {
		log.Fatal("Set COPYBOT_LISTEN_ADDR to a proper listen address!")
	}
	log.Fatal(http.ListenAndServe(listenAddr, cp))
}
