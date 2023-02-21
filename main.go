package main

import (
	"net/http"

	admincontroller "github.com/husnain3184/creative-studio/controllers/Admin"
	"github.com/husnain3184/creative-studio/controllers/Admin/manageourangelscontroller"
	"github.com/husnain3184/creative-studio/controllers/Admin/managespeakerscontroller"
	homecontroller "github.com/husnain3184/creative-studio/controllers/Frontend/home"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)
	// http.HandleFunc("/about-us", homecontroller.About)
	// http.HandleFunc("/contact-us", homecontroller.Contact)
	// http.HandleFunc("/speaker", homecontroller.Speaker)
	// http.HandleFunc("/schedule", homecontroller.Schedule)
	// http.HandleFunc("/blog", homecontroller.Blog)
	// http.HandleFunc("/blog-details", homecontroller.BlogDetail)
	http.HandleFunc("/admin", admincontroller.AdminLogin)
	http.HandleFunc("/register", admincontroller.Register)
	http.HandleFunc("/dashboard", admincontroller.Dashboard)
	http.HandleFunc("/manage-speakers", managespeakerscontroller.Index)
	http.HandleFunc("/manage-speakers/store", managespeakerscontroller.Store)
	http.HandleFunc("/manage-speakers/get_form", managespeakerscontroller.GetForm)
	http.HandleFunc("/manage-speakers/delete", managespeakerscontroller.Delete)

	http.HandleFunc("/manage-angels", manageourangelscontroller.Index)
	http.HandleFunc("/manage-angels/store", manageourangelscontroller.Store)
	http.HandleFunc("/manage-angels/get_form", manageourangelscontroller.GetForm)
	http.HandleFunc("/manage-angels/delete", manageourangelscontroller.Delete)

	http.HandleFunc("/logout", admincontroller.Logout)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)
}
