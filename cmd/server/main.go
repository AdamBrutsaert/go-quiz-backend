package main

import "github.com/AdamBrutsaert/go-quiz-backend/server"

func main() {
	err := server.New().Run()
	if err != nil {
		panic(err)
	}
}
