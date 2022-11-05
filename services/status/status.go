package status

import "net/http"

func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive"))
}
