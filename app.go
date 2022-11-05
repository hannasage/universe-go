package main

import (
	"net/http"
	"universe-go/services/status"
)

func main() {
	http.HandleFunc("/status", status.StatusCheck)
	http.ListenAndServe(":8080", nil)
}
