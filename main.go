package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Go Devops")

	fmt.Println("server is on 3600")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h2>Welcome to go devops</h2>"))
	})

	log.Fatal(http.ListenAndServe(":3600", nil))

}
