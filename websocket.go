package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println(os.Args[0]) //第一个参数，显示文件路径
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Println(err)
	}
}

func Echo(ws *websocket.Conn) {
	for {
		var raply string
		if err := websocket.Message.Receive(ws, &raply); err != nil { //get infomation,write in adress
			log.Println("can't receive")
			break
		}
		msg := "Received:" + raply
		log.Println(msg)
		if err := websocket.Message.Send(ws, "come back infomation"); err != nil { //send infomation
			log.Println("can't send")
			break
		}
	}
}
