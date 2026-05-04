package main

import (
	"music-server/internal/handler"
	"music-server/internal/logging"
)

func main() {
	logging.InitSlog()
	handler.Run()
}
