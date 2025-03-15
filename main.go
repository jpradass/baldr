package main

import (
	"fmt"
	"os"

	"github.com/jpradass/baldr/log"
)

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Printf("There was an error instantiating the logger. Error: %w", err)
		os.Exit(1)
	}

	logger.Info("loading baldr gateway...")
}
