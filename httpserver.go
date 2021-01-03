package main

import (
"net/http"
"fmt"
)


func httpserver(port *string, directory *string){
  fmt.Println("--HTTP-Server--")
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	fmt.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	http.ListenAndServe(":"+*port, nil)
}
