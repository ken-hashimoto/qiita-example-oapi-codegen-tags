package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ken-hashimoto/oapi-codegen-swagger/handler"
	"github.com/ken-hashimoto/oapi-codegen-swagger/middleware"
	"github.com/ken-hashimoto/oapi-codegen-swagger/router"
)

func main() {
	r := chi.NewRouter()
	middleware := middleware.NewMiddlewares()
	var apiHandler = handler.NewApiHandler(
		middleware,
	)

	// yamlの情報を構造化された形で復元
	swagger, err := router.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger spec: %v\n", err)
	}
	m := apiHandler.Middleware
	r.Use(m.ForbidMySecret(swagger))
	router.HandlerFromMux(apiHandler, r)
	http.ListenAndServe(":8080", r)
}
