package main

import "github.com/mcuadros/gop.kg"

type ServerCommand struct {
	Addr     string `long:"addr" default:":8080" description:"http server addr"`
	CertFile string `long:"cert" description:"TLS certificate file path."`
	KeyFile  string `long:"key" description:"TLS key file path."`
	s        *gopkg.Server
}

func (c *ServerCommand) Execute(args []string) error {
	c.s = gopkg.NewServer()
	return c.s.RunTLS(c.Addr, c.CertFile, c.KeyFile)
}
