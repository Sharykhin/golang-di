package main

import (
	"fmt"
	"net/http"
)

type (
	// Middleware is a middleware func for income requests
	Middleware func(h http.Handler) http.Handler
)

// Chain is a helper fun that allows chan multiple middlewares
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Income request: %s\n", r.RequestURI)
		h.ServeHTTP(w, r)
	})
}
