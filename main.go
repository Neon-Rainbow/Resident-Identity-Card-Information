package main

import (
	"id-information/repository"
	"id-information/router"
)

func main() {
	repository.InitRepository()

	r := router.InitRouter()
	r.Run()
}
