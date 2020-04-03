package main

import (
	"flag"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/config"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/service"
)

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	checkErr(service.Start(cfg))

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
