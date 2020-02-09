package main

import (
	setup "github.com/InsideCI/nego/src"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("app.env"); err != nil {
		panic("error loading .env file")
	}
	setup.Init()
}
