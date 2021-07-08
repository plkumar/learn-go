package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/plkumar/learn-go/models"
)

type todoController struct {
	userIdPattern *regexp.Regexp
}

func (tc todoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Welcome to Todo Controller"))
	if r.URL.Path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			tc.Get(w, r)
		case http.MethodPost:
			tc.Post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := tc.userIdPattern.FindStringSubmatch(r.URL.Path)
		fmt.Println(matches)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case http.MethodPut:
			tc.Put(id, w, r)
		case http.MethodDelete:
			tc.Delete(id, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (tc *todoController) Delete(id int, w http.ResponseWriter, r *http.Request) {
	err := models.RemoveByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		encodeToJson("Success", w)
	}
}

func (tc *todoController) Put(id int, w http.ResponseWriter, r *http.Request) {

	todo, err := tc.parseRequest(r)
	fmt.Println(todo, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id != todo.ID {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("id did not match"))
		return
	}

	todo, err = models.Update(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeToJson(todo, w)
}

func (tc *todoController) Post(w http.ResponseWriter, r *http.Request) {
	//par
	todo, err := tc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	todo, err = models.Add(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeToJson(todo, w)
}

func (tc *todoController) parseRequest(r *http.Request) (models.Todo, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Todo
	err := dec.Decode(&u)
	if err != nil {
		return models.Todo{}, err
	}
	return u, nil
}

func (tc *todoController) Get(w http.ResponseWriter, r *http.Request) {
	encodeToJson(models.All(), w)
}

func newTodoController() *todoController {
	return &todoController{
		userIdPattern: regexp.MustCompile(`^/todos/(\d+)/?`),
	}
}
