package main

import (
	"net/http"

	"matrix-web/views"

	"github.com/gorilla/mux"
)

const portNum = ":3000"

var homeView *views.View
var contactView *views.View

//home is the handler function for the / route
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

//contact is the handler function for the /contact route
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

//Must - a helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(portNum, r)
}
