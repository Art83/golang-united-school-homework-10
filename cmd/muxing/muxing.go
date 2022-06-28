package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func NameParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", vars["id"])
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func PostParamsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("couldn't read body: %s\n", err)
	}
	fmt.Fprintf(w, "I got message:\n%s", string(body))
}

func PostHeadersHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	i := a + b
	w.Header().Set("a+b", strconv.Itoa(i))
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", BadRequestHandler).Methods("GET")
	router.HandleFunc("/data", PostParamsHandler).Methods("POST")
	router.HandleFunc("/headers", PostHeadersHandler).Methods("POST")
	router.HandleFunc("/name/{id}", NameParamHandler).Methods("GET")
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
