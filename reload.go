package reload

import (
	"bytes"
	"log"
	"net/http"
)

var hub *Hub

func StartReloadServer(port string) {
	hub = newHub()
	go hub.run()
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	go StartServer(port)
	log.Println("Reload server listening at", port)
}

func StartServer(port string) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Println("Failed to start up the Reload server: ", err)
		return
	}
}

func SendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	hub.broadcast <- message
}
