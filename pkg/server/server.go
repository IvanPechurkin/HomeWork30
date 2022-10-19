package server

import (
	"30/cmd/utils"
	"30/pkg/api/routers"
	"30/pkg/client/mongodb"
	"30/pkg/storage/repositories"
	"github.com/go-chi/chi"
)

func Init(config utils.Configuration) *chi.Mux {
	client := mongodb.ConnectMongoDb(config.Database.Url)
	repo := repositories.NewMongoDbRepository(client, &config)
	r := routers.NewRouter(repo)

	return r
}
