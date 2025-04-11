package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/popflix-live/api/src/application/router"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func loadRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.RegisterRoutes(r)

	return r
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}
	serverErrors := make(chan error, 1)

	go func() {
		fmt.Println("Server started on :8000")
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		return server.Shutdown(ctx)
	}
}
