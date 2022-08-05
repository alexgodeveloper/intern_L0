package services

import (
	"intern_L0/pkg/database"
	"intern_L0/pkg/order"
)

// Сохраняет заказ в БД
func SaveOrder(order *order.Order) {
	database.Db.Create(order)
}

// Берет все заказы из базы
func GetAllOrders() []order.Order {
	var ord []order.Order
	database.Db.Preload("Payment").Preload("Delivery").Preload("Items").Find(&ord)
	return ord
}
