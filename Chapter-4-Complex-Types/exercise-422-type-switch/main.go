package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error) {
	switch t := v.(type) {

	case string:
		return fmt.Sprint(t + t), nil

	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32:
		return fmt.Sprint(t * 2), nil

	case float64:
		return fmt.Sprint(t * 2), nil

	case int:
		return fmt.Sprint(t * 2), nil

	case int8:
		return fmt.Sprint(t * 2), nil

	case int16:
		return fmt.Sprint(t * 2), nil

	case int32:
		return fmt.Sprint(t * 2), nil

	case int64:
		return fmt.Sprint(t * 2), nil

	case uint:
		return fmt.Sprint(t * 2), nil

	case uint8:
		return fmt.Sprint(t * 2), nil

	case uint16:
		return fmt.Sprint(t * 2), nil

	case uint32:
		return fmt.Sprint(t * 2), nil

	case uint64:
		return fmt.Sprint(t * 2), nil

	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := doubler("yum")
	fmt.Println("string     yum    ->", res)

	res, _ = doubler(true)
	fmt.Println("bool       true   ->", res)

	res, _ = doubler(float32(3.14))
	fmt.Println("float32    3.14   ->", res)

	res, _ = doubler(float64(3.14))
	fmt.Println("float64    3.14   ->", res)

	res, _ = doubler(int(5))
	fmt.Println("int        5      ->", res)

	res, _ = doubler(int8(10))
	fmt.Println("int8       10     ->", res)

	res, _ = doubler(int16(100))
	fmt.Println("int16      100    ->", res)

	res, _ = doubler(int32(1000))
	fmt.Println("int32      1000   ->", res)

	res, _ = doubler(int64(10000))
	fmt.Println("int64      10000  ->", res)

	res, _ = doubler(uint(5))
	fmt.Println("uint       5      ->", res)

	res, _ = doubler(uint8(10))
	fmt.Println("uint8      10     ->", res)

	res, _ = doubler(uint16(100))
	fmt.Println("uint16     100    ->", res)

	res, _ = doubler(uint32(1000))
	fmt.Println("uint32     1000   ->", res)

	res, _ = doubler(uint64(10000))
	fmt.Println("uint64     10000  ->", res)

	_, err := doubler(struct{}{})
	fmt.Println("unsupported      ->", err)
}
