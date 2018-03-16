package main

import (
	"net/http"
	"reflect"
)

var di = make(map[string]interface{})
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

	up := CacheUserProvider{
		users: users,
	}

	di["main.UserProvider"] = up
}

// Register registers dependencies
func Register(dep map[string]interface{}) {
	di = dep
}

// HttpDI resolves dependencies for http handler func
func HttpDI(w http.ResponseWriter, r *http.Request, fn interface{}) {
	if reflect.TypeOf(fn).Kind().String() != "func" {
		panic("third parameter must be a func")
	}

	in := reflect.ValueOf(fn).Type().NumIn()

	method := reflect.ValueOf(fn)
	params := make([]reflect.Value, in, in)
	params[0] = reflect.ValueOf(w)
	params[1] = reflect.ValueOf(r)
	for i := 2; i < in; i++ {
		param := method.Type().In(i)
		if val, ok := di[param.String()]; ok {
			params[i] = reflect.ValueOf(val)
		} else {
			panic("found dependency that is not in list")
		}
	}

	method.Call(params)
}

// DI is a helper method just for benchmarks
func diForBench(fn interface{}) {
	if reflect.TypeOf(fn).Kind().String() != "func" {
		panic("third parameter must be a func")
	}

	in := reflect.ValueOf(fn).Type().NumIn()

	method := reflect.ValueOf(fn)
	params := make([]reflect.Value, in, in)

	for i := 0; i < in; i++ {
		param := method.Type().In(i)
		if val, ok := di[param.String()]; ok {
			params[i] = reflect.ValueOf(val)
		} else {
			panic("found dependency that is not in list")
		}
	}
}
