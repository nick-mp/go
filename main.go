package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and"+
		" he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	Bob := User{"Bob", 25, -50, 4.2, 0.8}
	//Bob.setNewName("Alex")
	//fmt.Fprintf(page, Bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute (page, Bob)
}



func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// var Bob user =
	//Bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

	handleRequest()
}
