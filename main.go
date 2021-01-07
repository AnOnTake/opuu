package main

import (
"fmt"
"os"
"flag"
)

var BuildNumber string
var BuildDate string
var GitHash string

func oops(){
  fmt.Println("usage: opuu <arguments>")
  fmt.Println("'opuu -c help' for help")
  os.Exit(1)
}


func main(){
  fmt.Println("opuu! (OnTake Power User Utilities)");
  fmt.Printf("Build Date: %s\n",BuildDate)
  fmt.Printf("Commit: %s\n",GitHash)
  fmt.Printf("Build: %s\n",BuildNumber)
  fmt.Printf("\n")

  command := flag.String("c", "null", "command")
  port := flag.String("p", "8080", "port to serve on")
  directory := flag.String("d", ".", "the directory of static file to host")
  flag.Parse()
  if (*command=="http-server"){
      httpserver(port,directory)
  } else{
    oops();
  }
}
