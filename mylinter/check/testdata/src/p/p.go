package main

import (
	"fmt"
	"time"
)

func switchOK(v interface{}) {
	switch v := v.(type) {
	case *Document:
		fmt.Println(v)
	case Array:
		fmt.Println(v)
	case float64:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	case Binary:
		fmt.Println(v)
	case ObjectID:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	case time.Time:
		fmt.Println(v)
	case NullType:
		fmt.Println(v)
	case Regex:
		fmt.Println(v)
	case int32:
		fmt.Println(v)
	case Timestamp:
		fmt.Println(v)
	case int64:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}
}

func caseOK(v interface{}) {
	switch v := v.(type) {
	case *Document:
		fmt.Println(v)
	case *Array:
		fmt.Println(v)
	case float64, int32, int64:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	case Binary:
		fmt.Println(v)
	case ObjectID:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	case time.Time:
		fmt.Println(v)
	case NullType:
		fmt.Println(v)
	case Regex:
		fmt.Println(v)
	case Timestamp:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}
}

// func switchWrong(v interface{}) {
// 	switch v := v.(type) { // want "non-observance of the preferred order of types"
// 	case float64:
// 		fmt.Println(v)
// 	case *Array, int8:
// 		fmt.Println(v)
// 	case *Document:
// 		fmt.Println(v)
// 	case string:
// 		fmt.Println(v)
// 	case Binary:
// 		fmt.Println(v)
// 	case ObjectID:
// 		fmt.Println(v)
// 	case time.Time:
// 		fmt.Println(v)
// 	case bool:
// 		fmt.Println(v)
// 	case NullType:
// 		fmt.Println(v)
// 	case Regex:
// 		fmt.Println(v)
// 	case int32:
// 		fmt.Println(v)
// 	case Timestamp:
// 		fmt.Println(v)
// 	case int64:
// 		fmt.Println(v)
// 	default:
// 		fmt.Println(v)
// 	}
// }

// func caseWrong(v interface{}) {
// 	switch v := v.(type) { // want "non-observance of the preferred order of types"
// 	case *Document:
// 		fmt.Println(v)
// 	case *Array:
// 		fmt.Println(v)
// 	case float64, int64, int32:
// 		fmt.Println(v)
// 	case string:
// 		fmt.Println(v)
// 	case Binary:
// 		fmt.Println(v)
// 	case ObjectID:
// 		fmt.Println(v)
// 	case bool:
// 		fmt.Println(v)
// 	case time.Time:
// 		fmt.Println(v)
// 	case NullType:
// 		fmt.Println(v)
// 	case Regex:
// 		fmt.Println(v)
// 	case Timestamp:
// 		fmt.Println(v)
// 	default:
// 		fmt.Println(v)
// 	}
// }

// func severalWrong(v interface{}) {
// 	switch v := v.(type) { // want "non-observance of the preferred order of types"
// 	case *Document:
// 		fmt.Println(v)
// 	case *Array:
// 		fmt.Println(v)
// 	case float64, int32, int64:
// 		fmt.Println(v)
// 	case bool:
// 		fmt.Println(v)
// 	case string:
// 		fmt.Println(v)
// 	case Binary:
// 		fmt.Println(v)
// 	case ObjectID:
// 		fmt.Println(v)
// 	case time.Time:
// 		fmt.Println(v)
// 	case NullType:
// 		fmt.Println(v)
// 	case Regex:
// 		fmt.Println(v)
// 	case Timestamp:
// 		fmt.Println(v)
// 	default:
// 		fmt.Println(v)
// 	}
// }
