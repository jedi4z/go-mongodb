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
		panic(err)
	}

	engine := rest.NewEngine(ctn)
	if err := engine.Run(); err != nil {
		log.Fatalf("failed to init engine: %v", err)
		panic(err)
	}
}
