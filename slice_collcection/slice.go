package slice_collcection

import (
	"reflect"

	"github.com/ZHOUXING1997/collection/utils"
)

// NewCollection 创建并返回一个新的 Collection 实例。
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
func NewCollection[T any](values []T) *Collection[T] {
	var zero T
	typ := reflect.TypeOf(zero)
	coll := &Collection[T]{value: values, typ: typ}

	coll.compareFunc = utils.NewCompareFunc(typ.Kind())

	return coll
}

// NewEmptyCollection 返回一个空的Collection
func NewEmptyCollection[T any]() *Collection[T] {
	return NewCollection[T](nil)
}
