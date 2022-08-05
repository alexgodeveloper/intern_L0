package main

import (
	"embed"
	"fmt"
	"html/template"
	"intern_L0/pkg/cache"
	"intern_L0/pkg/database"
	"intern_L0/pkg/nats"
	"intern_L0/pkg/order"
	"log"
	"net/http"
)

//go:embed index.html
var html embed.FS
var templates = template.Must(template.ParseFS(html, "*.html"))

func main() {
	database.InitDb(database.Dsn)      // Инициализация БД
	cache.SetCacheFromDb(&cache.Cache) // "Греем" кеш
	nats.NatsJsConnect()               // Подключаем JetStream
	nats.Subscribe("ORDERS.*")         // Получаем заказы из стрима

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func handler(w http.ResponseWriter, r *http.Request) {
	var ord order.Order
	if guid := r.URL.Query().Get("guid"); guid != "" {
		or, ok := cache.Cache.Load(guid)
		if ok == true {
			ord = or.(order.Order)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	}
	log.Println(templates.Execute(w, ord))
}
