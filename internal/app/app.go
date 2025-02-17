package app

import (
	"computerextra/elaerning-go/db"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"io/fs"

	"github.com/rs/cors"
)

type App struct {
	config Config
	files  fs.FS
	db     *db.PrismaClient
}

func New(config Config, files fs.FS) (*App, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &App{
		config: config,
		files:  files,
		db:     client,
	}, nil
}

func (a *App) Start(ctx context.Context) error {

	router, err := a.loadRoutes()
	if err != nil {
		return fmt.Errorf("failed when loading routes: %w", err)
	}

	// TODO: Middlewares

	port := getPort(3000)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(router)

	// Start Server
	srv := &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%v", port),
		// Good Pratice: enfoce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	errCh := make(chan error, 1)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to listen and serve: %w", err)
		}

		close(errCh)
	}()

	select {
	// Wait until we receive SIGINT (ctrl+c on cli)
	case <-ctx.Done():
		break
	case err := <-errCh:
		return err
	}

	sCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(sCtx)

	return nil
}

func getPort(defaultPort int) int {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return defaultPort
	}

	return port
}
