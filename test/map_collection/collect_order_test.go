package map_collection

import (
	"testing"

	"github.com/ZHOUXING1997/collection/map_collection"
)

func TestSetKeyCompare(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	// 设置比较函数
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 验证返回自身
	if c == nil {
		t.Error("SetKeyCompare should return itself")
	}
}

func TestOrderKey(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	// 先设置比较函数
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 执行排序
	result, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey returned error: %v", err)
	}

	if result != c {
		t.Error("OrderKey should return itself")
	}

	// 验证First返回最小的key
	k, v, ok := c.First()
	if !ok || k != "a" || v != 1 {
		t.Errorf("First after OrderKey should return 'a', got '%s'", k)
	}
}

func TestOrderKeyWithoutCompareFunc(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	// 不设置比较函数直接调用OrderKey应返回错误
	_, err := c.OrderKey()
	if err == nil {
		t.Error("OrderKey without compare function should return error")
	}
}

func TestOrderKeyByFunc(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	// 设置比较函数并立即排序
	result, err := c.OrderKeyByFunc(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	if err != nil {
		t.Errorf("OrderKeyByFunc returned error: %v", err)
	}

	if result != c {
		t.Error("OrderKeyByFunc should return itself")
	}

	// 验证排序效果
	k, v, ok := c.First()
	if !ok || k != "a" || v != 1 {
		t.Errorf("First after OrderKeyByFunc should return 'a', got '%s'", k)
	}
}

func TestSetValCompare(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	// 设置值比较函数
	result := c.SetValCompare(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	if result != c {
		t.Error("SetValCompare should return itself")
	}
}

func TestOrderValue(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	c := map_collection.NewCollection(m)

	// 设置值比较函数
	c.SetValCompare(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 按值排序
	result, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue returned error: %v", err)
	}

	if result != c {
		t.Error("OrderValue should return itself")
	}

	// 验证First返回最小值对应的key
	k, v, ok := c.First()
	if !ok || v != 1 || k != "b" {
		t.Errorf("First after OrderValue should return key 'b' with value 1, got key '%s' with value %d", k, v)
	}
}

func TestOrderValueWithoutCompareFunc(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	c := map_collection.NewCollection(m)

	// 不设置比较函数直接调用OrderValue应返回错误
	_, err := c.OrderValue()
	if err == nil {
		t.Error("OrderValue without compare function should return error")
	}
}

func TestOrderByValueFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := map[string]Person{
		"p1": {Name: "Alice", Age: 30},
		"p2": {Name: "Bob", Age: 25},
		"p3": {Name: "Charlie", Age: 35},
	}
	c := map_collection.NewCollection(m)

	// 按Age字段排序
	result, err := c.OrderByValueFunc(
		func(p Person) any {
			return p.Age
		},
		func(a, b any) int {
			ageA := a.(int)
			ageB := b.(int)
			if ageA < ageB {
				return -1
			} else if ageA > ageB {
				return 1
			}
			return 0
		},
	)

	if err != nil {
		t.Errorf("OrderByValueFunc returned error: %v", err)
	}

	if result != c {
		t.Error("OrderByValueFunc should return itself")
	}

	// 验证First返回年龄最小的
	k, v, ok := c.First()
	if !ok || v.Age != 25 || k != "p2" {
		t.Errorf("First after OrderByValueFunc should return 'p2' with age 25, got '%s' with age %d", k, v.Age)
	}

	// 验证Last返回年龄最大的
	k, v, ok = c.Last()
	if !ok || v.Age != 35 || k != "p3" {
		t.Errorf("Last after OrderByValueFunc should return 'p3' with age 35, got '%s' with age %d", k, v.Age)
	}
}

func TestOrderByValueFuncWithNilFunc(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	c := map_collection.NewCollection(m)

	// extractFunc为nil应返回错误
	_, err := c.OrderByValueFunc(nil, func(a, b any) int {
		return 0
	})
	if err == nil {
		t.Error("OrderByValueFunc with nil extractFunc should return error")
	}

	// compareFunc为nil应返回错误
	_, err = c.OrderByValueFunc(func(v int) any {
		return v
	}, nil)
	if err == nil {
		t.Error("OrderByValueFunc with nil compareFunc should return error")
	}
}

func TestOrderKeyDescending(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	// 设置降序比较函数
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return 1 // 反转比较结果实现降序
		} else if a > b {
			return -1
		}
		return 0
	})

	// 执行排序
	_, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey returned error: %v", err)
	}

	// 验证First返回最大的key
	k, v, ok := c.First()
	if !ok || k != "c" || v != 3 {
		t.Errorf("First after descending OrderKey should return 'c', got '%s'", k)
	}
}

func TestOrderValueDescending(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	c := map_collection.NewCollection(m)

	// 设置降序值比较函数
	c.SetValCompare(func(a, b int) int {
		if a < b {
			return 1 // 反转比较结果实现降序
		} else if a > b {
			return -1
		}
		return 0
	})

	// 按值排序
	_, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue returned error: %v", err)
	}

	// 验证First返回最大值对应的key
	k, v, ok := c.First()
	if !ok || v != 3 || k != "a" {
		t.Errorf("First after descending OrderValue should return key 'a' with value 3, got key '%s' with value %d", k, v)
	}
}

// ========== 极限场景测试 ==========

// 测试空集合的OrderKey
func TestOrderKeyOnEmptyCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})

	// 设置比较函数
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 空集合排序应成功
	result, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey on empty collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderKey should return itself")
	}

	// 验证仍为空
	if !c.IsEmpty() {
		t.Error("Empty collection should remain empty after OrderKey")
	}

	// First应返回false
	_, _, ok := c.First()
	if ok {
		t.Error("First on empty sorted collection should return false")
	}

	// Last应返回false
	_, _, ok = c.Last()
	if ok {
		t.Error("Last on empty sorted collection should return false")
	}
}

// 测试空集合的OrderValue
func TestOrderValueOnEmptyCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})

	// 设置值比较函数
	c.SetValCompare(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 空集合排序应成功
	result, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue on empty collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderValue should return itself")
	}

	// 验证仍为空
	if !c.IsEmpty() {
		t.Error("Empty collection should remain empty after OrderValue")
	}
}

// 测试空集合的OrderKeyByFunc
func TestOrderKeyByFuncOnEmptyCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})

	result, err := c.OrderKeyByFunc(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	if err != nil {
		t.Errorf("OrderKeyByFunc on empty collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderKeyByFunc should return itself")
	}

	if !c.IsEmpty() {
		t.Error("Empty collection should remain empty after OrderKeyByFunc")
	}
}

// 测试空集合的OrderByValueFunc
func TestOrderByValueFuncOnEmptyCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})

	result, err := c.OrderByValueFunc(
		func(v int) any {
			return v
		},
		func(a, b any) int {
			return 0
		},
	)

	if err != nil {
		t.Errorf("OrderByValueFunc on empty collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderByValueFunc should return itself")
	}

	if !c.IsEmpty() {
		t.Error("Empty collection should remain empty after OrderByValueFunc")
	}
}

// 测试单元素集合的OrderKey
func TestOrderKeyOnSingleElementCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{"a": 1})

	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	result, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey on single element collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderKey should return itself")
	}

	// 验证元素仍存在
	if c.Count() != 1 {
		t.Errorf("Single element collection should still have count 1, got %d", c.Count())
	}

	// First和Last应返回相同元素
	k1, v1, ok1 := c.First()
	k2, v2, ok2 := c.Last()

	if !ok1 || !ok2 {
		t.Error("First and Last on single element collection should return true")
	}

	if k1 != k2 || v1 != v2 || k1 != "a" || v1 != 1 {
		t.Error("First and Last should return the same single element")
	}
}

// 测试单元素集合的OrderValue
func TestOrderValueOnSingleElementCollection(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{"a": 1})

	c.SetValCompare(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	result, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue on single element collection should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderValue should return itself")
	}

	if c.Count() != 1 {
		t.Errorf("Single element collection should still have count 1, got %d", c.Count())
	}
}

// 测试排序后删除所有元素
func TestOrderKeyThenDeleteAll(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{
		"a": 1, "b": 2, "c": 3,
	})

	// 先排序
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	_, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey returned error: %v", err)
	}

	// 删除所有元素
	c.Remove("a")
	c.Remove("b")
	c.Remove("c")

	// 验证为空
	if !c.IsEmpty() {
		t.Error("After deleting all elements, collection should be empty")
	}

	if c.Count() != 0 {
		t.Errorf("After deleting all elements, Count should be 0, got %d", c.Count())
	}

	// First和Last应返回false
	_, _, ok := c.First()
	if ok {
		t.Error("First on empty collection should return false")
	}

	_, _, ok = c.Last()
	if ok {
		t.Error("Last on empty collection should return false")
	}
}

// 测试排序后Filter为空
func TestOrderValueThenFilterToEmpty(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{
		"a": 1, "b": 3, "c": 5,
	})

	// 先按值排序
	c.SetValCompare(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	_, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue returned error: %v", err)
	}

	// Filter出偶数(没有)
	filtered := c.Filter(func(v int, k string) bool {
		return v%2 == 0
	})

	// 验证结果为空
	if !filtered.IsEmpty() {
		t.Error("Filter with no matches should return empty collection")
	}

	if filtered.Count() != 0 {
		t.Errorf("Filtered collection should have count 0, got %d", filtered.Count())
	}

	// 原集合不应改变
	if c.Count() != 3 {
		t.Error("Filter should not modify original collection")
	}
}

// 测试单元素排序后删除
func TestOrderSingleElementThenDelete(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{"only": 99})

	// 排序
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	_, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey returned error: %v", err)
	}

	// 删除唯一元素
	c.Remove("only")

	// 验证为空
	if !c.IsEmpty() {
		t.Error("After deleting single element, collection should be empty")
	}
}

// 测试空集合排序后添加元素再排序
func TestOrderEmptyThenAddAndReorder(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})

	// 设置比较函数
	c.SetKeyCompare(func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// 对空集合排序
	_, err := c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey on empty collection returned error: %v", err)
	}

	// 添加元素
	c.Set("b", 2)
	c.Set("a", 1)
	c.Set("c", 3)

	// 再次排序
	_, err = c.OrderKey()
	if err != nil {
		t.Errorf("OrderKey after adding elements returned error: %v", err)
	}

	// 验证排序效果
	k, v, ok := c.First()
	if !ok || k != "a" || v != 1 {
		t.Errorf("First after reorder should return 'a', got '%s'", k)
	}

	k, v, ok = c.Last()
	if !ok || k != "c" || v != 3 {
		t.Errorf("Last after reorder should return 'c', got '%s'", k)
	}
}

// 测试OrderByValueFunc在单元素上
func TestOrderByValueFuncOnSingleElement(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	c := map_collection.NewCollection(map[string]Person{
		"p1": {Name: "Alice", Age: 30},
	})

	result, err := c.OrderByValueFunc(
		func(p Person) any {
			return p.Age
		},
		func(a, b any) int {
			ageA := a.(int)
			ageB := b.(int)
			if ageA < ageB {
				return -1
			} else if ageA > ageB {
				return 1
			}
			return 0
		},
	)

	if err != nil {
		t.Errorf("OrderByValueFunc on single element should not return error, got: %v", err)
	}

	if result != c {
		t.Error("OrderByValueFunc should return itself")
	}

	if c.Count() != 1 {
		t.Errorf("Single element collection should still have count 1, got %d", c.Count())
	}

	// First和Last应返回同一元素
	k1, v1, _ := c.First()
	k2, v2, _ := c.Last()

	if k1 != k2 || v1.Name != v2.Name || k1 != "p1" {
		t.Error("First and Last should return the same single element")
	}
}
