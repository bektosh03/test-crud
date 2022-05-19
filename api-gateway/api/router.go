package api

import (
	"fmt"

	_ "github.com/bektosh03/test-crud/api-gateway/api/docs"
	"github.com/bektosh03/test-crud/api-gateway/config"
	"github.com/bektosh03/test-crud/api-gateway/ports"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title        HilolTest API
// @version      1.0
// @description  This is a first version of HilolTest APIs

// @contact.name   Bektosh Madaminov
// @contact.email  madaminovsbektosh@gmail.com
func InitRoutes(cfg config.Config, router chi.Router, s ports.HTTPServer) {
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/api/swagger/doc.json", cfg.ListenAddress())), //The url pointing to API definition
	))

	router.Post("/posts/download", s.DownloadPosts)
	router.Get("/posts/download/status", s.GetDownloadStatus)
	router.Get("/posts/{post-id}", s.GetPost)
}
