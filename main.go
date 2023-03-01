package main

import (
	"net/http"

	admincontroller "github.com/husnain3184/creative-studio/controllers/Admin"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurAngelsController"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurBlogController"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurPortfolioController"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurServiceController"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurTestimonialController"
	"github.com/husnain3184/creative-studio/controllers/Admin/ManageOurUsersController"
	homecontroller "github.com/husnain3184/creative-studio/controllers/Frontend/home"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/contact-us", homecontroller.SendEmail)
	// http.HandleFunc("/blog-details", homecontroller.BlogDetail)
	http.HandleFunc("/admin", admincontroller.AdminLogin)
	http.HandleFunc("/register", admincontroller.Register)
	http.HandleFunc("/dashboard", admincontroller.Dashboard)

	http.HandleFunc("/users/logout", ManageOurUsersController.Logout)
	http.HandleFunc("/users", ManageOurUsersController.UserLogin)
	http.HandleFunc("/manage-users", ManageOurUsersController.Index)
	http.HandleFunc("/manage-users/store", ManageOurUsersController.Store)
	http.HandleFunc("/manage-users/get_form", ManageOurUsersController.GetForm)
	http.HandleFunc("/manage-users/delete", ManageOurUsersController.Delete)

	http.HandleFunc("/manage-angels", ManageOurAngelsController.Index)
	http.HandleFunc("/manage-angels/store", ManageOurAngelsController.Store)
	http.HandleFunc("/manage-angels/get_form", ManageOurAngelsController.GetForm)
	http.HandleFunc("/manage-angels/delete", ManageOurAngelsController.Delete)

	http.HandleFunc("/manage-service", ManageOurServiceController.Index)
	http.HandleFunc("/manage-service/store", ManageOurServiceController.Store)
	http.HandleFunc("/manage-service/get_form", ManageOurServiceController.GetForm)
	http.HandleFunc("/manage-service/delete", ManageOurServiceController.Delete)

	http.HandleFunc("/manage-portfolio", ManageOurPortfolioController.Index)
	http.HandleFunc("/manage-portfolio/store", ManageOurPortfolioController.Store)
	http.HandleFunc("/manage-portfolio/get_form", ManageOurPortfolioController.GetForm)
	http.HandleFunc("/manage-portfolio/delete", ManageOurPortfolioController.Delete)

	http.HandleFunc("/manage-testimonial", ManageOurTestimonialController.Index)
	http.HandleFunc("/manage-testimonial/store", ManageOurTestimonialController.Store)
	http.HandleFunc("/manage-testimonial/get_form", ManageOurTestimonialController.GetForm)
	http.HandleFunc("/manage-testimonial/delete", ManageOurTestimonialController.Delete)

	http.HandleFunc("/manage-blogs", ManageOurBlogController.Index)
	http.HandleFunc("/manage-blogs/store", ManageOurBlogController.Store)
	http.HandleFunc("/manage-blogs/get_form", ManageOurBlogController.GetForm)
	http.HandleFunc("/manage-blogs/delete", ManageOurBlogController.Delete)

	http.HandleFunc("/logout", admincontroller.Logout)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)
}
