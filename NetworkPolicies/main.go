package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte("Hello from " + hostname))
}

func main() {
	mux := mux.NewRouter()
	mux.Handle("/", &handler{})
	fmt.Println("Server Started")
	http.ListenAndServe(":8080", mux)

}
