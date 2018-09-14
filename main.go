package main

import (
	"github.com/graphql-go/handler"
	"net/http"
)

func Register() *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:   &GSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}

func main() {
	h := Register()
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
