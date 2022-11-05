package main

import (
	"net/http"

	"github.com/kevinhaube/universe-go/services/status"
)

func main() {
	http.HandleFunc("/status", status.StatusCheck)
	http.ListenAndServe(":8080", nil)
}
