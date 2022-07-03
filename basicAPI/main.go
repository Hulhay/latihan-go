package main

import (
	f "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type User struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	LoginDate string `json:"LoginDate"`
}

var Users []User

func home(w http.ResponseWriter, r *http.Request) {
	f.Println("home")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	for _, user := range Users {
		if user.Id == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var user User

	json.Unmarshal(reqBody, &user)
	
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	for index, user := range Users {
		if user.Id == key {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	for index, user := range Users {
		if user.Id == key {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var user User

	json.Unmarshal(reqBody, &user)
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func handleRequests() {
	route := mux.NewRouter().StrictSlash(true)
	
	route.HandleFunc("/", home)
	route.HandleFunc("/users", getAllUsers).Methods("GET")
	route.HandleFunc("/users/{Id}", getUserById).Methods("GET")
	route.HandleFunc("/users", addUser).Methods("POST")
	route.HandleFunc("/users/{Id}", deleteUserById).Methods("DELETE")
	route.HandleFunc("/users/{Id}", updateUserById).Methods("PUT")
	
	http.ListenAndServe(":3000", route)
}

func main() {
	
	f.Println("starting web server at http://localhost:3000/")
	
	Users = []User{
		{Id: "1", Name: "Prasetyo", LoginDate: "Tue Nov 10 23:00:00 2009"},
		{Id: "2", Name: "Wahyu", LoginDate: "Tue Jun 11 23:00:00 2010"},
	}
	handleRequests()

}