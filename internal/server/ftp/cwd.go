package ftp

import (
	"log/slog"
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
	workDir := filepath.Join(c.workDir, args[0])
	_, err := os.Stat(filepath.Join(c.rootDir, workDir))
	if err != nil {
		slog.Error("failed to read new path", "error", err)
		c.respond(status550)
		return
	}
	c.workDir = workDir
	c.respond(status200)
}
