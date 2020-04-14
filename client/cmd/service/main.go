package main

import (
	"flag"
	"log"

	"github.com/fgituser/management-client-psychologist.services/client/internal/config"
	"github.com/fgituser/management-client-psychologist.services/client/internal/service"
)

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	log.Println("----------- ", cfg.Server.Port)
	checkErr(service.Start(cfg))

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
