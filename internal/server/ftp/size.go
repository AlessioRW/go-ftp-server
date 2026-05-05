package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) size(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	filename := args[0]
	fullPath := filepath.Join(c.rootDir, c.workDir, filename)
	f, err := os.Stat(fullPath)
	if err != nil {
		if err == os.ErrNotExist {
			c.respond(status550)
		} else {
			log.Print("ERROR failed to read file: ", err)
			c.respond(status451)
		}
		return
	}

	res := fmt.Sprintf(status213, f.Size())
	c.respond(res)
}
