package main

import (
        "fmt"
        "os"
        "log"
        "net/http"
)
func handler(w http.ResponseWriter, r "http.Request){
    name, err := os.Hostname()
    if err != nil {
      panic(err)
    }
    
    fmt.fprintln(w, "hostname:", name)
}
func main() {
    fmt.Fprintln(os.Stdout,"Starting GoApp Server......")
      http.HandleFunc("/",handler)
      log.Fatal(http.listenAndServe(":8080",nil))
}
