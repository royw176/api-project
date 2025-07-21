package main

import (
    "net/http"
    "api/router"
)

func main() {
    r := router.NewRouter()
    http.ListenAndServe(":3000", r)
}