package main

import (
	repo "github.com/mblancoa/go-fun-events/adapters/mongodb-repository"
	"github.com/mblancoa/go-fun-events/core"
)

func main() {
	repo.SetupMongodbRepositoryConfiguration()
	core.SetupCoreConfiguration()
	//api.SetupApiConfiguration()
}
