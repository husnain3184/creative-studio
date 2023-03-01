package ManageOurUsersController

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
	"github.com/husnain3184/creative-studio/models"
	"golang.org/x/crypto/bcrypt"
)

var UsersModel = models.NewUsers()

type UserInput struct {
	Username string
	Password string
	Email    string
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/backend/Login/userlogin.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		UserInput := &UserInput{

			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.Admin
		// fmt.Println(user)
		UsersModel.Where(&user, "username", UserInput.Username)

		var message error
		if UserInput.Username == "" {

			message = errors.New("username and password required")
		} else {

			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			log.Println(errPassword)

			if errPassword != nil {

				message = errors.New("username and password not matched")
			}

		}

		if message != nil {

			data := map[string]interface{}{

				"error": message,
			}
			temp, _ := template.ParseFiles("views/backend/Login/userlogin.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Name
			session.Values["name"] = user.Name

			session.Save(r, w)

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

		}

	}

}

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}
	// fmt.Println(data)

	temp, _ := template.ParseFiles("views/backend/Users/index.html", "views/backend/Dashboard/dashboardlayout.html")
	temp.Execute(w, data)

}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/backend/Users/data.html")

	var users []entities.Users
	err := UsersModel.FindAll(&users)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"users": users,
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
			"title": "Add users Data",
		}
	} else {

		var users entities.Users
		err := UsersModel.Find(id, &users)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title": "Edit users Data",
			"users": users,
		}
	}
	temp, _ := template.ParseFiles("views/backend/Users/form.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseMultipartForm(10 << 20)
		var users entities.Users
		users.Name = r.Form.Get("name")
		users.Username = r.Form.Get("Username")
		users.Email = r.Form.Get("email")
		users.Password = r.Form.Get("password")
		users.Roll = r.Form.Get("roll")
		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {

			err := UsersModel.Create(&users)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}

		} else {

			users.Id = id
			err := UsersModel.Update(users)
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
	err = UsersModel.Delete(id)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"message": "Our users Delete Successfully",
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

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
