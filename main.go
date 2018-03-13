package main

import (
	"encoding/json"
	"log"
	"net/http"

	"reflect"

	"fmt"

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

var users []User

func init() {
	user1 := User{
		Id:   "1",
		Name: "John",
	}
	user2 := User{
		Id:   "2",
		Name: "Mark",
	}

	users = append(users, user1)
	users = append(users, user2)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		in := reflect.ValueOf(handler).Type().NumIn()

		method := reflect.ValueOf(handler)
		params := make([]reflect.Value, in, in)
		//fmt.Println(reflect.ValueOf(handler).Type().In(2))
		for i := 0; i < in; i++ {
			param := reflect.ValueOf(handler).Type().In(i)
			fmt.Println(param.String())
			if param.Kind().String() == "interface" && param.String() == "main.UserProvider" {
				up := CacheUserProvider{
					users: users,
				}
				params[i] = reflect.ValueOf(up)
			} else {
				params[i] = reflect.ValueOf(param)
			}

		}
		fmt.Println(params)
		method.Call(params)
		//up := CacheUserProvider{
		//	users: users,
		//}
		//handler(w, r, up)
	})
	log.Fatal(http.ListenAndServe(":8003", nil))
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
