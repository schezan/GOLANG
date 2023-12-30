package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{} //map of functions i can use on a template(currently empty)

//var app *AppConfig

//NewTemplates sets the template for the new package

//func NewTemplates(a *AppConfig) {

//	app =a
//}

func addDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) error {
	// get the template cache from the app config
	var tc map[string]*template.Template //declaring the variable that will hold my template cache
	if app.UseCache {
		tc = app.TemplateCache //if usecache is true,store in app.tc otherwise create a new one
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl] //store template name (tmpl) into t(the ok var is a bool to check if said template exists or not)
	if !ok {
		log.Fatal("could not get template from template cache") //shutdown app if there is an error
	}

	buf := new(bytes.Buffer) //var "buf" just holds some bytes

	td = addDefaultData(td, r)

	_ = t.Execute(buf, td) //take the "t" var,passed it the template data(hence td),store the value to the "buf" variable

	_, err := buf.WriteTo(w) //writes to the response writer and gets the number of byters and an error in return
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
	return nil

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl") //grab  the pages in temp folder with .page.tmpl
	if err != nil {
		return myCache, err //if there is an error return the error along with the  cache
	}

	for _, page := range pages { //for loop through pages
		name := filepath.Base(page)                                     //gets the name of the page(like about/home atm)without extension like .page.tmpl
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) //create tmpl set,set a name,set ability to implement functions(from the funcmap structure above)
		if err != nil {
			return myCache, err //if there is an error return the error along with the  cache
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl") //lookfor anyfiles that matches .layout.tmpl
		if err != nil {
			return myCache, err //if there is an error return the error along with the  cache
		}

		if len(matches) > 0 { //if something matches,look for the tmpl set ts,parse it for errors
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err //if there is an error return the error along with the  cache
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
