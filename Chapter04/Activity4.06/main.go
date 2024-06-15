package main

import "fmt"

func typeChecker(v any) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	test_types := []any{2, 3.14, "hello", true, struct{ name string }{"boy"}}
	for _, t := range test_types {
		res := typeChecker(t)
		fmt.Printf("%v is %v\n", t, res)
	}
}
