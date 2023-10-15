package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"os"

	"application/config"

	"log/slog"

	slogmulti "github.com/samber/slog-multi"
	"go.uber.org/zap"
)

func main() {

	// ...

	cfg := &config.ViperConfig{}
	conf, err := config.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	// load config
	if err := conf.Load(cfg); err != nil {
		panic(err)
	}

	// logger := initZapLogger(cfg)
	// logger.Debug("config", zap.Any("cfg", cfg))
	logger := initSlogLogger(cfg)
	logger.Info("config", "cfg", cfg)
	engine, err := wireApp(cfg, logger)
	if err != nil {
		logger.Error("failed to init app", zap.Error(err))
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			logger.Error("failed to run app", zap.Error(err))
		}

	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal := <-quit

	logger.Info("app stopping...")
	ctx, cancell := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancell()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown app", err)
		panic(err)
	}

	logger.Info("app stopped", "signal", signal)

}

func initSlogLogger(conf *config.ViperConfig) *slog.Logger {

	// create list of slog handlers
	slogHandlerOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	switch conf.Observability.Logging.Level {
	case "debug":
		slogHandlerOptions.Level = slog.LevelDebug
	case "info":
		slogHandlerOptions.Level = slog.LevelInfo
	case "warn":
		slogHandlerOptions.Level = slog.LevelWarn
	case "error":
		slogHandlerOptions.Level = slog.LevelError
	default:
		slogHandlerOptions.Level = slog.LevelInfo
	}

	slogHandlers := []slog.Handler{}

	// add stdout
	slogHandlers = append(slogHandlers, slog.NewJSONHandler(os.Stdout, slogHandlerOptions))

	// add udp if logstash enabled
	if conf.Observability.Logging.Logstash.Enabled {
		// options := slogsyslog.Option{}
		fmt.Println("logstash enabled")
		// address := conf.Observability.Logging.Logstash.Address
		con, err := net.Dial("udp", "127.0.0.1:9000")
		if err != nil {
			panic(err)
		}
		slogHandlers = append(slogHandlers, slog.NewJSONHandler(con, slogHandlerOptions))
	}

	// init logger
	// con, err := net.Dial("udp", "127.0.0.1:9000")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(con, "hello world")

	logger := slog.New(slogmulti.Fanout(slogHandlers...))
	slog.SetDefault(logger)

	return logger

}
