package main

import "fmt"

func getData() []any {
	return []any{
		1,
		3.14,
		"hello",
		true,
		struct{}{},
	}
}

func getType(v any) string {
	switch v.(type) {

	case string:
		return "string"
	case bool:
		return "boolean"
	case float32, float64:
		return "float"
	case int, int8, int16, int32, int64, uint, uint16, uint8, uint32, uint64:
		return "int"
	default:
		return "unknown"
	}
}

func main() {
	data := getData()

	for _, value := range data {
		fmt.Printf("%v is %v\n", value, getType(value))
	}
}
