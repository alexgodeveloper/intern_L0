package main

import (
	"intern_L0/pkg/cache"
	"intern_L0/pkg/database"
	"intern_L0/pkg/generator"
	"intern_L0/pkg/services"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

const dsn_test = "host=localhost user=test password=12345 dbname=l0-test port=5433 sslmode=disable TimeZone=Europe/Moscow"

// Тестируем http
func Test_http(t *testing.T) {
	want := 200
	req := httptest.NewRequest("GET", "http://localhost", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != want {
		t.Errorf("Status code %d not equal %d", resp.StatusCode, want)
	}
}

// Тестируем на ненайденный заказ
func Test_orderNotFound(t *testing.T) {
	want := 404
	wantMsg := "Уточните номер заказа..."
	req := httptest.NewRequest("GET", "http://localhost/?guid=test123", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != want {
		t.Errorf("Status code %d not equal %d", resp.StatusCode, want)
	}

	if !strings.Contains(string(body), wantMsg) {
		t.Errorf("Expected message %v not found on page \n%v ", wantMsg, string(body))
	}
}

// Тестируем подключение, запись БД.
func TestDb(t *testing.T) {
	database.InitDb(dsn_test)
	newOrder := generator.GenerateOrder()

	services.SaveOrder(&newOrder)
	cache.Cache.Store(newOrder.OrderUID, newOrder)

	want := 200
	wantMsg := newOrder.TrackNumber
	req := httptest.NewRequest("GET", "http://localhost/?guid="+newOrder.OrderUID, nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	t.Run("check status code", func(t *testing.T) {
		if resp.StatusCode != want {
			t.Errorf("Status code %d not equal %d", resp.StatusCode, want)
		}
	})

	t.Run("check response body", func(t *testing.T) {
		if !strings.Contains(string(body), wantMsg) {
			t.Errorf("Expected message %v not found on page \n%v ", wantMsg, string(body))
		}
	})
}
