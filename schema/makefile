api-gen:
	oapi-codegen -package router -generate "types" openapi.yaml > ../app/router/types.gen.go
	oapi-codegen -package router -generate "spec" openapi.yaml > ../app/router/spec.gen.go
	oapi-codegen -package router -generate "chi-server" openapi.yaml > ../app/router/api.gen.go
