package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)
	http.ListenAndServe("127.0.0.1:8990", nil)

	port := 8990
	err := http.ListenAndServe(string(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)

	}
}

type User struct {
	FirstName string `json:"first_name" :"first_name"`
	LastName  string `json:"last_name" :"last_name"`
	IsFired   bool
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	//user := User{"Vitaliy", "Peremyslov"}
	//js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return

	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := []User{{FirstName: "Vasy", LastName: "Ivanov", IsFired: false}, {FirstName: "Vitaliy", LastName: "Peremyslov", IsFired: true}}
	//js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return

	}
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
