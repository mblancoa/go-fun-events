package main

import (
	"github.com/astaxie/beego"
	"github.com/mblancoa/go-fun-events/api"
	repo "github.com/mblancoa/go-fun-events/pkg/adapters/mongodb-repository"
	"github.com/mblancoa/go-fun-events/pkg/core"
)

func main() {
	repo.SetupMongodbRepositoryConfiguration()
	core.SetupCoreConfiguration()
	api.SetupApiConfiguration()

	beego.Run()
}
