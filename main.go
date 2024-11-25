package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/islombektoshev/RocksQL/engine"
	"github.com/islombektoshev/RocksQL/server"
)

const DefaultDbPath = "/tmp/rocksql/db"

var (
	dbPath   string
	port     int
	host     string
	maxConn  int
	logLevel string
)

func ParseParams() {
	var cli = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cli.StringVar(&dbPath, "dbpath", "rocksql-data", "Please set you some dbpath to run")
	cli.StringVar(&host, "host", "localhost", "Please set sever host to run")
	cli.StringVar(&logLevel, "log-level", "info", "Please set log level, values are: debug, info, warn, error")
	cli.IntVar(&port, "port", 6666, "Pease set server port to run")
	cli.IntVar(&maxConn, "max-conn", 20, "Plase set max num of simultanious connections")
	cli.Parse(os.Args[1:])
}

func ParseLogLevel() (slog.Level, error) {
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	}
	return 0, fmt.Errorf("unexpected value")
}

func main() {
	ParseParams()

	var db, err = engine.OpenDB(dbPath)
	if err != nil {
		slog.Error("Cannot open db", "path", dbPath, "error", err)
		os.Exit(1)
	}
	defer db.Close()

	eng := engine.NewEngine(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var terminate = MakeTerminateChann()

	StartServer(ctx, eng, maxConn, &wg)

	select {
	case <-terminate:
		cancel()
	}
	wg.Wait()
}

func MakeTerminateChann() chan os.Signal {
	var terminate = make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	return terminate
}

func StartServer(ctx context.Context, eng engine.Engine, numOfConn int, wg *sync.WaitGroup) {
	svr := server.NewServer(&CmdRouter{Engine: eng}, fmt.Sprintf("%s:%d", host, port))
	err := svr.StartListening()
	if err != nil {
		slog.Error("Unable to listern port", "port", port, "error", err)
		os.Exit(1)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Accepting connections")
		svr.AcceptContinously(ctx, numOfConn)
		slog.Info("Server is stoped")
	}()
}
