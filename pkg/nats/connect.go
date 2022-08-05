package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"intern_L0/pkg/cache"
	"intern_L0/pkg/order"
	"intern_L0/pkg/services"
	"intern_L0/pkg/utils"
	"log"
	"os"
)

var Nc *nats.Conn
var Js nats.JetStreamContext

func NatsJsConnect() {
	var err error
	natsHost := os.Getenv("NATS_HOST")
	if natsHost == "" {
		natsHost = "localhost"
	}
	var url = "nats://" + natsHost + ":4222"
	Nc, err = nats.Connect(url)
	fmt.Println(err)
	var e error
	Js, e = Nc.JetStream(nats.PublishAsyncMaxPending(256))
	fmt.Println(e)
	log.Println(Js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
		NoAck:    false,
	}))
}

// Публикация  в stream
func Publish(topic string, Json []byte) {
	_, err := Js.Publish(topic, Json) //"ORDERS.*"
	if err != nil {
		fmt.Println(err)
	}
}

// Подписка на stream, получение заказа
func Subscribe(topic string) {
	_, err := Js.Subscribe(topic, func(m *nats.Msg) {
		m.Ack()
		switch topic {
		case "ORDERS.*":
			var ord order.Order
			strct := utils.JsonToStruct(m.Data, ord)
			ProcessOrder(strct)
		default:
			fmt.Println("Нет такого топика")
		}
		log.Printf("monitor service subscribes from subject:%s\n", m.Subject)
	}, nats.Durable("MONITOR"), nats.ManualAck())

	if err != nil {
	} else {
		fmt.Println("Received a JetStream message")
	}
}

func ProcessOrder(order order.Order) {

	services.SaveOrder(&order)
	cache.Cache.Store(order.OrderUID, order)

}
