package collection

import (
	"fmt"
	"reflect"
	"strconv"
)

// isComputableKind 检查类型是否可计算
func isComputableKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func any2Float(val interface{}) (float64, error) {
	switch val.(type) {
	case int:
		return float64(val.(int)), nil
	case int8:
		return float64(val.(int8)), nil
	case int16:
		return float64(val.(int16)), nil
	case int32:
		return float64(val.(int32)), nil
	case int64:
		return float64(val.(int64)), nil
	case uint:
		return float64(val.(uint)), nil
	case uint8:
		return float64(val.(uint8)), nil
	case uint16:
		return float64(val.(uint16)), nil
	case uint32:
		return float64(val.(uint32)), nil
	case uint64:
		return float64(val.(uint64)), nil
	case float32:
		return float64(val.(float32)), nil
	case float64:
		return val.(float64), nil
	case string:
		number, err := strconv.ParseFloat(val.(string), 64)
		return number, err
	default:
		return 0, fmt.Errorf("invalid type %T", val)
	}
}
