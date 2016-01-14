package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	instance := os.Getenv("CF_INSTANCE_INDEX")
	hostname, _ := os.Hostname()

	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<br>Hello %s from My Cloud Foundry Instance....</br>", r.URL.Path[1:])
	fmt.Fprintf(w, "<br>Instance# :%s</br>", instance)
	fmt.Fprintf(w, "<br>Hostname :%v</br>", hostname)
	fmt.Fprintf(w, "<br>Hostname :%v</br>", hostname)
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func template(w http.ResponseWriter, r *http.Request) {

}

const (
	defaultPort = "8080"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+vn", defaultPort)
		port = defaultPort
	}

	router := mux.NewRouter()
	router.HandleFunc("/info", handler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: %v", err)
	}
}
