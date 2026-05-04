package ftp

import (
	"fmt"
)

func (c *Conn) pwd() {
	res := fmt.Sprintf(status257, c.workDir)
	c.respond(res)
}
