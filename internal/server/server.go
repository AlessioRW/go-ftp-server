package server

import (
	"fmt"
	"go-ftp-server/internal/config"
	"go-ftp-server/internal/server/ftp"
	"log"
	"net"
	"os"
	"path/filepath"
)

func InitServer() {
	srv := fmt.Sprintf(":%d", config.Config.HostPort)
	listener, err := net.Listen("tcp", srv)
	if err != nil {
		log.Print("ERROR failed to create listener: ", err)
		os.Exit(1)
	}

	log.Print("server listening on port ", config.Config.HostPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("ERROR failed to accept connection: ", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(config.Config.StorageRoot)
	if err != nil {
		log.Print("ERROR failed to get absolute path: ", err)
		return
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
