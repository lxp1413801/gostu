package main

// /root/goLib/src/github.com/golang/net
import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error
	var reply string
	for {
		err = websocket.Message.Receive(ws, &reply)

		fmt.Printf("new conn:%s,%s\r\n", ws.RemoteAddr().String(), ws.LocalAddr().String())
		//fmt.Printf("conn's id: ", ws.Remote...)
		if err != nil {
			fmt.Println("recevied err :", err.Error())
			break
		}
		msg := "received :" + reply
		fmt.Println("received back from client:" + msg)
		msg = reply
		err = websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("send err :", err.Error())
			break
		}
	}
}

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mothod", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("websocket.html")
		t.Execute(w, nil)
	} else {
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
	}
	fmt.Printf("remote addr:%s\r\n", r.RemoteAddr)

}
func main() {

	http.Handle("/ws", websocket.Handler(Echo))
	http.HandleFunc("/web", web)
	if err := http.ListenAndServe(":1112", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
