package main

import (
	"fmt"
	"main/internal"
	"os"

	logger "github.com/can-zanat/gologger"
	_ "github.com/go-sql-driver/mysql"
)

const serverPort = ":83"

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//nolint:unparam // run function is not using any parameters
func run() error {
	loggerInfoLevel := logger.NewWithLogLevel("info")
	defer func() {
		err := loggerInfoLevel.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	service := internal.NewService()
	handler := internal.NewHandler(service)

	New(serverPort, handler, loggerInfoLevel).Run()

	return nil
}
