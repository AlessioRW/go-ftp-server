package ftp

import "fmt"

const features = `Features:
	CWD
	PWD
	LIST
	NLST
	POR
	EPRT
	USER
	RETR
	SIZE
	TYPE
	FEAT
	SYST
	STAT
	PASV
	EPSV
	QUIT
`

func (c *Conn) feat() {
	res := fmt.Sprintf(status211, features)
	c.respond(res)
}
