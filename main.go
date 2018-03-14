package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserProvider interface {
	GetUser(id string) (*User, error)
}

type CacheUserProvider struct {
	users []User
}

func (c CacheUserProvider) GetUser(id string) (*User, error) {
	for _, u := range c.users {
		if u.Id == id {
			return &u, nil
		}
	}

	return nil, errors.New("no user found with id " + id)
}

func main() {

	http.HandleFunc("/di", func(w http.ResponseWriter, r *http.Request) {
		HttpDI(w, r, handler)
	})

	http.HandleFunc("/di2", func(w http.ResponseWriter, r *http.Request) {
		up := CacheUserProvider{
			users: users,
		}
		handler(w, r, up)
	})

	http.Handle("/hello", Chain(http.HandlerFunc(index), Logger))
	log.Fatal(http.ListenAndServe(":8003", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handler(w http.ResponseWriter, r *http.Request, up UserProvider) {
	id := r.FormValue("id")
	user, err := up.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
