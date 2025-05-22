package main

import (
	beego "github.com/beego/beego/v2/server/web"
	repo "github.com/mblancoa/go-fun-events/adapters/mongodb-repository"
	"github.com/mblancoa/go-fun-events/api"
	"github.com/mblancoa/go-fun-events/core"
)

func main() {
	repo.SetupMongodbRepositoryConfiguration()
	core.SetupCoreConfiguration()
	api.SetupApiConfiguration()

	beego.Run()
}
