package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var port int = 9999
var portString = strconv.Itoa(port)

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

func Client(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Insert title here</title>
      <script type="text/javascript">
         var sock = null;
         var wsuri = "ws://127.0.0.1:` + portString + `";
         window.onload = function() {
			var msg = document.getElementById('msg');
            console.log("onload");
            sock = new WebSocket(wsuri);
            sock.onopen = function() {
               console.log("connected to " + wsuri);
            }
            sock.onclose = function(e) {
               console.log("connection closed (" + e.code + ")");
            }
            sock.onmessage = function(e) {
               console.log("message received: " + e.data);
			msg.innerText = "Server response " + e.data + "\n" + msg.innerText;
            }
         };
         function send() {
            var message = document.getElementById('message').value;
			msg.innerText = "Client msg sending " + message + "\n";
            sock.send(message);
         };
      </script>
      <h1>WebSocket Echo Test</h1>
      <form>
         <p>
            Message: <input id="message" type="text" value="Hello, world!">
         </p>
      </form>
      <button onclick="send();">Send Message</button>
</head>
<body>
	<div id="msg"></div>
</body>
</html>`
	io.WriteString(w, html)
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/Client", Client)

	if err := http.ListenAndServe(":"+portString, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
