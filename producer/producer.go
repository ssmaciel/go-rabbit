package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/wesleywillians/go-rabbitmq/queue"
	"html/template"
	"log"
	"net/http"
)


type User struct {
	Name   string
	Email string
}

type Result struct {
	Status string
}

func init() {
	log.Println("Init producer")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	user := User{
		Name:   name,
		Email: email,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error parsing to json")
	}

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	err = rabbitMQ.Notify(string(jsonUser), "application/json", "users.ex", "")
	if err != nil {
		log.Fatal("Error sending message to the queue")
	}

	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, "")
}
