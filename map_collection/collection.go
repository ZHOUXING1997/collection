package map_collection

import (
	"encoding/json"
	"fmt"
	"slices"
)

// Copy 复制一个新的 Collection
func (c *Collection[K, V]) Copy() *Collection[K, V] {
	return c.cloneWithSortedKeys(Clone(c.value))
}

// IsEmpty 判断是否为空
func (c *Collection[K, V]) IsEmpty() bool {
	return len(c.value) == 0
}

// IsNotEmpty 判断是否不为空
func (c *Collection[K, V]) IsNotEmpty() bool {
	return len(c.value) != 0
}

// Count 返回 map 中键值对的数量
func (c *Collection[K, V]) Count() int {
	return len(c.value)
}

// Keys 返回所有的 key 组成的切片
func (c *Collection[K, V]) Keys() []K {
	return Keys(c.value)
}

// Values 返回所有的 value 组成的切片
func (c *Collection[K, V]) Values() []V {
	return Values(c.value)
}

// GetValue 获取指定 key 的值，如果不存在返回零值
func (c *Collection[K, V]) GetValue(key K) V {
	return Get(c.value, key)
}

// Get 获取指定 key 的值
// 返回值和是否存在的标志
func (c *Collection[K, V]) Get(key K) (V, bool) {
	val, ok := c.value[key]

	return val, ok
}

// GetOr 获取 key 对应的值；如果不存在，返回默认值 def
func (c *Collection[K, V]) GetOr(key K, def V) V {
	return GetOr(c.value, key, def)
}

// Has 判断是否存在指定的 key
func (c *Collection[K, V]) Has(key K) bool {
	return Has(c.value, key)
}

// Set 设置 key->val（直接修改当前 Collection，返回自身以支持链式调用）
func (c *Collection[K, V]) Set(key K, val V) *Collection[K, V] {
	_, exists := c.value[key]
	c.value[key] = val

	// 如果是新增的 key，插入到 sortedKeys
	if !exists {
		c.insertKeyInOrder(key)
	}

	return c
}

// Put 设置 key->val（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) Put(key K, val V) *Collection[K, V] {
	// 检查是否是新增的 key
	_, exists := c.value[key]

	newMap := Set(c.value, key, val)
	newColl := c.cloneWithSortedKeys(newMap)

	// 如果是新增的 key，插入到 sortedKeys
	if !exists {
		newColl.insertKeyInOrder(key)
	}

	return newColl
}

// Delete 删除指定的 key（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) Delete(key K) *Collection[K, V] {
	newMap := Delete(c.value, key)
	newColl := c.cloneWithSortedKeys(newMap)

	// 从新 Collection 的 sortedKeys 中移除
	newColl.removeKeyFromSorted(key)

	return newColl
}

// DeleteByFunc 删除满足条件的 key（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) DeleteByFunc(fn func(K, V) bool) *Collection[K, V] {
	newMap := Clone(c.value)
	deletedKeys := make([]K, 0)
	for k, v := range newMap {
		if fn(k, v) {
			delete(newMap, k)
			deletedKeys = append(deletedKeys, k)
		}
	}

	newColl := c.cloneWithSortedKeys(newMap)

	// 从新 Collection 的 sortedKeys 中移除被删除的keys
	newColl.removeKeyFromSorted(deletedKeys...)

	return newColl
}

// Remove 删除指定的 key（直接修改当前 Collection，返回自身以支持链式调用）
func (c *Collection[K, V]) Remove(key K) *Collection[K, V] {
	delete(c.value, key)
	c.removeKeyFromSorted(key)

	return c
}

// Merge 合并另一个 map 到当前 Collection（不修改原 Collection，返回新的 Collection）
// 键冲突时，以 other 为准
func (c *Collection[K, V]) Merge(other map[K]V) *Collection[K, V] {
	newMap := Merge(c.value, other)
	newColl := c.cloneWithSortedKeys(newMap)

	// 归并排序更新 sortedKeys
	if newColl.sortedKeys != nil {
		otherKeys := Keys(other)
		if newColl.keyCompareFunc != nil {
			slices.SortFunc(otherKeys, func(a, b K) int {
				return newColl.keyCompareFunc(a, b)
			})
		}
		newColl.sortedKeys = newColl.mergeKeys(otherKeys)
	}

	return newColl
}

// MergeCollection 合并另一个 Collection 到当前 Collection（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) MergeCollection(other *Collection[K, V]) *Collection[K, V] {
	if other == nil {
		return c.Copy()
	}

	newMap := Merge(c.value, other.value)
	newColl := c.cloneWithSortedKeys(newMap)

	// 归并排序更新 sortedKeys（直接使用 other 的 sortedKeys）
	if newColl.sortedKeys != nil && other.sortedKeys != nil {
		newColl.sortedKeys = newColl.mergeKeys(other.sortedKeys)
	}

	return newColl
}

// MergeInPlace 将另一个 map 合并到当前 Collection（直接修改当前 Collection，返回自身以支持链式调用）
func (c *Collection[K, V]) MergeInPlace(other map[K]V) *Collection[K, V] {
	// 记录新增的 keys
	newKeys := make([]K, 0)
	for k := range other {
		if _, exists := c.value[k]; !exists {
			newKeys = append(newKeys, k)
		}
	}

	MergeInPlace(c.value, other)

	// 增量更新 sortedKeys
	if c.sortedKeys != nil {
		if c.keyCompareFunc != nil {
			slices.SortFunc(newKeys, func(a, b K) int {
				return c.keyCompareFunc(a, b)
			})
		}
		c.sortedKeys = c.mergeKeys(newKeys)
	}

	return c
}

// Only 仅保留指定的 keys（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) Only(keys []K) *Collection[K, V] {
	newMap := Only(c.value, keys)
	newColl := c.cloneWithSortedKeys(newMap)

	// 从 sortedKeys 中过滤，保持顺序
	newColl.sortedKeys = c.filterSortedKeys(newMap)

	return newColl
}

// Except 排除指定的 keys（不修改原 Collection，返回新的 Collection）
func (c *Collection[K, V]) Except(keys []K) *Collection[K, V] {
	newMap := Except(c.value, keys)
	newColl := c.cloneWithSortedKeys(newMap)

	// 从 sortedKeys 中过滤，保持顺序
	newColl.sortedKeys = c.filterSortedKeys(newMap)

	return newColl
}

// Filter 过滤键值对，返回新的 Collection
// fn 函数返回 true 的键值对会被保留
func (c *Collection[K, V]) Filter(fn func(value V, key K) bool) *Collection[K, V] {
	newMap := Filter(c.value, fn)
	newColl := c.cloneWithSortedKeys(newMap)

	// 从 sortedKeys 中过滤，保持顺序
	newColl.sortedKeys = c.filterSortedKeys(newMap)

	return newColl
}

// Each 对每个键值对执行回调函数
func (c *Collection[K, V]) Each(fn func(value V, key K)) *Collection[K, V] {
	Each(c.value, fn)

	return c
}

// Foreach 对每个键值对执行回调函数（有序的）
func (c *Collection[K, V]) Foreach(fn func(value V, key K)) *Collection[K, V] {
	if c.sortedKeys == nil {
		c.initSortedKeys()
	}
	for _, k := range c.sortedKeys {
		fn(c.value[k], k)
	}

	return c
}

// Map 对 value 进行映射转换，返回新的 Collection
// 注意：此方法会改变 value 的类型，因此返回的是 map[K]R 而不是 Collection
func (c *Collection[K, V]) Map(fn func(value V, key K) V) *Collection[K, V] {
	newMap := MapValues(c.value, fn)

	return NewCollection(newMap)
}

// Reduce 聚合：将 map 折叠为一个结果
func (c *Collection[K, V]) Reduce(init any, fn func(acc any, value V, key K) any) any {
	return Reduce(c.value, init, fn)
}

// First 返回排序后的第一个键值对及是否存在（稳定，基于 sortedKeys）
func (c *Collection[K, V]) First() (K, V, bool) {
	if c.sortedKeys == nil {
		c.initSortedKeys()
	}
	if len(c.sortedKeys) == 0 {
		var zk K
		var zv V
		return zk, zv, false
	}
	firstKey := c.sortedKeys[0]
	return firstKey, c.value[firstKey], true
}

// FirstWhere 返回第一个满足条件的键值对及是否存在（稳定，按 sortedKeys 顺序查找）
func (c *Collection[K, V]) FirstWhere(fn func(value V, key K) bool) (K, V, bool) {
	if c.sortedKeys == nil {
		c.initSortedKeys()
	}
	for _, k := range c.sortedKeys {
		v := c.value[k]
		if fn(v, k) {
			return k, v, true
		}
	}
	var zk K
	var zv V
	return zk, zv, false
}

// Last 返回排序后的最后一个键值对及是否存在（稳定，基于 sortedKeys）
func (c *Collection[K, V]) Last() (K, V, bool) {
	if c.sortedKeys == nil {
		c.initSortedKeys()
	}
	if len(c.sortedKeys) == 0 {
		var zk K
		var zv V
		return zk, zv, false
	}
	lastKey := c.sortedKeys[len(c.sortedKeys)-1]
	return lastKey, c.value[lastKey], true
}

// LastWhere 返回最后一个满足条件的键值对及是否存在（稳定，按 sortedKeys 逆序查找）
func (c *Collection[K, V]) LastWhere(fn func(value V, key K) bool) (K, V, bool) {
	if c.sortedKeys == nil {
		c.initSortedKeys()
	}
	for i := len(c.sortedKeys) - 1; i >= 0; i-- {
		k := c.sortedKeys[i]
		v := c.value[k]
		if fn(v, k) {
			return k, v, true
		}
	}
	var zk K
	var zv V
	return zk, zv, false
}

// ToMap 返回底层的 map（这是一个引用，修改会影响 Collection）
func (c *Collection[K, V]) All() map[K]V {
	return c.value
}

// ToJSON 将 Collection 转换为 JSON 字符串
func (c *Collection[K, V]) ToJSON() (string, error) {
	data, err := json.Marshal(c.value)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// DD 打印 Collection 的内容（用于调试）
func (c *Collection[K, V]) DD() *Collection[K, V] {
	fmt.Printf("Collection: %+v\n", c.value)

	return c
}

// Pluck 从所有 value 中提取指定字段，返回类型安全的切片
// 适用于 V 是结构体或结构体指针的场景
// fieldName: 字段名
// R: 字段类型（需要与实际字段类型匹配）
// 注意：此方法返回的切片顺序取决于是否有 sortedKeys
func (c *Collection[K, V]) Pluck(fieldName string) []any {
	result := make([]any, 0, len(c.value))

	// 使用 sortedKeys 保持顺序（如果存在）
	if c.sortedKeys != nil {
		for _, k := range c.sortedKeys {
			if v, ok := c.value[k]; ok {
				if extracted := extractField(v, fieldName); extracted != nil {
					result = append(result, extracted)
				}
			}
		}
	} else {
		// 无序遍历
		for _, v := range c.value {
			if extracted := extractField(v, fieldName); extracted != nil {
				result = append(result, extracted)
			}
		}
	}

	return result
}

// PluckFunc 从所有 value 中提取指定内容，返回切片
// extractFunc: 用户自定义提取函数
func (c *Collection[K, V]) PluckFunc(extractFunc func(V) any) []any {
	result := make([]any, 0, len(c.value))

	// 使用 sortedKeys 保持顺序（如果存在）
	if c.sortedKeys != nil {
		for _, k := range c.sortedKeys {
			if v, ok := c.value[k]; ok {
				result = append(result, extractFunc(v))
			}
		}
	} else {
		// 无序遍历
		for _, v := range c.value {
			result = append(result, extractFunc(v))
		}
	}

	return result
}
