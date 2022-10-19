package routers

import (
	"30/pkg/api/handlers"
	"30/pkg/storage/interfaces"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(storage interfaces.Repository) *chi.Mux {
	//Инициализация маршрутизатора
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Инициализация хендлов запросов
	r.Get("/users", handlers.GetAll(storage))
	r.Post("/create", handlers.Add(storage))
	r.Post("/makeFriends", handlers.Link(storage))
	r.Delete("/delete", handlers.Delete(storage))
	r.Put("/{user_id}", handlers.Update(storage))
	r.Get("/friends/{user_id}", handlers.GetFriends(storage))

	return r
}
