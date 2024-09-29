package main

import (
	"elerphore/cybersport-parser/internal/scheduler"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	scheduler.StatTask()
}
