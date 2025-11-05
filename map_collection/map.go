package map_collection

import (
	"reflect"
	"slices"
	"sort"
)

// extractField 从值中提取指定字段
func extractField(v any, fieldName string) any {
	val := reflect.ValueOf(v)

	// 处理指针类型
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	// 确保是结构体
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 获取字段值
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Interface()
}

// NewCollection 创建并返回一个新的 Collection 实例（默认初始化时排序）。
func NewCollection[T ~map[K]V, K comparable, V any](values T) *Collection[K, V] {
	var vzero V
	vType := reflect.TypeOf(vzero)
	var kZero K
	kType := reflect.TypeOf(kZero)

	coll := &Collection[K, V]{
		value:          values,
		vType:          vType,
		kType:          kType,
		valCompareFunc: nil,
		keyCompareFunc: nil,
		sortedKeys:     Keys(values),
	}

	return coll
}
