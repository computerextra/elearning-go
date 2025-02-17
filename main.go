package main

import (
	"computerextra/elaerning-go/internal/app"

	"context"
	"embed"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

var files embed.FS

func main() {
	godotenv.Load()

	app, err := app.New(app.Config{}, files)
	if err != nil {
		panic(err)

	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
