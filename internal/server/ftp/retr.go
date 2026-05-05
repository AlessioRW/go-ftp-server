package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Print("ERROR failed to open file '"+path+"': ", err)
		c.respond(status550)
		return
	}
	defer file.Close()

	// open new connetion to client
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	c.respond(status150)

	// write requested file to client
	_, err = io.Copy(dataConn, file)
	if err != nil {
		log.Print("ERROR failed to copy file to writer: ", err)
		c.respond(status451)
		return
	}

	c.respond(status226)
}
