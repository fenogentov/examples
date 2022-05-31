package testdata

import (
	"fmt"
	"time"
)

func examp(v interface{}) {
	switch v := v.(type) {
	case float64:
		fmt.Println(v)
	// case *types.Document:
	// 	fmt.Println(v)
	// case *types.Array, int8:
	// 	fmt.Println(v)
	case string:
		fmt.Println(v)
	// case types.Binary:
	// 	fmt.Println(v)
	// case types.ObjectID:
	// 	fmt.Println(v)

	case time.Time:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	// case types.NullType:
	// 	fmt.Println(v)
	// case types.Regex:
	// 	fmt.Println(v)
	case int32:
		fmt.Println(v)
	// case types.Timestamp:
	// 	fmt.Println(v)
	case int64:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}
}
