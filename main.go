package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/satori/go.uuid"
)

func pub() {
	topic := "user.created"

	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		u, _ := uuid.NewV4()
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("[%s] %d: user-id: %s", time.Now().String(), i, u.String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub]", topic, string(msg.Body))
		}
		i++
	}
}

func sub() {
	topic := "user.created"

	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	_, err := broker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub]", topic, "body:", string(p.Message().Body), "action:", "DO FANCY SHIT WITH THE USER")
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	forever := make(chan struct{})

	go pub()
	go sub()

	<-forever
}
