package collection

import (
	"github.com/ZHOUXING1997/collection/map_collection"
	"github.com/ZHOUXING1997/collection/slice_collcection"
)

// NewSliceCollect 创建并返回一个新的 Collection 实例。
//
// 该函数接收一个切片作为输入，并自动初始化适合该类型的比较函数。
// 对于基本类型（如 int, string, float 等），会自动设置默认的比较函数。
// 对于结构体类型，可能需要手动设置比较函数。
//
// 参数:
//   - values: 要初始化的切片
//
// 返回:
//   - *Collection[T]: 新创建的 Collection 实例
func NewSliceCollect[T any](values []T) *slice_collcection.Collection[T] {
	return slice_collcection.NewCollection[T](values)
}

// NewEmptyCollection 返回一个空的Collection
func NewEmptyCollection[T any]() *slice_collcection.Collection[T] {
	return slice_collcection.NewEmptyCollection[T]()
}

// NewMapCollect 创建并返回一个新的 map Collection 实例（默认初始化时排序）。
//
// 该函数接收一个 map 作为输入，并返回一个支持链式操作的 Collection 对象。
// 创建时会对 keys 进行排序，以支持稳定的 First/Last 操作。
//
// 参数:
//   - values: 要初始化的 map
//
// 返回:
//   - *map_collection.Collection[K, V]: 新创建的 map Collection 实例
func NewMapCollect[K comparable, V any](values map[K]V) *map_collection.Collection[K, V] {
	return map_collection.NewCollection[map[K]V, K, V](values)
}

// NewMapCollectUnsorted 创建并返回一个新的 map Collection 实例（不排序）。
//
// 该函数适合不使用 First/Last 方法的场景，可以避免初始化时的排序开销。
//
// 参数:
//   - values: 要初始化的 map
//
// 返回:
//   - *map_collection.Collection[K, V]: 新创建的 map Collection 实例（未排序）
func NewMapCollectUnsorted[K comparable, V any](values map[K]V) *map_collection.Collection[K, V] {
	return map_collection.NewCollectionUnsorted[map[K]V, K, V](values)
}

// NewEmptyMapCollection 返回一个空的 map Collection（默认排序）
func NewEmptyMapCollection[K comparable, V any]() *map_collection.Collection[K, V] {
	return map_collection.NewCollection[map[K]V, K, V](make(map[K]V))
}

// NewEmptyMapCollectionUnsorted 返回一个空的 map Collection（不排序）
func NewEmptyMapCollectionUnsorted[K comparable, V any]() *map_collection.Collection[K, V] {
	return map_collection.NewCollectionUnsorted[map[K]V, K, V](make(map[K]V))
}
