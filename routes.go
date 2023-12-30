package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *AppConfig) http.Handler {

	mux := chi.NewRouter() //declaring my new router

	//going to use one of chi package's middleware ,the recoverer(that gracefully handles error?)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(WritetoConsole) //calling our write to function in middleware file
	mux.Use(SessionLoad)    //loads and saves the session on every requests

	//routing and calling our repository and handler functions
	mux.Get("/", Repo.Home)
	mux.Get("/shop", Repo.Shop)
	mux.Post("/shop", Repo.PostShop)
	mux.Get("/Order-Summary", Repo.OrderSummary)
	//route to file server,Allows our templates to access the static folder which holds our website's js,css and images
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
