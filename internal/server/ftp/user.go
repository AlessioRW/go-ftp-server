package ftp

import (
	"fmt"
	"strings"
)

func (c *Conn) user(args []string) {
	res := fmt.Sprintf(status230, strings.Join(args, " "))
	c.respond(res)
}
