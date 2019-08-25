package main

import (
	"bufio"
	"log"
	"os"
	rl "talentlessguy/golang-reload-browser/reload"
)

func main() {
	log.Println("Starting reload server.")

	rl.StartReloadServer(":3000")

	log.Println("Reload server started.")
	log.Println("Press Enter to reload the browser!")
	for {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		log.Println("Reloading browser.")
		rl.SendReload()
	}
}
