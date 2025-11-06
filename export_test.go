package collection

import (
	"testing"

	"github.com/ZHOUXING1997/collection/map_collection"
)

// TestNewMapCollect_Basic 测试基本功能
func TestNewMapCollect_Basic(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := NewMapCollect(m)

	if c == nil {
		t.Error("NewMapCollect returned nil")
	}
	if c.Count() != 3 {
		t.Errorf("Expected count 3, got %d", c.Count())
	}
}

// TestNewMapCollect_WithKeyCompare 测试使用key比较函数
func TestNewMapCollect_WithKeyCompare(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := NewMapCollect(m, map_collection.WithKeyCompare[string, int](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}))

	// 验证First返回排序后的第一个
	k, v, ok := c.First()
	if !ok || k != "a" || v != 1 {
		t.Errorf("Expected first key 'a' with value 1, got key '%s' with value %d", k, v)
	}
}

// TestNewMapCollect_WithValCompare 测试使用value比较函数
func TestNewMapCollect_WithValCompare(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	c := NewMapCollect(m, map_collection.WithValCompare[string, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}))

	// SetValCompare后需要OrderValue才能排序
	c, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue returned error: %v", err)
	}

	k, v, ok := c.First()
	if !ok || v != 1 || k != "b" {
		t.Errorf("Expected first key 'b' with value 1, got key '%s' with value %d", k, v)
	}
}

// TestNewMapCollect_WithBothCompare 测试同时使用key和value比较函数
func TestNewMapCollect_WithBothCompare(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	c := NewMapCollect(m,
		map_collection.WithKeyCompare[string, int](func(a, b string) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		}),
		map_collection.WithValCompare[string, int](func(a, b int) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		}),
	)

	// 默认按key排序
	k, _, ok := c.First()
	if !ok || k != "a" {
		t.Errorf("Expected first key 'a', got '%s'", k)
	}

	// 切换到按value排序
	c, err := c.OrderValue()
	if err != nil {
		t.Errorf("OrderValue returned error: %v", err)
	}

	k, v, ok := c.First()
	if !ok || v != 1 || k != "a" {
		t.Errorf("Expected first key 'a' with value 1, got key '%s' with value %d", k, v)
	}
}

// TestNewMapCollect_EmptyMap 测试空map
func TestNewMapCollect_EmptyMap(t *testing.T) {
	m := map[string]int{}
	c := NewMapCollect(m)

	if c.Count() != 0 {
		t.Errorf("Expected count 0, got %d", c.Count())
	}

	_, _, ok := c.First()
	if ok {
		t.Error("First on empty map should return false")
	}
}

// TestNewMapCollect_NilMap 测试nil map
func TestNewMapCollect_NilMap(t *testing.T) {
	var m map[string]int
	c := NewMapCollect(m)

	if c == nil {
		t.Error("NewMapCollect with nil map should not return nil")
	}

	if c.Count() != 0 {
		t.Errorf("Expected count 0 for nil map, got %d", c.Count())
	}
}

// TestNewMapCollect_LargeMap 测试大量数据
func TestNewMapCollect_LargeMap(t *testing.T) {
	m := make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		m[i] = i * 2
	}

	c := NewMapCollect(m, map_collection.WithKeyCompare[int, int](func(a, b int) int {
		return a - b
	}))

	if c.Count() != 10000 {
		t.Errorf("Expected count 10000, got %d", c.Count())
	}

	k, v, ok := c.First()
	if !ok || k != 0 || v != 0 {
		t.Errorf("Expected first key 0 with value 0, got key %d with value %d", k, v)
	}

	k, v, ok = c.Last()
	if !ok || k != 9999 || v != 19998 {
		t.Errorf("Expected last key 9999 with value 19998, got key %d with value %d", k, v)
	}
}

// TestNewMapCollect_NoOpts 测试不传opts参数
func TestNewMapCollect_NoOpts(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	c := NewMapCollect(m)

	// 不传opts时,不自动排序
	if c.Count() != 2 {
		t.Errorf("Expected count 2, got %d", c.Count())
	}
}

// TestNewMapCollect_MultipleOpts 测试多个opts
func TestNewMapCollect_MultipleOpts(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	c := NewMapCollect(m,
		map_collection.WithKeyCompare[string, int](func(a, b string) int {
			if a < b {
				return -1
			}
			return 1
		}),
		map_collection.WithValCompare[string, int](func(a, b int) int {
			return a - b
		}),
	)

	if c == nil {
		t.Error("NewMapCollect with multiple opts returned nil")
	}
}

// TestNewMapCollect_DescendingOrder 测试降序排序
func TestNewMapCollect_DescendingOrder(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := NewMapCollect(m, map_collection.WithKeyCompare[string, int](func(a, b string) int {
		// 反转比较实现降序
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}))

	k, v, ok := c.First()
	if !ok || k != "c" || v != 3 {
		t.Errorf("Expected first key 'c' with value 3, got key '%s' with value %d", k, v)
	}

	k, v, ok = c.Last()
	if !ok || k != "a" || v != 1 {
		t.Errorf("Expected last key 'a' with value 1, got key '%s' with value %d", k, v)
	}
}

// TestNewMapCollect_DuplicateKeys 测试重复的key(map本身不允许,但测试数据完整性)
func TestNewMapCollect_DuplicateKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	m["a"] = 10 // 覆盖

	c := NewMapCollect(m)

	if c.GetValue("a") != 10 {
		t.Errorf("Expected value 10 for key 'a', got %d", c.GetValue("a"))
	}
	if c.Count() != 2 {
		t.Errorf("Expected count 2, got %d", c.Count())
	}
}

// TestNewEmptyMapCollection 测试空集合创建
func TestNewEmptyMapCollection(t *testing.T) {
	c := NewEmptyMapCollection[string, int]()

	if c == nil {
		t.Error("NewEmptyMapCollection returned nil")
	}
	if c.Count() != 0 {
		t.Errorf("Expected count 0, got %d", c.Count())
	}
	if !c.IsEmpty() {
		t.Error("Expected empty collection")
	}
}

// TestNewEmptyMapCollection_AddData 测试空集合添加数据
func TestNewEmptyMapCollection_AddData(t *testing.T) {
	c := NewEmptyMapCollection[string, int]()
	c.Set("a", 1).Set("b", 2)

	if c.Count() != 2 {
		t.Errorf("Expected count 2, got %d", c.Count())
	}
	if c.GetValue("a") != 1 {
		t.Errorf("Expected value 1 for key 'a', got %d", c.GetValue("a"))
	}
}

// TestNewMapCollect_StructKey 测试结构体作为key
func TestNewMapCollect_StructKey(t *testing.T) {
	type Key struct {
		ID int
	}
	m := map[Key]string{
		{ID: 1}: "one",
		{ID: 2}: "two",
	}

	c := NewMapCollect(m, map_collection.WithKeyCompare[Key, string](func(a, b Key) int {
		return a.ID - b.ID
	}))

	if c.Count() != 2 {
		t.Errorf("Expected count 2, got %d", c.Count())
	}

	k, v, ok := c.First()
	if !ok || k.ID != 1 || v != "one" {
		t.Errorf("Expected first key {ID:1} with value 'one', got key {ID:%d} with value '%s'", k.ID, v)
	}
}

// TestNewMapCollect_PointerValue 测试指针类型的value
func TestNewMapCollect_PointerValue(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{Name: "Alice", Age: 30}
	p2 := &Person{Name: "Bob", Age: 25}

	m := map[string]*Person{
		"p1": p1,
		"p2": p2,
	}

	c := NewMapCollect(m)

	if c.Count() != 2 {
		t.Errorf("Expected count 2, got %d", c.Count())
	}

	v := c.GetValue("p1")
	if v == nil || v.Name != "Alice" {
		t.Error("Failed to get pointer value correctly")
	}
}
