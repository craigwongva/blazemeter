package main

import (
        "net/http"
)

func main() {

        mux := http.NewServeMux()

        mux.HandleFunc("/blaze", blaze)
        mux.HandleFunc("/ping", ping)
        http.ListenAndServe(":8076", mux)
}
