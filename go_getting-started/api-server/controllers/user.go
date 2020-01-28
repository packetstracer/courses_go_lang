package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

// The user controller
type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, error := strconv.Atoi(matches[1])

		if error != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}

	}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJSON(models.GetUsers(), w)
}

func (uc userController) get(ID int, w http.ResponseWriter) {
	u, err := models.GetUserByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	EncodeResponseAsJSON(u, w)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	u, error := uc.parseRequest(r)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse the user object"))
		return
	}

	u, error = models.AddUser(u)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	EncodeResponseAsJSON(u, w)
}

func (uc userController) put(ID int, w http.ResponseWriter, r *http.Request) {
	u, error := uc.parseRequest(r)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse the user object"))
		return
	}

	u, error = models.UpdateUser(ID, u)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	EncodeResponseAsJSON(u, w)
}

func (uc userController) delete(ID int, w http.ResponseWriter) {
	error := models.DeleteUserByID(ID)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)

	var u models.User
	error := dec.Decode(&u)

	if error != nil {
		return models.User{}, error
	}

	return u, nil
}

// User controlle constructor
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
