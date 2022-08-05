package utils

import (
	"encoding/json"
	"fmt"
)

// Перевод любой структуры в Json
func StructToJson(str any) []byte {
	Json, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	return Json
}

// Перевод любого джейсона в структуру
func JsonToStruct[T any](jsn []byte, strct T) T {
	if err := json.Unmarshal(jsn, &strct); err != nil {
		fmt.Println(err)
	}
	return strct
}
