package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// NoComputableError 不可计算
var NoComputableError = errors.New("collection is not computable")
var ElementNoComputableError = errors.New("element can not be computable")

// KeyUnComparableError 不可比较
var KeyUnComparableError = errors.New("key has unComparable type")
var NoComparableError = errors.New("collection is not comparable")

// 是结构体或者指针
func (c *Collection[T]) isStructOrPointer() bool {
	switch c.typ.Kind() {
	case reflect.Struct:
		return true
	case reflect.Pointer:
		if c.typ.Elem().Kind() == reflect.Struct {
			return true
		}
	default:
		return false
	}

	return false
}

// 是可比较的
func (c *Collection[T]) isComparable() bool {
	return c.compareFunc != nil
}

// 是可以进行计算的
func (c *Collection[T]) isComputable() bool {
	return isComputableKind(c.typ.Kind())
}

// -1：小于，0：等于，1：大于
func (c *Collection[T]) newCompare(kind reflect.Kind) func(any, any) int {
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
	// case reflect.String: // 字符串排序有问题，这里不再自动注册，如果需要使用，请使用 SetCompare 进行注册
	// 	return func(a, b any) int {
	// 		vala := a.(string)
	// 		valb := b.(string)
	// 		return strings.Compare(vala, valb)
	// 	}
	default:
		return nil
	}
}

// Collection 主体
type Collection[T any] struct {
	value []T // 数组

	// err error        // 错误信息
	typ reflect.Type // collection 中每个元素的类型，在new的时候就定义了

	compareFunc func(any, any) int // 比较函数，在new的时候定义了，也可以通过 SetCompare 方法进行设置
}

// NewCollection 初始化一个compare
func NewCollection[T any](values []T) *Collection[T] {
	var zero T
	typ := reflect.TypeOf(zero)
	coll := &Collection[T]{value: values, typ: typ}

	coll.compareFunc = coll.newCompare(typ.Kind())

	return coll
}

// NewEmptyCollection 返回一个空的Collection
func NewEmptyCollection[T any]() *Collection[T] {
	return NewCollection[T](nil)
}

// // Err 返回Collection的错误信息
// func (c *Collection[T]) Err() error {
// 	return c.err
// }
//
// // SetErr 设置Collection的错误信息
// func (c *Collection[T]) SetErr(err error) *Collection[T] {
// 	c.err = err
// 	return c
// }

// SetCompare 设置比较函数
// -1：小于，0：等于，1：大于
func (c *Collection[T]) SetCompare(compareFunc func(a any, b any) int) *Collection[T] {
	c.compareFunc = compareFunc
	return c
}

// Copy 复制一个新的Collection
func (c *Collection[T]) Copy() *Collection[T] {
	// coll := NewCollection[T](c.value).SetErr(c.err)
	coll := NewCollection[T](c.value)
	return coll
}

// IsEmpty 判断是否为空
func (c *Collection[T]) IsEmpty() bool {
	return len(c.value) == 0
}

// IsNotEmpty 判断是否不为空
func (c *Collection[T]) IsNotEmpty() bool {
	return len(c.value) != 0
}

// Append 添加元素
func (c *Collection[T]) Append(item T) *Collection[T] {
	c.value = append(c.value, item)
	return c
}

// Remove 删除元素
func (c *Collection[T]) Remove(index int) *Collection[T] {
	c.value = append(c.value[:index], c.value[index+1:]...)
	return c
}

// Insert 插入元素
func (c *Collection[T]) Insert(index int, item T) *Collection[T] {
	arr := make([]T, 0, len(c.value)+1)
	arr = append(arr, c.value[:index]...)
	arr = append(arr, item)
	arr = append(arr, c.value[index:]...)
	return NewCollection[T](arr)
}

// Search 查找元素
func (c *Collection[T]) Search(item T) (int, error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("element can not be comparable"))
		// return 0
		return -1, ElementNoComputableError
	}

	for i, v := range c.value {
		if c.compareFunc(v, item) == 0 {
			return i, nil
		}
	}

	return -1, fmt.Errorf("not found")
}

// Unique 去重
func (c *Collection[T]) Unique() (*Collection[T], error) {
	// if c.err != nil {
	// 	return c
	// }
	// if !c.isComparable() {
	// 	c.SetErr(errors.New("element can not be comparable"))
	// 	return c
	// }

	if !c.isComparable() {
		return c, ElementNoComputableError
	}

	// 过滤数组中重复的元素，仅对基础Collection生效
	res := make([]T, 0, len(c.value))
	inArr := func(item T, arr []T) bool {
		for _, val := range arr {
			if c.compareFunc(item, val) == 0 {
				return true
			}
		}
		return false
	}
	for i, v := range c.value {
		if !inArr(v, res) {
			res = append(res, c.value[i])
		}
	}

	return NewCollection[T](res), nil
}

// Filter 过滤
func (c *Collection[T]) Filter(f func(item T, key int) bool) *Collection[T] {
	// if c.err != nil {
	// 	return c
	// }

	// 按照某个方法进行过滤, 保留符合的
	res := make([]T, 0, len(c.value))
	for i, v := range c.value {
		if f(v, i) {
			res = append(res, v)
		}
	}

	return NewCollection[T](res)
}

// Reject 过滤
func (c *Collection[T]) Reject(f func(item T, key int) bool) *Collection[T] {
	// if c.err != nil {
	// 	return c
	// }

	res := make([]T, 0, len(c.value))
	for i, v := range c.value {
		if !f(v, i) {
			res = append(res, v)
		}
	}

	return NewCollection[T](res)
}

// First 获取第一个元素
func (c *Collection[T]) First() T {
	var zero T
	if c.IsEmpty() {
		return zero
	}

	return c.value[0]
}

// Last 获取最后一个元素
func (c *Collection[T]) Last() T {
	var zero T
	if c.IsEmpty() {
		return zero
	}

	return c.value[len(c.value)-1]
}

// Slice 获取数组片段
func (c *Collection[T]) Slice(params ...int) (*Collection[T], error) {
	if len(params) == 0 {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("invalid params"))
		return NewCollection[T](nil), fmt.Errorf("invalid params")
	}
	start := params[0]
	if start < 0 || start >= len(c.value) {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("invalid start index"))
		return NewCollection[T](nil), fmt.Errorf("invalid start index")
	}
	if len(params) == 1 {
		// return NewCollection(c.value[start:]).SetErr(nil)
		return NewCollection(c.value[start:]), nil
	}
	end := params[1]
	if end < 0 || end > len(c.value) {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("invalid end index"))
		return NewCollection[T](nil), fmt.Errorf("invalid end index")
	}
	if start > end {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("start index should be less than end index"))
		return NewCollection[T](nil), fmt.Errorf("start index should be less than end index")
	}

	return NewCollection(c.value[start:end]), nil
}

// Index 获取某个下标
func (c *Collection[T]) Index(i int) T {

	// 获取某个下标，对所有Collection生效
	var zero T
	if i < 0 || i >= len(c.value) {
		return zero
	}

	return c.value[i]
}

// SetIndex 设置某个下标
func (c *Collection[T]) SetIndex(i int, val T) *Collection[T] {
	// if c.err != nil {
	// 	return c
	// }

	// 设置数组的下标为某个值
	if i < 0 || i >= len(c.value) {
		return c
	}
	c.value[i] = val
	return c
}

// Count 获取数组长度
func (c *Collection[T]) Count() int {
	// 获取数组长度，对所有Collection生效
	return len(c.value)
}

// Merge 合并数组
func (c *Collection[T]) Merge(arr *Collection[T]) *Collection[T] {
	// 将两个数组进行合并
	if arr == nil {
		return c
	}

	// if c.err != nil {
	// 	return c
	// }
	//
	// if arr.err != nil {
	// 	return c.SetErr(arr.err)
	// }

	res := c.Copy()
	for i := 0; i < arr.Count(); i++ {
		res.value = append(res.value, arr.Index(i))
	}

	return res
}

// Each 遍历
// 当 f 返回 false 时，终止
func (c *Collection[T]) Each(f func(item T, key int) bool) {
	for i, v := range c.value {
		if !f(v, i) {
			return
		}
	}
}

// ForEach 遍历，能保证每一个元素都被遍历到，内部决定是否跳过
func (c *Collection[T]) Foreach(f func(item T, key int)) {
	for i, v := range c.value {
		f(v, i)
	}
}

// Map 映射
func (c *Collection[T]) Map(f func(item T, key int) T) *Collection[T] {
	res := make([]T, 0, len(c.value))
	for i, v := range c.value {
		res = append(res, f(v, i))
	}

	return NewCollection[T](res)
}

// Map 映射并过滤
// 映射后，可以过滤掉一些元素
func (c *Collection[T]) MapFilter(f func(item T, key int) (T, bool)) *Collection[T] {
	res := make([]T, 0, len(c.value))
	for i, v := range c.value {
		r, ok := f(v, i)
		if ok {
			res = append(res, r)
		}
	}

	return NewCollection[T](res)
}

// Reduce 求和
func (c *Collection[T]) Reduce(f func(carry T, item T) T) T {
	var zero T
	if c.IsEmpty() {
		return zero
	}
	res := c.value[0]
	for i := 1; i < len(c.value); i++ {
		res = f(res, c.value[i])
	}

	return res
}

// Every 判断是否所有元素都满足条件
func (c *Collection[T]) Every(f func(item T, key int) bool) bool {
	for i, v := range c.value {
		if !f(v, i) {
			return false
		}
	}

	return true
}

// ForPage 分页
func (c *Collection[T]) ForPage(page int, perPage int) (*Collection[T], error) {
	if page <= 0 || perPage <= 0 {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("invalid page or perPage"))
		return NewCollection[T](nil), fmt.Errorf("invalid page or perPage")
	}
	start := (page - 1) * perPage
	end := start + perPage
	if start >= len(c.value) {
		return NewCollection[T](nil), nil
	}
	if end > len(c.value) {
		end = len(c.value)
	}

	return &Collection[T]{value: c.value[start:end]}, nil
}

// Nth 每隔n个取一个
func (c *Collection[T]) Nth(n int, offset int) (*Collection[T], error) {
	if n <= 0 {
		// return NewCollection[T](nil).SetErr(fmt.Errorf("invalid n"))

		return NewCollection[T](nil), fmt.Errorf("invalid n")
	}
	res := make([]T, 0, len(c.value)/n+1)
	for i := offset; i < len(c.value); i += n {
		res = append(res, c.value[i])
	}

	return NewCollection[T](res), nil
}

// Pad 填充
func (c *Collection[T]) Pad(count int, def T) *Collection[T] {
	if count <= len(c.value) {
		return c.Copy()
	}
	res := make([]T, count)
	for i := 0; i < len(c.value); i++ {
		res[i] = c.value[i]
	}
	for i := len(c.value); i < count; i++ {
		res[i] = def
	}

	return NewCollection[T](res)
}

// Pop 弹出最后一个元素
func (c *Collection[T]) Pop() T {
	var zero T
	if c.IsEmpty() {
		return zero
	}
	res := c.value[len(c.value)-1]
	c.value = c.value[:len(c.value)-1]
	return res
}

// Push 添加元素
func (c *Collection[T]) Push(item T) *Collection[T] {
	c.value = append(c.value, item)
	return c
}

// Prepend 添加元素到头部
func (c *Collection[T]) Prepend(item T) *Collection[T] {
	res := make([]T, 0, len(c.value)+1)
	res = append(res, item)
	res = append(res, c.value...)
	return NewCollection[T](res)
}

// Random 随机取一个元素
func (c *Collection[T]) Random() T {
	var zero T
	if c.IsEmpty() {
		return zero
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return c.value[random.Intn(len(c.value))]
}

// Reverse 反转
func (c *Collection[T]) Reverse() *Collection[T] {
	res := make([]T, 0, len(c.value))
	for i := len(c.value) - 1; i >= 0; i-- {
		res = append(res, c.value[i])
	}

	return NewCollection(res)
}

// Shuffle 随机排序
func (c *Collection[T]) Shuffle() *Collection[T] {
	res := make([]T, len(c.value))
	perm := rand.Perm(len(c.value))
	for i, v := range perm {
		res[v] = c.value[i]
	}

	return NewCollection(res)
}

// GroupBy 分组
func (c *Collection[T]) GroupBy(f func(T, int) interface{}) map[interface{}]*Collection[T] {
	res := make(map[interface{}]*Collection[T])
	for i, v := range c.value {
		key := f(v, i)
		if _, ok := res[key]; !ok {
			res[key] = NewCollection[T](nil)
		}
		res[key].Push(v)
	}

	return res
}

// Split 按照size个数进行分组
func (c *Collection[T]) Split(size int) []*Collection[T] {
	if size <= 0 {
		return []*Collection[T]{}
	}
	var res []*Collection[T]
	for i := 0; i < len(c.value); i += size {
		end := i + size
		if end > len(c.value) {
			end = len(c.value)
		}
		res = append(res, NewCollection(c.value[i:end]))
	}

	return res
}

// DD 打印出当前数组结构
func (c *Collection[T]) DD() {
	ret := fmt.Sprintf("Collection(%d, %s):{\n", c.Count(), c.typ.String())
	for k, v := range c.value {
		ret = ret + fmt.Sprintf("\t%d:\t%v\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

// PluckString 按照某个字段进行筛选
func (c *Collection[T]) PluckString(key string) (*Collection[string], error) {
	res := make([]string, 0, len(c.value))
	if c.typ.Kind() != reflect.Struct && c.typ.Kind() != reflect.Pointer {
		// c.SetErr(errors.New("invalid collection"))
		// return nil
		return nil, errors.New("invalid collection")
	}

	for _, v := range c.value {
		var val reflect.Value
		if c.typ.Kind() == reflect.Ptr {
			val = reflect.ValueOf(v).Elem().FieldByName(key)
		} else if c.typ.Kind() == reflect.Struct {
			val = reflect.ValueOf(v).FieldByName(key)
		}
		kind := val.Type().Kind()
		if kind != reflect.String {
			// c.SetErr(errors.New("invalid type"))
			// return nil
			return nil, errors.New("invalid type")
		}

		res = append(res, val.String())
	}

	return NewCollection(res), nil
}

// PluckInt64 按照某个字段进行筛选
func (c *Collection[T]) PluckInt64(key string) (*Collection[int64], error) {
	res := make([]int64, 0, len(c.value))
	for _, v := range c.value {
		var val reflect.Value
		if c.typ.Kind() == reflect.Ptr {
			val = reflect.ValueOf(v).Elem().FieldByName(key)
		} else if c.typ.Kind() == reflect.Struct {
			val = reflect.ValueOf(v).FieldByName(key)
		}
		if !val.CanInt() {
			// c.SetErr(errors.New("invalid type"))
			// return nil

			return nil, errors.New("invalid type")
		}

		res = append(res, val.Int())
	}

	return NewCollection(res), nil
}

// PluckFloat64 按照某个字段进行筛选
func (c *Collection[T]) PluckFloat64(key string) (*Collection[float64], error) {
	res := make([]float64, 0, len(c.value))
	for _, v := range c.value {
		var val reflect.Value
		if c.typ.Kind() == reflect.Ptr {
			val = reflect.ValueOf(v).Elem().FieldByName(key)
		} else if c.typ.Kind() == reflect.Struct {
			val = reflect.ValueOf(v).FieldByName(key)
		}
		if !val.CanFloat() {
			// c.SetErr(errors.New("invalid type"))
			// return nil

			return nil, errors.New("invalid type")
		}

		res = append(res, val.Float())
	}

	return NewCollection(res), nil
}

// PluckUint64 按照某个字段进行筛选
func (c *Collection[T]) PluckUint64(key string) (*Collection[uint64], error) {
	res := make([]uint64, 0, len(c.value))
	for _, v := range c.value {
		var val reflect.Value
		if c.typ.Kind() == reflect.Ptr {
			val = reflect.ValueOf(v).Elem().FieldByName(key)
		} else if c.typ.Kind() == reflect.Struct {
			val = reflect.ValueOf(v).FieldByName(key)
		}
		if !val.CanUint() {
			// c.SetErr(errors.New("invalid type"))
			// return nil

			return nil, errors.New("invalid type")
		}

		res = append(res, val.Uint())
	}

	return NewCollection(res), nil
}

// PluckBool 按照某个字段进行筛选
func (c *Collection[T]) PluckBool(key string) (*Collection[bool], error) {
	res := make([]bool, 0, len(c.value))
	for _, v := range c.value {
		var val reflect.Value
		if c.typ.Kind() == reflect.Ptr {
			val = reflect.ValueOf(v).Elem().FieldByName(key)
		} else if c.typ.Kind() == reflect.Struct {
			val = reflect.ValueOf(v).FieldByName(key)
		}
		if val.Kind() != reflect.Bool {
			// c.SetErr(errors.New("invalid type"))
			// return nil

			return nil, errors.New("invalid type")
		}

		res = append(res, val.Bool())
	}

	return NewCollection(res), nil
}

// SortBy 按照某个字段进行排序
func (c *Collection[T]) SortBy(key string) (*Collection[T], error) {
	if c.IsEmpty() || !c.isStructOrPointer() {
		return c, errors.New("collection `T` must be a struct or pointer to struct")
	}

	// 预先检查字段类型
	baseType := reflect.TypeOf(c.value[0])
	isPointer := baseType.Kind() == reflect.Pointer

	// 获取实际的结构体类型
	if isPointer {
		baseType = baseType.Elem()
	}

	// 检查字段是否存在和类型是否可比较
	field, exists := baseType.FieldByName(key)
	if !exists {
		return c, fmt.Errorf("field %s does not exist", key)
	}

	// 使用闭包避免重复反射操作
	lessFunc := c.newCompare(field.Type.Kind())
	if lessFunc == nil {
		return c, KeyUnComparableError
	}

	sort.Slice(c.value, func(i, j int) bool {
		var val1, val2 reflect.Value
		// 处理指针类型
		if isPointer {
			val1 = reflect.ValueOf(c.value[i]).Elem().FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).Elem().FieldByName(key)
		} else {
			val1 = reflect.ValueOf(c.value[i]).FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).FieldByName(key)
		}
		r := lessFunc(val1.Interface(), val2.Interface())

		return r < 0
	})

	return c, nil
}

// SortBy 按照某个字段转换为float进行排序，且字段必须为 string、int、float
// 如果转换为float失败后，会视为没满足条件
func (c *Collection[T]) SortFloatBy(key string) (*Collection[T], error) {
	if c.IsEmpty() || !c.isStructOrPointer() {
		return c, errors.New("collection `T` must be a struct or pointer to struct")
	}

	// 预先检查字段类型
	baseType := reflect.TypeOf(c.value[0])
	isPointer := baseType.Kind() == reflect.Pointer

	// 获取实际的结构体类型
	if isPointer {
		baseType = baseType.Elem()
	}

	// 检查字段是否存在和类型是否可比较
	field, exists := baseType.FieldByName(key)
	if !exists {
		return c, fmt.Errorf("field %s does not exist", key)
	}

	if !isComputableKind(field.Type.Kind()) && field.Type.Kind() != reflect.String {
		return c, fmt.Errorf("field %s is not computable", key)
	}

	// 使用闭包避免重复反射操作
	lessFunc := c.newCompare(reflect.Float64)
	if lessFunc == nil {
		return c, KeyUnComparableError
	}

	var err error
	sort.Slice(c.value, func(i, j int) bool {
		var val1, val2 reflect.Value
		// 处理指针类型
		if isPointer {
			val1 = reflect.ValueOf(c.value[i]).Elem().FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).Elem().FieldByName(key)
		} else {
			val1 = reflect.ValueOf(c.value[i]).FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).FieldByName(key)
		}

		var v1f, v2f float64
		v1f, err = any2Float(val1.Interface())
		if err != nil {
			return false
		}

		v2f, err = any2Float(val2.Interface())
		if err != nil {
			return false
		}

		return lessFunc(v1f, v2f) < 0
	})

	return c, nil
}

// SortBy 按照某个字段进行排序
func (c *Collection[T]) SortByFunc(fn func(v1, v2 T) bool) (*Collection[T], error) {
	sort.Slice(c.value, func(i, j int) bool {
		return fn(c.value[i], c.value[j])
	})

	return c, nil
}

// SortByDesc 按照某个字段进行排序,倒序
func (c *Collection[T]) SortByDesc(key string) (*Collection[T], error) {
	if c.IsEmpty() || !c.isStructOrPointer() {
		return c, errors.New("collection `T` must be a struct or pointer to struct")
	}

	// 预先检查字段类型
	baseType := reflect.TypeOf(c.value[0])
	isPointer := baseType.Kind() == reflect.Pointer

	// 获取实际的结构体类型
	if isPointer {
		baseType = baseType.Elem()
	}

	// 检查字段是否存在和类型是否可比较
	field, exists := baseType.FieldByName(key)
	if !exists {
		return c, fmt.Errorf("field %s does not exist", key)
	}

	// 使用闭包避免重复反射操作
	lessFunc := c.newCompare(field.Type.Kind())
	if lessFunc == nil {
		return c, KeyUnComparableError
	}

	sort.Slice(c.value, func(i, j int) bool {
		var val1, val2 reflect.Value
		// 处理指针类型
		if isPointer {
			val1 = reflect.ValueOf(c.value[i]).Elem().FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).Elem().FieldByName(key)
		} else {
			val1 = reflect.ValueOf(c.value[i]).FieldByName(key)
			val2 = reflect.ValueOf(c.value[j]).FieldByName(key)
		}
		r := lessFunc(val1.Interface(), val2.Interface())

		return r > 0
	})

	return c, nil
}

// KeyByStrField 根据某个字段为key，返回一个map,要求key对应的field是string
func (c *Collection[T]) KeyByStrField(key string) (map[string]T, error) {
	res := make(map[string]T)
	for _, v := range c.value {
		val := reflect.ValueOf(v).FieldByName(key)
		if val.IsValid() && val.CanInterface() {
			if str, ok := val.Interface().(string); ok {
				res[str] = v
			} else {
				return nil, fmt.Errorf("key is not string")
			}
		}
	}

	return res, nil
}

// Max 数组中最大的元素，仅对基础Collection生效, 可以传递一个比较函数
func (c *Collection[T]) Max() (T, error) {
	var zero T

	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return zero

		return zero, NoComparableError
	}

	if c.IsEmpty() {
		return zero, nil
	}

	maxElem := c.value[0]
	for _, v := range c.value {
		if c.compareFunc(v, maxElem) > 0 {
			maxElem = v
		}
	}

	return maxElem, nil
}

// Min 数组中最小的元素，仅对基础Collection生效
func (c *Collection[T]) Min() (T, error) {
	var zero T

	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return zero

		return zero, NoComparableError
	}

	if c.IsEmpty() {
		return zero, nil
	}

	minElem := c.value[0]
	for _, v := range c.value {
		if c.compareFunc(v, minElem) < 0 {
			minElem = v
		}
	}

	return minElem, nil
}

// Contains 判断是否包含某个元素，（并不进行定位），对基础Collection生效
func (c *Collection[T]) Contains(obj T) (bool, error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return false

		return false, NoComparableError
	}

	if c.IsEmpty() {
		return false, nil
	}

	for _, v := range c.value {
		if c.compareFunc(v, obj) == 0 {
			return true, nil
		}
	}

	return false, nil
}

// ContainsCount 判断包含某个元素的个数，返回0代表没有找到，返回正整数代表个数。必须设置compare函数
func (c *Collection[T]) ContainsCount(obj T) (int, error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return 0

		return 0, NoComparableError
	}

	count := 0
	for _, v := range c.value {
		if c.compareFunc(v, obj) == 0 {
			count++
		}
	}

	return count, nil
}

// Diff 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
func (c *Collection[T]) Diff(arr *Collection[T]) (*Collection[T], error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		return c, NoComparableError
	}

	res := NewCollection([]T{})
	for _, v := range c.value {
		if ok, err := arr.Contains(v); !ok && err == nil {
			res.Append(v)
		} else if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (c *Collection[T]) Sort() (*Collection[T], error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return nil

		return nil, NoComparableError
	}

	sort.Slice(c.value, func(i, j int) bool {
		return c.compareFunc(c.value[i], c.value[j]) < 0
	})
	return c, nil
}

// SortDesc 进行排序，倒序
func (c *Collection[T]) SortDesc() (*Collection[T], error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return nil

		return nil, NoComparableError
	}

	sort.Slice(c.value, func(i, j int) bool {
		return c.compareFunc(c.value[i], c.value[j]) > 0
	})

	return c, nil
}

// Join 进行拼接
func (c *Collection[T]) Join(split string, format ...func(item interface{}) string) string {
	var res string
	for i, v := range c.value {
		if len(format) > 0 {
			res += format[0](v)
		} else {
			res += fmt.Sprintf("%v", v)
		}
		if i != len(c.value)-1 {
			res += split
		}
	}

	return res
}

// Union 两个集合的并集
func (c *Collection[T]) Union(arr *Collection[T]) (*Collection[T], error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return nil

		return nil, NoComparableError
	}

	res := c.Copy()
	for _, v := range arr.value {
		if ok, err := c.Contains(v); !ok && err == nil {
			res.Append(v)
		} else if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// Intersect 两个集合的交集
func (c *Collection[T]) Intersect(arr *Collection[T]) (*Collection[T], error) {
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		// return nil

		return nil, NoComparableError
	}

	res := NewCollection([]T{})
	for _, v := range c.value {
		if ok, err := arr.Contains(v); ok && err == nil {
			res.Append(v)
		} else if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// Avg 获取平均值
func (c *Collection[T]) Avg() (float64, error) {
	if !c.isComputable() {
		// c.SetErr(errors.New("collection is not floatable"))
		// return 0.0

		return 0.0, NoComputableError
	}

	if c.IsEmpty() {
		return 0.0, nil
	}
	sum, err := c.Sum()
	if err != nil {
		return 0.0, err
	}

	return sum / float64(len(c.value)), nil
}

// Median 获取中位值。
// 中位数（Median）又称中值，统计学中的专有名词，是按顺序排列的一组数据中居于中间位置的数，代表一个样本、种群或概率分布中的一个数值，其可将数值集合划分为相等的上下两部分。
// 对于有限的数集，可以通过把所有观察值高低排序后找出正中间的一个作为中位数。如果观察值有偶数个，通常取最中间的两个数值的平均数作为中位数。
func (c *Collection[T]) Median() (float64, error) {
	if !c.isComputable() {
		// c.SetErr(errors.New("collection is not floatable"))

		return 0.0, NoComputableError
	}

	coll, err := c.Sort()
	if err != nil {
		return 0.0, err
	}
	newColl := NewCollection([]T{})
	if len(coll.value)%2 == 0 {
		newColl.Append(coll.Index(len(coll.value)/2 - 1)).Append(coll.Index(len(coll.value) / 2))
		return newColl.Avg()
	}
	newColl.Append(coll.Index(len(coll.value) / 2))
	return newColl.Avg()
}

// 记录每个元素出现个数的结构，只有Mode用
type tCount struct {
	item   any // 元素
	count  int // 出现的次数
	cindex int // 在原来collection中的index
}

// Mode 获取Mode值，众数，一组数据中出现最多的
func (c *Collection[T]) Mode() (T, error) {
	var zero T
	if !c.isComparable() {
		// c.SetErr(errors.New("collection is not comparable"))
		return zero, NoComparableError
	}

	if c.IsEmpty() {
		return zero, nil
	}

	summary := make([]tCount, 0, c.Count())

	// 查找index的地址
	indexSummary := func(item any, summary []tCount) int {
		for i, val := range summary {
			if c.compareFunc(val.item, item) == 0 {
				return i
			}
		}
		return -1
	}

	for i, item := range c.value {
		index := indexSummary(item, summary)
		if index == -1 {
			summary = append(summary, tCount{
				item:   item,
				count:  1,
				cindex: i,
			})
		} else {
			summary[index].count++
		}
	}

	var maxCount int
	var maxIndex int
	for _, tcount := range summary {
		if tcount.count > maxCount {
			maxCount = tcount.count
			maxIndex = tcount.cindex
		}
	}

	return c.value[maxIndex], nil
}

// Sum 获取sum值
func (c *Collection[T]) Sum() (float64, error) {
	if c.IsEmpty() {
		return 0.0, nil
	}
	if !c.isComputable() {
		// c.SetErr(errors.New("collection is not floatable"))
		return 0.0, NoComputableError
	}

	sum := float64(0)
	for _, item := range c.value {
		switch reflect.ValueOf(item).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sum += float64(reflect.ValueOf(item).Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			sum += float64(reflect.ValueOf(item).Uint())
		case reflect.Float32, reflect.Float64:
			sum += reflect.ValueOf(item).Float()
		default:
			// c.SetErr(errors.New("invalid type"))

			return 0.0, fmt.Errorf("unsupported type: %T", item)
		}
	}

	return sum, nil
}

// Values 获取值
func (c *Collection[T]) Values() []T {
	return c.value
}

// ToJson 获取json
func (c *Collection[T]) ToJson() ([]byte, error) {
	return json.Marshal(c.value)
}

// FromJson 从json中获取数据
func (c *Collection[T]) FromJson(data []byte) error {
	return json.Unmarshal(data, &c.value)
}
