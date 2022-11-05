package status

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StatusCheck(w http.ResponseWriter, r *http.Request) {
	services := []string{"thisService", "thatService", "theOtherService"}
	var statuses []string

	checkChannel := make(chan string)

	go getFakeStatusesWorker(services, checkChannel)

	for check := range checkChannel {
		statuses = append(statuses, check)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func getFakeStatusesWorker(services []string, c chan string) {
	for i, service := range services {
		fmt.Printf("Checking: %s\n", service)
		c <- fmt.Sprintf("%d-%s: alive", i, service)
	}
	close(c)
}
