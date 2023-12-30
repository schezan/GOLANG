package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//Everytime someone hits a page,write this line to the console.
func WritetoConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)

	})
}

//adds csrf protection on every POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)        //creates a new handler for us
	csrfHandler.SetBaseCookie(http.Cookie{ //this one is to make sure the token it generates is in a per page basis

		HttpOnly: true,
		Path:     "/", //applying cookie path to the entire site hence "/"
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
