package syncman

import "flag"

func parseFlags(args []string) *flags {
	f := flags{}
	flag.UintVar(&f.Port, "port", 6000, "port number")
	flag.StringVar(&f.ConfigPath, "config", "config.xml", "path to config file")
	return &f
}

type flags struct {
	Port       uint
	ConfigPath string
}
