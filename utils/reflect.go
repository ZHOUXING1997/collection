package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// IsComputableKind 检查类型是否可计算
func IsComputableKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func Any2Float(val interface{}) (float64, error) {
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

// NewCompareFunc 创建比较函数, -1 小于，0 等于，1 大于
func NewCompareFunc(kind reflect.Kind) func(any, any) int {
	switch kind {
	case reflect.Int:
		return func(a, b any) int {
			vala := a.(int)
			valb := b.(int)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Int8:
		return func(a, b any) int {
			vala := a.(int8)
			valb := b.(int8)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Int16:
		return func(a, b any) int {
			vala := a.(int16)
			valb := b.(int16)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Int32:
		return func(a, b any) int {
			vala := a.(int32)
			valb := b.(int32)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Int64:
		return func(a, b any) int {
			vala := a.(int64)
			valb := b.(int64)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Uint:
		return func(a, b any) int {
			vala := a.(uint)
			valb := b.(uint)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Uint8:
		return func(a, b any) int {
			vala := a.(uint8)
			valb := b.(uint8)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Uint16:
		return func(a, b any) int {
			vala := a.(uint16)
			valb := b.(uint16)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Uint32:
		return func(a, b any) int {
			vala := a.(uint32)
			valb := b.(uint32)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Uint64:
		return func(a, b any) int {
			vala := a.(uint64)
			valb := b.(uint64)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Float32:
		return func(a, b any) int {
			vala := a.(float32)
			valb := b.(float32)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	case reflect.Float64:
		return func(a, b any) int {
			vala := a.(float64)
			valb := b.(float64)
			if vala > valb {
				return 1
			} else if vala < valb {
				return -1
			}
			return 0
		}
	default:
		return nil
	}
}
