package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ChatWith(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received from " + ws.Request().Host + "  " + reply

		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
     <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:8888";
        window.onload = function() {
           console.log("onload");
           try
           {
            sock = new WebSocket(wsuri);
           }catch (e) {
               alert(e.Message);
           }
           sock.onopen = function() {
              console.log("connected to " + wsuri);
           }
           sock.onerror = function(e) {
              console.log(" error from connect " + e);
           }
           sock.onclose = function(e) {
              console.log("connection closed (" + e.code + ")");
           }
           sock.onmessage = function(e) {
              console.log("message received: " + e.data);
			  var divMsg = document.getElementById('log');
		   	  divMsg.innerText = "qwerty\n" + e.data;
           }
        };
        function send() {
           var msg = document.getElementById('message').value;
		   var divMsg = document.getElementById('log');
		   divMsg.innerText = "qwerty\n" + divMsg.innerText;
           sock.send(msg);
        };
     </script>
     <h1>WebSocket chat with server </h1>
         <div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
         </div>
         <div>
           <form>
               <p>
                   Message: <input id="message" type="text" value="Hello, world!"><button onclick="send();">Send Message</button>
               </p>
           </form>
         </div>
</html>`
	io.WriteString(w, html)
}

func main() {
	//
	http.Handle("/", websocket.Handler(ChatWith))
	http.HandleFunc("/Client", Client)

	fmt.Println("listen on port 8888")
	fmt.Println("visit http://127.0.0.1:8888/")
	fmt.Println("chat with web browser(recommend: chrome)")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
		fmt.Println("ListenAndServer failed")
	}
}
