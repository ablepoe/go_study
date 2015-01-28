package handler

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index", http.StatusFound)
	}
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout/main.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/chats/index.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, nil)
}

func mainServerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/server.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, nil)
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received from:  " + ws.Request().Host + " " + reply
		fmt.Println("Sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func Test() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/server", mainServerHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/index", indexHandler)
	//http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/topic/", assetsHandler)
	http.ListenAndServe(":8888", nil)

}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("writer is ", w)
	//fmt.Println("request is ", r)
	//fmt.Println("r.URL.Path is ", r.URL.Path)
	//fmt.Println("r.URL.Path is ", r.URL.Path[len("/"):])
	http.ServeFile(w, r, r.URL.Path[len("/"):])
}
