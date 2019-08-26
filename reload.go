package reload

import (
	"bytes"
	"log"
	"net/http"
)

var hub *Hub

// StartReloadServer - start reloading server with websockets
func StartReloadServer(port string) {
	hub = newHub()
	go hub.run()
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	go StartServer(port)
}

// StartServer - launch static server
func StartServer(port string) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Println("Failed to start up dev server: ", err)
		return
	}
}

// SendReload - send socket with message "reload"
func SendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	hub.broadcast <- message
}
