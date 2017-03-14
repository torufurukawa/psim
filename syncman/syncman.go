package syncman

import "flag"

func parseFlags(args []string) *flags {
	f := flags{}
	fs := flag.NewFlagSet("syncman", flag.ExitOnError)
	fs.UintVar(&f.Port, "port", 6000, "port number")
	fs.StringVar(&f.ConfigPath, "config", "config.xml", "path to config file")
	fs.Parse(args[1:])
	return &f
}

type flags struct {
	Port       uint
	ConfigPath string
}
