# Golang page reloader

__I forked the repository to add go module and custom port.__

This is a small sample application implemented in Golang which can programmatically refresh a tab in a browser.
It is accompanying a blog post I've written about the subject with all the details: .

It hosts a small WebSocket service with a single `reload` endpoint, to which we can connect from the browser, and send a message every time we want it to be reloaded.

## Install

```sh
go get -u -v github.com/talentlessguy/golang-reload-browser
```

Open a browser and press "Enter" in console.

## Example

Go code:

```go
package main

import (
	"bufio"
	"log"
	"os"
	rl "talentlessguy/golang-reload-browser"
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
```

HTML page:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Reload test page</title>
    <script>
      const tryConnectToReload = address => {
        const conn = new WebSocket(address)

        conn.onclose = () => {
          setTimeout(() => {
            tryConnectToReload(address)
          }, 2000)
        }

        conn.onmessage = evt => location.reload()
      }
      try {
        if (window.WebSocket) {
          // The reload endpoint is hosted on a statically defined port.
          tryConnectToReload('ws://localhost:3000/reload')
        } else {
          console.log( 'Your browser does not support WebSockets :(')
        }
      } catch (e) {
        console.error(`Exception during connecting to Reload: ${e}`)
      }
    </script>
  </head>
</html>
```

Then run go file:

```sh
go run main.go
```
