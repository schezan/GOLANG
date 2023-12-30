package main

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the template cache as a global variable (app wide use)
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template //Holds html... to be precise go templates (go webpages)
	Infolog       *log.Logger
	InProduction  bool //this variable was made to handle security/inproduction  status
	Session       *scs.SessionManager
	MailChan      chan *MailData //creating a channel that holds our maildata

}
