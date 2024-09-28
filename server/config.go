package server

import "flag"

type Config struct {
	Host string
	Port string
}

func (c *Config) bindFlags(fs flag.FlagSet) {
	fs.StringVar(&c.Host, "host", "127.0.0.1", "hostname defaults to '127.0.0.1'")
	fs.StringVar(&c.Port, "port", "8080", "port defaults to '8080'")
}
