package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Repository struct {
	App *AppConfig
}

var Repo *Repository //the repository used by the handlers

//Creates a new repository
func NewRepo(a *AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//Sets the repository for our handlers
func Newhandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//remoteIP := r.RemoteAddr
	//m.App.Session.Put(r.Context(), "remote ip", remoteIP)
	RenderTemplate(w, r, "home.page.tmpl", &TemplateData{}) //calling our render function from render to build the webpage
}

//we have added reciever of type repository to both Home and about
// shop is the handler for the shop page
func (m *Repository) Shop(w http.ResponseWriter, r *http.Request) {

	//var emptyOrder OrderData

	//data := emptyOrder

	RenderTemplate(w, r, "shop.page.tmpl", &TemplateData{
		//Data: data,
	})

}

//Calling the post function  that triggers after Customer Orders
func (m *Repository) PostShop(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//grabing the strigified cart that I posted into the "cart" form
	cart := r.Form.Get("cart")

	//Unmarshalling json file to golang Array
	var items []Items
	err = json.Unmarshal([]byte(cart), &items)
	if err != nil {
		fmt.Println(err)

	}

	//Fetching data from html forms to the webserver
	Ordersummary := OrderData{
		Customername: r.Form.Get("Name"),
		Phone:        r.Form.Get("Phone"),
		Address:      r.Form.Get("Address"),
		Items:        r.Form.Get("totalitems"),
		Price:        r.Form.Get("cartTotal"),
		Cart:         items,
	}

	var emailMessage string

	emailMessage = fmt.Sprintln("Customer name :", Ordersummary.Customername)
	emailMessage += fmt.Sprintln("Phone  :", Ordersummary.Phone)
	emailMessage += fmt.Sprintln("Address :", Ordersummary.Address)

	emailMessage += fmt.Sprintln("Cart :")
	emailMessage += fmt.Sprintln(" ")
	emailMessage += fmt.Sprintf("%-24s%10s%6s%12s", "Item Name", "Item Price", "Count", "Item total")
	emailMessage += fmt.Sprintln(" ")

	for _, i := range items {
		var t = i
		emailMessage += fmt.Sprintf("%-24s%10s%6s%12s", t.Name, strconv.FormatInt(int64(t.Price), 10), strconv.FormatInt(int64(t.Count), 10), t.Total)
		emailMessage += fmt.Sprintln(" ")
	}
	emailMessage += fmt.Sprintf("              Total Items : %s    Total Price : %s", Ordersummary.Items, Ordersummary.Price)

	subject := fmt.Sprintf("New Order from %s", Ordersummary.Customername)
	msg := MailData{
		To:      "Sheezanalamgir@gmail.com",
		From:    "wattehgreens@yahoo.com",
		Subject: subject,
		Content: emailMessage,
	}

	m.App.MailChan <- &msg

	m.App.Session.Put(r.Context(), "order", Ordersummary)

	http.Redirect(w, r, "/Order-Summary#order", http.StatusSeeOther)

}

func (m *Repository) OrderSummary(w http.ResponseWriter, r *http.Request) {
	Orderdata, ok := m.App.Session.Get(r.Context(), "order").(OrderData)
	if !ok {
		log.Println("Cannot get item from session")
		return
	}

	data := make(map[string]interface{})
	data["order"] = Orderdata

	RenderTemplate(w, r, "Order-Summary.page.tmpl", &TemplateData{
		Data: data,
	})
}
