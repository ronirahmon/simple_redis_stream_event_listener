package main

import (
	"fmt"
	"log"
	"os"

	"simple_redis_stream_event_listener/event"
	"simple_redis_stream_event_listener/handlers"
	"simple_redis_stream_event_listener/utils"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config := utils.Configs{
		AppName:        fmt.Sprintf("%s v%s", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION")),
		AppEnvironment: os.Getenv("APP_ENVIRONMENT"),
		Redis: utils.DB{
			DbName:             "1",
			Host:               os.Getenv("REDIS_HOST"),
			StreamSubject:      os.Getenv("REDIS_STREAM_SUBJECT"),
			StreamConsumeGroup: os.Getenv("REDIS_STREAM_CONSUMER_GROUP"),
		},
	}

	_ = utils.NewConfig(config)

}

func main() {

	log.Println("Redis Event Consumer Start")

	h := handlers.Handler{}

	e := event.New()
	e.AddEvent("status", h.ChangeStatusEvent)

	e.Start()
}
