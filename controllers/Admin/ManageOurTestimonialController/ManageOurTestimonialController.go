package ManageOurTestimonialController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/husnain3184/creative-studio/entities"
	"github.com/husnain3184/creative-studio/models"
)

var TestimonialModel = models.NewTestimonial()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}
	// fmt.Println(data)

	temp, _ := template.ParseFiles("views/backend/OurTestimonial/index.html", "views/backend/Dashboard/dashboardlayout.html")
	temp.Execute(w, data)

}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/backend/OurTestimonial/data.html")

	var testimonial []entities.Testimonial
	err := TestimonialModel.FindAll(&testimonial)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"testimonial": testimonial,
	}
	// fmt.Println(data)

	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {

		data = map[string]interface{}{
			"title": "Add testimonial Data",
		}
	} else {

		var testimonial entities.Testimonial
		err := TestimonialModel.Find(id, &testimonial)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title":       "Edit testimonial Data",
			"testimonial": testimonial,
		}
	}
	fmt.Println(data)
	temp, _ := template.ParseFiles("views/backend/OurTestimonial/form.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseMultipartForm(10 << 20)
		var testimonial entities.Testimonial
		testimonial.Name = r.Form.Get("name")
		testimonial.Details = r.Form.Get("details")
		id, errs := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			fmt.Println(r.FormFile("image"))
			return
		}

		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)

		var imagename = time.Now().UnixNano()
		dst, err := os.Create(fmt.Sprintf("./static/uploads/%d%s", imagename, filepath.Ext(handler.Filename)))
		var dbimagename = fmt.Sprintf("%d%s", imagename, filepath.Ext(handler.Filename))
		testimonial.Image = dbimagename
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data map[string]interface{}

		if errs != nil {

			errs := TestimonialModel.Create(&testimonial)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, errs.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}

		} else {

			testimonial.Id = id
			err := TestimonialModel.Update(testimonial)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{

				"message": "Data Successfully Updated",
				"data":    template.HTML(GetData()),
			}

		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

	if err != nil {

		panic(err)
	}
	err = TestimonialModel.Delete(id)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"message": "Our testimonial Delete Successfully",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func RepsonseError(w http.ResponseWriter, code int, message string) {

	ResponseJson(w, code, map[string]string{"error": message})

}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {

	// fmt.Println("Payload", payload)
	response, _ := json.Marshal(payload)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
