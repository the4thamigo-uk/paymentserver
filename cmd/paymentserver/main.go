package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/the4thamigo-uk/paymentserver/pkg/server"
	"os"
)

func main() {
	var cfgPath = pflag.StringP("config", "c", "", "Path or URI to config file in YAML,JSON or TOML format.")
	var address = pflag.StringP("address", "l", "", "Specify listen address. Overrides the value in the config file.")
	pflag.Parse()

	var err error
	cfg := &server.Config{}
	if *cfgPath != "" {
		cfg, err = server.LoadConfig(*cfgPath)
		if err != nil {
			fmt.Println(err)
			pflag.Usage()
			os.Exit(1)
			return
		}
	}
	if *address != "" {
		cfg.Address = *address
	}
	s := server.NewServer(cfg)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
