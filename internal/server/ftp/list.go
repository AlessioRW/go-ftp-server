package ftp

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	var target string

	// form list target
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := os.ReadDir(target)
	if err != nil {
		slog.Error(
			"failed to read directory",
			"target", target,
			"error", err,
		)
		c.respond(status550)
		return
	}

	// send 150 before opening data connection
	c.respond(status150)

	// open connection to client
	// used to write files back as a response
	// new connection is used so controller connection can stay open (i think)
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	// loop over files and write to client
	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), c.EOL())
		if err != nil {
			slog.Error("failed to write file name", "error", err)
			c.respond(status426)
			return
			// should command be terminated here??
			// code 426 specifed transfer aborted but example code doesn't
		}
	}

	// why send this last bit?
	// i should read the FTP spec properly
	_, err = fmt.Fprint(dataConn, c.EOL())
	if err != nil {
		slog.Error("failed to write", "error", err)
		c.respond(status426)
		return
	}

	c.respond(status226)
}
