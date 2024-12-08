package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/library/memory" // Или импортируйте нужный storage (например, sqlite)

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type App struct {
	config *Config
	router *chi.Mux
	http   *http.Server
}

func New(ctx context.Context, config *Config) (*App, error) {
	r := chi.NewRouter()
	return &App{
		config: config,
		router: r,
		http: &http.Server{
			Addr:              config.Host + ":" + config.Port,
			Handler:           r,
			ReadTimeout:       0,
			ReadHeaderTimeout: 0,
			WriteTimeout:      0,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
		},
	}, nil
}

// Initialize db, service and setup Handler with HTTP requests
func (a *App) Setup(ctx context.Context) error {
	// Initialize db
	store := memory.NewMemoryBookStore()

	// Initialize service
	service := library.NewBookService(store)

	// Create User
	user := library.NewUserServiceClient(a.config.UserHost + ":" + a.config.UserInternalPort)

	// Create Handler
	handler := library.NewHandler(a.router, service, user)
	handler.Register()

	return nil
}

// Run HTTP-server
func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	errs, ctx := errgroup.WithContext(ctx)

	log.Printf("starting web server on port %s", a.config.Port)

	// Run server
	errs.Go(func() error {
		if err := a.http.ListenAndServe(); err != nil {
			return fmt.Errorf("listen and serve error: %w", err)
		}
		return nil
	})

	<-ctx.Done()

	// Graceful shutdown (we got the interrupt signal)
	stop()
	log.Println("shutting down gracefully")

	// Perform application shutdown with a maximum timeout of 5 seconds
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.http.Shutdown(timeoutCtx); err != nil {
		log.Println("error during shutdown:", err)
	}

	return nil
}
