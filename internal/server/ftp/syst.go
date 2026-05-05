package ftp

import "fmt"

func (c *Conn) syst() {
	// TODO: make dynamic
	res := fmt.Sprintf(status215, "MACOS")
	c.respond(res)
}
