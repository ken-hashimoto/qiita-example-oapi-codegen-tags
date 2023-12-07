package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ken-hashimoto/oapi-codegen-swagger/middleware"
	"github.com/ken-hashimoto/oapi-codegen-swagger/router"
)

func NewApiHandler(
	middleware middleware.Middleware,
) *ApiHandler {
	return &ApiHandler{
		Middleware: middleware,
	}
}

type ApiHandler struct {
	Middleware middleware.Middleware
}

func (h ApiHandler) GetAuthorProfile(w http.ResponseWriter, r *http.Request) {
	profile := router.Profile{
		Name:          "ken",
		SelfIntroduce: "こんにちは、kenです。好きなアニメはBANANA FISHです。",
	}
	render.JSON(w, r, profile)
}

func (h ApiHandler) GetAuthorSecretProfile(w http.ResponseWriter, r *http.Request) {
	profile := router.Profile{
		Name:          "ken",
		SelfIntroduce: "こんにちは、kenです。最近12の倍数の年齢のときに歳男になるのだと気づきました。",
	}
	render.JSON(w, r, profile)
}
