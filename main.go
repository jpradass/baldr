package main

import (
	"fmt"
	"os"

	"github.com/jpradass/baldr/log"
	"github.com/jpradass/baldr/modules/server"
)

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Printf("There was an error instantiating the logger. Error: %s", err.Error())
		os.Exit(1)
	}

	logger.Info("loading baldr gateway...")
	server := server.New("", 9000, false)
	server.Start()
}
