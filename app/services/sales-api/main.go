package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/grayjunzi/gola/app/foundation/logger"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New("SALES-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {

	// --------------------------------------------------------------
	// GOMAXPROCS: 操作系统线程数，调整并发的运行性能。
	// --------------------------------------------------------------
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// --------------------------------------------------------------
	// shutdown: 友好关闭程序
	// --------------------------------------------------------------
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown

	log.Infow("shutdown", "status", "shutdown started", "signal", sig)
	defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

	return nil
}
