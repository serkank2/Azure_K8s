package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Testone"))
}

func main() {
	mux := mux.NewRouter()
	mux.Handle("/", &handler{})
	fmt.Println("Server Started")
	http.ListenAndServe(":9091", mux)

}
