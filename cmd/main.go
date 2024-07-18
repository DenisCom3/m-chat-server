package main

import (
	"fmt"
	"log"

	"github.com/DenisCom3/m-chat-server/internal/config"
)

func main() {
	err := config.MustLoad()
	if err != nil {
		log.Fatalf("failed to init config. %v", err)
	}
	fmt.Println(config.GetPostgres().Dsn())
}
