package cache

import (
	"intern_L0/pkg/services"
	"sync"
)

var Cache sync.Map

func SetCacheFromDb(c *sync.Map) {

	ord := services.GetAllOrders()
	for _, r := range ord {
		c.Store(r.OrderUID, r)
	}
}
