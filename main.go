package main

import (
	"log"

	api "github.com/crafting-dev/template-go-gin/src"
)

func main() {
	var ctx api.Context

	ctx.Mode = "release"
	ctx.Port = "8080"

	if ctx.Port == "" {
		log.Fatal("PORT must be set")
	}

	api.Init(ctx)
}
