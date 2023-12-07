package middleware

import (
	"log"
	"net/http"

	"golang.org/x/exp/slices"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/go-chi/render"
)

type Middleware interface {
	ForbidMySecret(swagger *openapi3.T) func(next http.Handler) http.Handler
}

type middleware struct{}

func NewMiddlewares() Middleware {
	return middleware{}
}

func (m middleware) ForbidMySecret(swagger *openapi3.T) func(next http.Handler) http.Handler {
	// Swagger仕様を基にしたrouterを生成
	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 受け取ったHTTPリクエストに基づいて、対応するルートを探す
			route, _, err := router.FindRoute(r)
			if err != nil {
				log.Fatal(err)
				return
			}
			// ルートが見つかった場合、そのルートのHTTPメソッドに関連付けられているタグを取得
			tags := route.PathItem.GetOperation(r.Method).Tags
			// タグが「OnlyPremium」を含むかどうかを確認
			isOnlyPremium := slices.Contains(tags, "OnlyPremium")
			if isOnlyPremium {
				/*
					ここで本当ならプレミアム会員かどうかを確認
					（ここではその実装を割愛して常にforbiddenを返す）
				*/
				render.Status(r, 403)
				render.JSON(w, r, "見せられないよ！")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
