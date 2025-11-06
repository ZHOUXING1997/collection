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
