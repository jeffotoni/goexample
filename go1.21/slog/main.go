package main

import (
	"log/slog"
	"os"
)

func main() {

	slog.Info("hello", "count", 3)

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("hello", "count", 3)
	logger.Error("hello", "count", 1)

	logger2 := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger2.Info("hello", "count", 3)
	logger2.Debug("hello debub", "count", 12)
	slog.SetDefault(logger2)

}
