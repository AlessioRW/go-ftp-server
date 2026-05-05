package ftp

import "fmt"

const statText = `FTP server status:
	Connected to localhost
`

func (c *Conn) stat() {
	res := fmt.Sprintf(status211, statText)
	c.respond(res)
}
