package main

import (
	"context"
	"github.com/aibotsoft/latency-metric/collector"
	"github.com/aibotsoft/latency-metric/config"
	"github.com/aibotsoft/latency-metric/logger"
	"github.com/aibotsoft/latency-metric/signals"
	"github.com/aibotsoft/latency-metric/version"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	log, err := logger.NewLogger(cfg.Zap.Level, cfg.Zap.Encoding, cfg.Zap.Caller)
	if err != nil {
		panic(err)
	}
	log.Info("start_service", zap.String("version", version.Version), zap.String("build_date", version.BuildDate), zap.Any("config", cfg))
	ctx, cancel := context.WithCancel(context.Background())
	c := collector.NewCollector(cfg, log, ctx)

	errCh := make(chan error)
	go func() {
		errCh <- c.Run()
	}()
	defer func() {
		log.Info("closing_services...")
		cancel()
		_ = log.Sync()
	}()
	stopCh := signals.SetupSignalHandler()
	select {
	case err := <-errCh:
		log.Error("stop_service_by_error", zap.Error(err))
	case sig := <-stopCh:
		log.Info("stop_service_by_os_signal", zap.String("signal", sig.String()))
	}

}
