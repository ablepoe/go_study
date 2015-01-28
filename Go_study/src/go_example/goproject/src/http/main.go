package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func SayGoodbye(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Goodbye"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/goodbye", SayGoodbye)
	http.ListenAndServe(":8201", nil)

}
