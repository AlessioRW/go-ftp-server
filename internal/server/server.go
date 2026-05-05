package server

import (
	"fmt"
	"go-ftp-server/internal/server/ftp"
	"log"
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
		log.Print("ERROR failed to create listener: ", err)
		os.Exit(1)
	}

	log.Print("server listening on port ", port)
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
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Print("ERROR failed to get absolute path: ", err)
		return
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
