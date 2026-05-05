package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) nlst(args []string) {
	var target string

	// form list target
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := os.ReadDir(target)
	if err != nil {
		log.Print("ERROR failed to read directory '"+target+"': ", err)
		c.respond(status550)
		return
	}

	// open connection to client
	// used to write files back as a response
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	c.respond(status125)

	// loop over files and write to client
	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), c.EOL())
		if err != nil {
			log.Print("ERROR failed to write file name: ", err)
			c.respond(status426)
			return
		}
	}

	_, err = fmt.Fprint(dataConn, c.EOL())
	if err != nil {
		log.Print("ERROR failed to write: ", err)
		c.respond(status426)
		return
	}

	c.respond(status226)
}
