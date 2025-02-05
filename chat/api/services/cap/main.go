package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/4hmedhabib/chat_go/chat/foundation/logger"
)

type TraceIDFn func(ctx context.Context) string

type Logger struct {
	handler   slog.Handler
	traceIDFn TraceIDFn
}

func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return "" // TODO: NEED TRACE IDs
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIDFn)

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	log.Info(ctx, "startup", "status", "started")
	defer log.Info(ctx, "startup", "status", "shutdown")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	return nil
}
