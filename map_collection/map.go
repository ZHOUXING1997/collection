package map_collection

import (
	"reflect"
)

// CollectionOption 是用于配置 Collection 的函数式选项
type CollectionOption[K comparable, V any] func(*Collection[K, V])

// WithKeyCompare 设置 key 的比较函数（用于排序）
// compareFunc: 比较函数，返回 -1(小于)、0(等于)、1(大于)
func WithKeyCompare[K comparable, V any](compareFunc func(K, K) int) CollectionOption[K, V] {
	return func(c *Collection[K, V]) {
		c.keyCompareFunc = compareFunc
	}
}

// WithValCompare 设置 value 的比较函数（用于按值排序）
// compareFunc: 比较函数，返回 -1(小于)、0(等于)、1(大于)
func WithValCompare[K comparable, V any](compareFunc func(V, V) int) CollectionOption[K, V] {
	return func(c *Collection[K, V]) {
		c.valCompareFunc = compareFunc
	}
}

// NewCollection 创建并返回一个新的 Collection 实例
// 支持函数式选项模式配置比较函数
func NewCollection[T ~map[K]V, K comparable, V any](values T, opts ...CollectionOption[K, V]) *Collection[K, V] {
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

	// 应用选项
	for _, opt := range opts {
		opt(coll)
	}

	return coll
}
