package films

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Get("/{id}", ListSingleFilm)

	router.Get("/search/{name}", SearchFilm)

	return router
}
