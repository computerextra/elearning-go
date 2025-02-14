package main

import (
	"computerextra/elaerning-go/internal/app"

	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

var files embed.FS

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app, err := app.New(logger, app.Config{}, files)
	if err != nil {
		logger.Error("failed to create app", slog.Any("error", err))
	}

	if err := app.Start(ctx); err != nil {
		logger.Error("failed to start app", slog.Any("error", err))
	}

	// router := mux.NewRouter()

	// // Static Assets
	// var dir string
	// flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()

	// // Get DB Client
	// client := db.NewClient()
	// if err := client.Prisma.Connect(); err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// routes.GetRoutes(router, client)
	// routes.GetApiRoutes(router, client)

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// })
	// handler := c.Handler(router)

	// env := env.GetEnv()

	// srv := &http.Server{
	// 	Handler:      handler,
	// 	Addr:         fmt.Sprintf(":%v", env.PORT),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())
}
