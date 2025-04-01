package main

import (
	"embed"

	"github.com/9688101/HX/internal/app"
)

//go:embed web/build/*
var BuildFS embed.FS

func main() {
	app.Run(BuildFS)
}
