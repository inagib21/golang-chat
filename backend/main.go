package main

import (
	"fmt"
	"net/http"
	"github.com/inagib21/golang-chat/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println("websocket endpoint reached")

	conn, err:- websocket.Upgrade(w, r)

	if err!=nil{
		fmt.Fprint(w, "%+v\n",err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}


func setupRoutess(){
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request)){
		serveWs(pool,w,r)
	}
}

func main(){
	fmt.Println("Nagib full stack project")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}