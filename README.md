# Golang page reloader

__I forked the repository to add go module and custom port.__

This is a small sample application implemented in Golang which can programmatically refresh a tab in a browser.
It is accompanying a blog post I've written about the subject with all the details: .

It hosts a small WebSocket service with a single `reload` endpoint, to which we can connect from the browser, and send a message every time we want it to be reloaded.

## Usage

Install dependencies:

```sh
go mod download
```

Then run go files:

```
go run *.go
```

Open a browser and press "Enter" in console.
