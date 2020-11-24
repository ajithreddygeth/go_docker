package main

import (
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./mywebsite")))

    log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}