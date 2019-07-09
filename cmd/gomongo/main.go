package main

import (
	"github.com/jedi4z/go-mongodb/app/interface/rest"
	"github.com/jedi4z/go-mongodb/app/registry"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	engine := rest.NewEngine(ctn)

	engine.Run()
}
