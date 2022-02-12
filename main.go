package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hrasiter/chat/trace"
)

func main() {
	addr := flag.String("addr", ":8080", "The address of the application")
	flag.Parse()
	http.Handle("/", &TemplateHandler{filename: "chat.html"})
	room := NewRoom()
	room.tracer = trace.NewTracer(os.Stdout)
	http.Handle("/room", room)

	go room.run()

	fmt.Println("Server is running on address: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Error in ListenAndServer: ", err)
	}
}
