package main

import (
	"intern_L0/pkg/generator"
	"intern_L0/pkg/nats"
	"intern_L0/pkg/utils"
	"time"
)

//Бесконечная генерация заказов
func main() {
	nats.NatsJsConnect() // Подключаем JetStream
	for {
		order := generator.GenerateOrder()
		OrdJson := utils.StructToJson(order)
		nats.Publish("ORDERS.*", OrdJson)
		time.Sleep(60 * time.Second)
	}
}
