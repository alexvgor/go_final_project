package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexvgor/go_final_project/tests"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Printf("starting app on %d port ... \n", tests.Port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", tests.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
