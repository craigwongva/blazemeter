package main

import (
        "net/http"
)

func main() {

        mux := http.NewServeMux()

        mux.HandleFunc("/blaze", blaze)
        http.ListenAndServe(":8076", mux)
}
