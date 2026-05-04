package server

import (
	"fmt"
	"log/slog"
	"music-server/internal/server/ftp"
	"net"
	"os"
	"path/filepath"
)

const port = 21
const rootDir = "store"

func InitServer() {
	srv := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", srv)
	if err != nil {
		slog.Error("failed to create listener", "port", port, "error", err)
		os.Exit(1)
	}

	slog.Info("server listening", "port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "error", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		slog.Error("failed to get absolute path", "error", err)
		return
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
