package server

import (
	"net/http"

	"github.com/bektosh03/test-crud/common/environment"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.Info("Starting HTTP server")
	logrus.Fatal(http.ListenAndServe(addr, rootRouter))
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
}

func addCorsMiddleware(router *chi.Mux) {
	debug := false
	if environment.Current() == environment.Development {
		debug = true
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            debug,
	})

	router.Use(corsMiddleware.Handler)
}
