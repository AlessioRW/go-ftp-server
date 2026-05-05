package ftp

import (
	"log"
	"os"
	"path/filepath"
)

// does moving back (`cd ..`) need to be implemented explicitly
func (c *Conn) cwd(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	// form workdir, validate, and set
	workDir := filepath.Join(c.rootDir, args[0])
	_, err := os.Stat(workDir)
	if err != nil {
		log.Print("ERROR failed to read new path: ", err)
		c.respond(status550)
		return
	}
	c.workDir = args[0]
	c.respond(status250)
}
