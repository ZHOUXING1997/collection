package map_collection

import (
	"reflect"
)

// Collection 是集合操作的核心结构，实现了对各种数据类型的集合操作。
// Collection 是针对map类型的集合操作核心结构
type Collection[K comparable, V any] struct {
	value map[K]V      // map数据
	vType reflect.Type // value中值的类型
	kType reflect.Type // key的类型

	// 比较函数，用于值的比较
	valCompareFunc func(V, V) int
	// key的比较函数，用于key的排序等操作
	keyCompareFunc func(K, K) int

	// 缓存排序后的keys,用于稳定的First/Last操作
	sortedKeys []K
}
