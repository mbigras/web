// Credit to the following illustrative examples:
// 1. https://pkg.go.dev/net/http#hdr-Servers
// 2. https://pkg.go.dev/os/signal#Notify
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func web() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	log.Println("running as pid", os.Getpid())
	log.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	go web()
	// Block until a signal is received.
	sig := <-ch
	log.Printf("caught signal: %s; shutting down; byeee!!", sig)
}
