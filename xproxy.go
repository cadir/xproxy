package main

import (
       "log"
       "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "text/plain")
     w.Write([]byte("This is an example server.\n"))
}

func main() {
     http.HandleFunc("/", handler)
     err := http.ListenAndServe(":8080", nil)
     if err != nil {
       log.Fatal("ListenAndServe: ", err)
     }
}