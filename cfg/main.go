package main

import (
	"flag"

	cfg "github.com/rustyeddy/goconfig"
)

var (
	config	 cfg.Configuration
)

func init() {
	flag.StringVar(&config.Addr, "addr", "", "Bind our HTTP to this address")
	flag.StringVar(&config.Broker, "broker", "tcp://localhost:1883", "Address of MQTT broker")
	flag.BoolVar(&config.Verbose, "verbose", false, "Print logs to stdout")

	flag.StringVar(&config.Filename, "fname", "config.json", "Write to file")

}

func main() {
	if len(flag.Args()) > 1 {
		config.Read(flag.Arg(1))
	} else {
		flag.Parse()
	}

	if config.Filename != "" {
		config.Write(config.Filename)
	} else {
		config.Dump()
	}
}

