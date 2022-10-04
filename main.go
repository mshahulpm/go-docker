package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn
var ec *nats.EncodedConn

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	nc, err = nats.Connect(os.Getenv("NATS_URI"))
	ec, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	defer nc.Close()
	defer ec.Close()

	fmt.Println("Go Devops")

	fmt.Println("server is on 3600")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
		<h2>Welcome to go devops2</h2>

		<br>
		<button onclick="createUser()">Create a random user in user service</button>
		<br>
		<br>
		<br>
		<script>
	   async function createUser(){
		const res = await fetch('/add-user',{
		  method:'POST'
		})
		const data = await res.text() 
		console.log(data) 
		alert('user created')
		}
		</script>
		`)
	})

	http.HandleFunc("/add-user", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		ec.Publish("create-user", nil)
		w.Write([]byte("User-created"))
	})

	log.Fatal(http.ListenAndServe(":3600", nil))

}
