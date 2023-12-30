package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":80"

var app AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	//registering what i am going to put in sessions
	gob.Register(OrderData{})
	app.InProduction = false
	//running our mail channel
	run()
	//closing our mail channel
	defer close(app.MailChan)

	listenForMail()
	session = scs.New() //created a new session that is going to last 24 hrs
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true                  //will the session persist after closing the web browser or window
	session.Cookie.SameSite = http.SameSiteLaxMode //how strict you wanna be about what site this cookie applies to
	session.Cookie.Secure = app.InProduction       //makes sure the session is encrypted and from https and not from http(8080 isnt secure anyways so we set it to false)

	app.Session = session

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := NewRepo(&app) // created a repository variable

	Newhandler(repo) //passing it back to the handlers

	//NewTemplates(&app)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	//making new variable called serv/srv
	//its gonna be set to a pointer to http server that has some fields
	// those fields are firstly the address which is our port number,also gonna pass
	//them our handlers and pass em our routes with a reference to  app
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err) //if there is an error shut it down
}

//Function to Run maildata Channels  (check out golang Channels)
func run() {
	mailChan := make(chan *MailData)
	app.MailChan = mailChan

}
