package main

import (
	"go-ftp-server/internal/config"
	"go-ftp-server/internal/handler"
	"os"
)

func main() {
	err := config.Parse(os.Args)
	if err != nil {
		return
	}

	handler.Run()
}
