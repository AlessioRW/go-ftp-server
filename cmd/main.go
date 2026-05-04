package main

import (
	"go-ftp-server/internal/handler"
	"go-ftp-server/internal/logging"
)

func main() {
	logging.InitSlog()
	handler.Run()
}
