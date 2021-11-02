package main

import api "github.com/crafting-dev/template-go-gin/pkg/api"

func main() {
	var ctx api.Context

	ctx.Mode = "release"
	ctx.Port = "3000"

	api.Init(ctx)
}
