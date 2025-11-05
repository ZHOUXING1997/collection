package slice_collcection

import (
	"reflect"
)

// Collection 是集合操作的核心结构，实现了对各种数据类型的集合操作。
//
// Collection 使用泛型参数 T 来支持不同类型的数据，包括基本类型和结构体类型。
// 它提供了丰富的方法来操作和转换这些数据，如过滤、映射、排序、查找、聚合等。
//
// 使用示例：
//
//	intColl := NewCollection([]int{1, 2, 3, 4, 5})
//	filtered := intColl.Filter(func(item int, key int) bool {
//	    return item > 2
//	})
type Collection[T any] struct {
	value []T // 数组

	// err error        // 错误信息
	typ reflect.Type // collection 中每个元素的类型，在new的时候就定义了

	compareFunc func(any, any) int // 比较函数，在new的时候定义了，也可以通过 SetCompare 方法进行设置
}
