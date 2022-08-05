package generator

import (
	"github.com/google/uuid"
	"intern_L0/pkg/order"
	"time"
)

// Создает заказ
func GenerateOrder() order.Order {

	order := order.Order{
		OrderUID:    uuid.NewString(),
		TrackNumber: uuid.NewString(),
		Entry:       uuid.NewString(),
		Delivery: order.Delivery{
			OrderID: 0,
			Name:    "Test",
			Phone:   "123123",
			Zip:     "",
			City:    "",
			Address: "",
			Region:  "",
			Email:   "",
		},
		Payment: order.Payment{
			OrderID:      0,
			Transaction:  "321",
			RequestID:    "",
			Currency:     "RUB",
			Provider:     "",
			Amount:       0,
			PaymentDt:    0,
			Bank:         "",
			DeliveryCost: 0,
			GoodsTotal:   0,
			CustomFee:    0,
		},
		Items:             []order.Items{{ChrtID: 1, TrackNumber: "asd", Price: 1, Rid: "asd", Name: "asd", Sale: 12, Size: "123", TotalPrice: 13, NmID: 1, Brand: "a", Status: 1}, {ChrtID: 2, TrackNumber: "asd", Price: 1, Rid: "asd", Name: "sdf", Sale: 12, Size: "123", TotalPrice: 13, NmID: 1, Brand: "a", Status: 1}},
		Locale:            "1",
		InternalSignature: "1",
		CustomerID:        "1",
		DeliveryService:   "1",
		Shardkey:          "1",
		SmID:              0,
		DateCreated:       time.Time{},
		OofShard:          "1",
	}

	return order
}
