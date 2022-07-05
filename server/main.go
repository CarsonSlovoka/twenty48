package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var (
	Debug = true
	Port  = "0"
)

func init() {
	if debug := os.Getenv("Debug"); debug == "0" || debug == "false" {
		Debug = false
	}
	if port := os.Getenv("Port"); port != "0" {
		Port = port
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootDir := "./pages"
		http.FileServer(http.Dir(rootDir)).ServeHTTP(w, r)
	})

	var server http.Server
	if Debug {
		server = http.Server{Addr: "127.0.0.1:0", Handler: mux}
	} else {
		server = http.Server{Addr: ":0", Handler: mux}
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("using port number:%d", listener.Addr().(*net.TCPAddr).Port)
	fmt.Printf("Addr:%q\n", listener.Addr().String())

	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
