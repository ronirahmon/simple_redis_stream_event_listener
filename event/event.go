package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"simple_redis_stream_event_listener/utils"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/redis/go-redis/v9"
	"github.com/rs/xid"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

type Event struct {
	keyFunc       map[string]HandlerFunc
	Redis         *redis.Client
	StreamSubject string
	CusumerGroup  string
	Firebase      *firebase.App
}

type HandlerFunc func(c context.Context) error

func New() *Event {

	err := utils.NewRedisClient()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nRegistered events :")

	return &Event{
		Redis:         utils.GetRedisClient(),
		StreamSubject: utils.GetConfig().Redis.StreamSubject,
		CusumerGroup:  utils.GetConfig().Redis.StreamConsumeGroup,
		keyFunc:       map[string]HandlerFunc{},
	}
}

func (e *Event) AddEvent(key string, h HandlerFunc) {
	fmt.Printf("   * Event :  %s\n", key)
	e.keyFunc[key] = h
}

func (e *Event) RunEvent(messageId string, data *map[string]interface{}, ch chan error) {
	log.Printf("%s %s event triggered%s", Red, (*data)["action"].(string), Reset)
	// var keyctx contextKeyData = "data"
	keyctx := KeyContextValue("data")
	js, _ := json.Marshal(data)
	fmt.Println(string(js))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, keyctx, data)

	if _, ok := (*data)["action"]; !ok {
		ch <- errors.New("action not found")
		return
	}
	if _, ok := e.keyFunc[(*data)["action"].(string)]; !ok {
		ch <- errors.New("Event not found")
		return
	}
	fmt.Println((*data)["action"].(string), " event start")
	err := e.keyFunc[(*data)["action"].(string)](ctx)
	if err != nil {
		ch <- err
		return
	}
	err = e.Redis.XAck(ctx, e.StreamSubject, e.CusumerGroup, messageId).Err()

	fmt.Printf("\n%sEvent Run  SUCCESS%s\n\n", Green, Reset)

	data = nil
	ch <- err
}

func (e *Event) Start() {
	fmt.Println("")
	log.Println("Start Event ")

	log.Println("Create cunsumer group")
	err := e.Redis.XGroupCreate(context.Background(), e.StreamSubject, e.CusumerGroup, "0").Err()
	if err != nil {
		log.Println(err)
	}
	uniqId := xid.New().String()

	ch := make(chan error, 10)
	fmt.Println("\nlisten event start..")
	fmt.Println("")
	for {
		entries, err := e.Redis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    e.CusumerGroup,
			Consumer: uniqId,
			Streams:  []string{e.StreamSubject, ">"},
			Count:    1,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			fmt.Println("ERR :", err)
			time.Sleep(2 * time.Second)
			continue

		}

		for _, message := range entries[0].Messages {
			// ch <- err
			go e.RunEvent(message.ID, &message.Values, ch)
			err := <-ch
			if err != nil {
				fmt.Println("Error: ", err)
			}

		}
		entries = nil
		err = nil
	}
}
