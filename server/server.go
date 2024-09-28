package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"
)

func NewServer(
	config *Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
	)
	var handler http.Handler = mux
	handler = Logging(handler)
	handler = PanicRecovery(handler)
	return handler
}

func Run(
	ctx context.Context,
	args []string,
) error {
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)

	config := &Config{}
	config.bindFlags(*flagSet)

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	srv := NewServer(
		config,
	)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}
	go func() {
		slog.Info(fmt.Sprintf("listening on %s\n", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(fmt.Sprintf("error listening and serving: %s\n", err))
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error(fmt.Sprintf("error shutting down http server: %s\n", err))
		}
	}()
	wg.Wait()
	return nil
}
