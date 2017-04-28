package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pborman/uuid"
	"registry"
)

var (
	service = &registry.Service{
		Name: "go.micro.srv.greeter",
		Nodes: []*registry.Node{
			{
				Id:      "go.micro.srv.greeter-" + uuid.NewUUID().String(),
				Address: "localhost",
				Port:    4000,
			},
		},
	}
)

func main() {
	l, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/greeter", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintf(w, "Hello %s!", r.Form.Get("name"))
	})

	go http.Serve(l, http.DefaultServeMux)

	Register(service)

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-notify

	Deregister(service)
	l.Close()
}
