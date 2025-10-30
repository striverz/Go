package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/striverz/100xGo/controllers"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/movies", controllers.GetAllMovies)
	router.Post("/movie", controllers.CreateMovie)
	router.Delete("/movie/{id}", controllers.DeleMovie)

	return router
}
