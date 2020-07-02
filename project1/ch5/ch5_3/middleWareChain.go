package main

import (
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "#", 0)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello!"))
	logger.Println("hello request finished!")
}

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func newRouter() *Router {
	return &Router{mux: make(map[string]http.Handler)}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

func (r *Router) Start() {
	for route, hander := range r.mux {
		http.Handle(route, hander)
	}
}

func logger1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("this is middleware:logger1 !")
		next.ServeHTTP(w, r)
	})
}
func logger2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("this is middleware:logger2 !")
		next.ServeHTTP(w, r)

	})
}

func main() {
	r := newRouter()
	r.Use(logger1)
	r.Use(logger2)

	// r.Start()
	// http.ListenAndServe(":9000", nil)
	var logger3 = middleware(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("this is middleware:logger3 !")
			next.ServeHTTP(w, r)
		})
	})
	r.Use(logger3)
	r.Add("/", http.HandlerFunc(hello))

}
