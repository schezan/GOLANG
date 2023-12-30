package main

//template data stores data from templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// Holds the email msg

type MailData struct {
	To      string
	From    string
	Subject string
	Content string
}

//creating a golang struct that matches the  js cart elements in the shop page
type Items struct {
	Name  string "json:name"
	Price int64  "json:price"
	Count int64  "json:count"
	Total string "json:total"
}

//Holds our customer Order details
type OrderData struct {
	Customername string
	Phone        string
	Address      string
	Items        string
	Price        string
	Cart         []Items
}
