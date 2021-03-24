package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World")
}
