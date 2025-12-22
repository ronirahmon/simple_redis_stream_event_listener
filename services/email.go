package services

import (
	"encoding/json"
	"fmt"
	"simple_redis_stream_event_listener/models"
)

func PrintStatus(status *models.Status) (err error) {

	fmt.Println("Print status")
	js, err := json.Marshal(status)

	fmt.Println(string(js))

	return
}
