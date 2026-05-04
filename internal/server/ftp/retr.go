package ftp

import (
	"io"
	"log/slog"
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
		slog.Error(
			"failed to open file",
			"path", path,
			"error", err,
		)
		c.respond(status550)
	}
	c.respond(status150)

	// open new connetion to client
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	// write requested file to client
	// wondering if this data is sent progessive
	_, err = io.Copy(dataConn, file)
	if err != nil {
		slog.Error("failed to copy file to writer")
		// set as 451, code example had 425
		c.respond(status451)
		return
	}

	io.WriteString(dataConn, c.EOL())
	c.respond(status226)
}
