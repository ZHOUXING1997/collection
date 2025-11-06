package map_collection

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ZHOUXING1997/collection/map_collection"
)

func TestNewCollection(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)
	if c == nil {
		t.Error("NewCollection returned nil")
	}
	if c.Count() != 3 {
		t.Errorf("Expected count 3, got %d", c.Count())
	}
}

func TestCopy(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)
	copied := c.Copy()

	if !reflect.DeepEqual(copied.All(), c.All()) {
		t.Error("Copy did not copy the values correctly")
	}

	// 修改副本不应影响原集合
	copied.Set("d", 4)
	if c.Count() == copied.Count() {
		t.Error("Copy should create independent collection")
	}
}

func TestIsEmpty(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{})
	if !c.IsEmpty() {
		t.Error("Expected collection to be empty")
	}

	c.Set("a", 1)
	if c.IsEmpty() {
		t.Error("Expected collection to be not empty")
	}
}

func TestIsNotEmpty(t *testing.T) {
	c := map_collection.NewCollection(map[string]int{"a": 1})
	if !c.IsNotEmpty() {
		t.Error("Expected collection to be not empty")
	}

	c2 := map_collection.NewCollection(map[string]int{})
	if c2.IsNotEmpty() {
		t.Error("Expected collection to be empty")
	}
}

func TestCount(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)
	if c.Count() != 3 {
		t.Errorf("Expected count 3, got %d", c.Count())
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)
	keys := c.Keys()

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	for _, k := range keys {
		if !expectedKeys[k] {
			t.Errorf("Unexpected key: %s", k)
		}
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)
	values := c.Values()

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	for _, v := range values {
		if !expectedValues[v] {
			t.Errorf("Unexpected value: %d", v)
		}
	}
}

func TestGetValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	if c.GetValue("a") != 1 {
		t.Error("GetValue returned wrong value")
	}

	// 获取不存在的key应返回零值
	if c.GetValue("d") != 0 {
		t.Error("GetValue should return zero value for non-existent key")
	}
}

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	v, ok := c.Get("a")
	if !ok || v != 1 {
		t.Error("Get returned wrong value")
	}

	_, ok = c.Get("d")
	if ok {
		t.Error("Get should return false for non-existent key")
	}
}

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	if c.GetOr("a", 99) != 1 {
		t.Error("GetOr returned wrong value")
	}

	if c.GetOr("d", 99) != 99 {
		t.Error("GetOr should return default value for non-existent key")
	}
}

func TestHas(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	if !c.Has("a") {
		t.Error("Has should return true for existing key")
	}

	if c.Has("d") {
		t.Error("Has should return false for non-existent key")
	}
}

func TestSet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	c.Set("c", 3)
	if c.Count() != 3 {
		t.Error("Set did not add new key")
	}
	if c.GetValue("c") != 3 {
		t.Error("Set did not set correct value")
	}

	// 修改已存在的key
	c.Set("a", 10)
	if c.GetValue("a") != 10 {
		t.Error("Set did not update existing key")
	}
}

func TestPut(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	newC := c.Put("c", 3)

	// 原集合不应改变
	if c.Count() != 2 {
		t.Error("Put should not modify original collection")
	}

	// 新集合应包含新值
	if newC.Count() != 3 || newC.GetValue("c") != 3 {
		t.Error("Put did not create correct new collection")
	}
}

func TestDelete(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	newC := c.Delete("b")

	// 原集合不应改变
	if c.Count() != 3 {
		t.Error("Delete should not modify original collection")
	}

	// 新集合应删除指定key
	if newC.Count() != 2 || newC.Has("b") {
		t.Error("Delete did not remove key correctly")
	}
}

func TestDeleteByFunc(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	c := map_collection.NewCollection(m)

	// 删除值为偶数的项
	newC := c.DeleteByFunc(func(k string, v int) bool {
		return v%2 == 0
	})

	// 原集合不应改变
	if c.Count() != 4 {
		t.Error("DeleteByFunc should not modify original collection")
	}

	// 新集合应只保留奇数值
	if newC.Count() != 2 || newC.Has("b") || newC.Has("d") {
		t.Error("DeleteByFunc did not delete correctly")
	}
}

func TestRemove(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	c.Remove("b")

	if c.Count() != 2 || c.Has("b") {
		t.Error("Remove did not remove key correctly")
	}
}

func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	c := map_collection.NewCollection(m1)

	m2 := map[string]int{"b": 20, "c": 3}
	newC := c.Merge(m2)

	// 原集合不应改变
	if c.Count() != 2 {
		t.Error("Merge should not modify original collection")
	}

	// 新集合应合并两个map，冲突时以other为准
	if newC.Count() != 3 || newC.GetValue("b") != 20 || newC.GetValue("c") != 3 {
		t.Error("Merge did not merge correctly")
	}
}

func TestMergeCollection(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	c1 := map_collection.NewCollection(m1)

	m2 := map[string]int{"b": 20, "c": 3}
	c2 := map_collection.NewCollection(m2)

	newC := c1.MergeCollection(c2)

	// 原集合不应改变
	if c1.Count() != 2 {
		t.Error("MergeCollection should not modify original collection")
	}

	// 新集合应合并两个Collection
	if newC.Count() != 3 || newC.GetValue("b") != 20 || newC.GetValue("c") != 3 {
		t.Error("MergeCollection did not merge correctly")
	}

	// 测试nil情况
	newC2 := c1.MergeCollection(nil)
	if newC2.Count() != c1.Count() {
		t.Error("MergeCollection with nil should return copy")
	}
}

func TestMergeInPlace(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	c := map_collection.NewCollection(m1)

	m2 := map[string]int{"b": 20, "c": 3}
	c.MergeInPlace(m2)

	// 应修改原集合
	if c.Count() != 3 || c.GetValue("b") != 20 || c.GetValue("c") != 3 {
		t.Error("MergeInPlace did not merge correctly")
	}
}

func TestOnly(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	c := map_collection.NewCollection(m)

	newC := c.Only([]string{"a", "c"})

	// 原集合不应改变
	if c.Count() != 4 {
		t.Error("Only should not modify original collection")
	}

	// 新集合应只包含指定的key
	if newC.Count() != 2 || !newC.Has("a") || !newC.Has("c") || newC.Has("b") {
		t.Error("Only did not filter correctly")
	}
}

func TestExcept(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	c := map_collection.NewCollection(m)

	newC := c.Except([]string{"b", "d"})

	// 原集合不应改变
	if c.Count() != 4 {
		t.Error("Except should not modify original collection")
	}

	// 新集合应排除指定的key
	if newC.Count() != 2 || newC.Has("b") || newC.Has("d") {
		t.Error("Except did not filter correctly")
	}
}

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	c := map_collection.NewCollection(m)

	// 过滤出偶数值
	newC := c.Filter(func(v int, k string) bool {
		return v%2 == 0
	})

	// 原集合不应改变
	if c.Count() != 4 {
		t.Error("Filter should not modify original collection")
	}

	// 新集合应只包含偶数值
	if newC.Count() != 2 || !newC.Has("b") || !newC.Has("d") {
		t.Error("Filter did not filter correctly")
	}
}

func TestEach(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	sum := 0
	c.Each(func(v int, k string) {
		sum += v
	})

	if sum != 6 {
		t.Errorf("Each did not iterate correctly, expected sum 6, got %d", sum)
	}
}

func TestForeach(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	sum := 0
	c.Foreach(func(v int, k string) {
		sum += v
	})

	if sum != 6 {
		t.Errorf("Foreach did not iterate correctly, expected sum 6, got %d", sum)
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	// 将所有值翻倍
	newC := c.Map(func(v int, k string) int {
		return v * 2
	})

	// 原集合不应改变
	if c.GetValue("a") != 1 {
		t.Error("Map should not modify original collection")
	}

	// 新集合的值应翻倍
	if newC.GetValue("a") != 2 || newC.GetValue("b") != 4 || newC.GetValue("c") != 6 {
		t.Error("Map did not transform values correctly")
	}
}

func TestReduce(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	// 求和
	sum := c.Reduce(0, func(acc any, v int, k string) any {
		return acc.(int) + v
	})

	if sum.(int) != 6 {
		t.Errorf("Reduce did not calculate correctly, expected 6, got %d", sum)
	}
}

func TestFirst(t *testing.T) {
	m := map[string]int{"a": 1}
	c := map_collection.NewCollection(m)

	k, v, ok := c.First()
	if !ok || k != "a" || v != 1 {
		t.Error("First did not return correct value")
	}

	// 空集合
	emptyC := map_collection.NewCollection(map[string]int{})
	_, _, ok = emptyC.First()
	if ok {
		t.Error("First should return false for empty collection")
	}
}

func TestFirstWhere(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	_, v, ok := c.FirstWhere(func(v int, k string) bool {
		return v > 1
	})

	if !ok || v <= 1 {
		t.Error("FirstWhere did not find correct value")
	}

	// 找不到匹配项
	_, _, ok = c.FirstWhere(func(v int, k string) bool {
		return v > 10
	})
	if ok {
		t.Error("FirstWhere should return false when no match found")
	}
}

func TestLast(t *testing.T) {
	m := map[string]int{"a": 1}
	c := map_collection.NewCollection(m)

	k, v, ok := c.Last()
	if !ok || k != "a" || v != 1 {
		t.Error("Last did not return correct value")
	}

	// 空集合
	emptyC := map_collection.NewCollection(map[string]int{})
	_, _, ok = emptyC.Last()
	if ok {
		t.Error("Last should return false for empty collection")
	}
}

func TestLastWhere(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	_, v, ok := c.LastWhere(func(v int, k string) bool {
		return v < 3
	})

	if !ok || v >= 3 {
		t.Error("LastWhere did not find correct value")
	}

	// 找不到匹配项
	_, _, ok = c.LastWhere(func(v int, k string) bool {
		return v > 10
	})
	if ok {
		t.Error("LastWhere should return false when no match found")
	}
}

func TestAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	all := c.All()
	if len(all) != 3 || all["a"] != 1 || all["b"] != 2 || all["c"] != 3 {
		t.Error("All did not return correct map")
	}
}

func TestToJSON(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	c := map_collection.NewCollection(m)

	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Errorf("ToJSON returned error: %v", err)
	}

	// 反序列化验证
	var result map[string]int
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if !reflect.DeepEqual(result, m) {
		t.Error("ToJSON did not produce correct JSON")
	}
}

func TestDD(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	c := map_collection.NewCollection(m)

	// DD应该返回自身
	result := c.DD()
	if result != c {
		t.Error("DD should return itself")
	}
}

type Person struct {
	Name string
	Age  int
}

func TestPluck(t *testing.T) {
	m := map[string]Person{
		"p1": {Name: "Alice", Age: 30},
		"p2": {Name: "Bob", Age: 25},
		"p3": {Name: "Charlie", Age: 35},
	}
	c := map_collection.NewCollection(m)

	names := c.Pluck("Name")
	if len(names) != 3 {
		t.Errorf("Pluck returned wrong number of items: %d", len(names))
	}

	// 验证是否包含所有名字
	nameMap := make(map[string]bool)
	for _, name := range names {
		if name != nil {
			nameMap[name.(string)] = true
		}
	}

	if !nameMap["Alice"] || !nameMap["Bob"] || !nameMap["Charlie"] {
		t.Error("Pluck did not extract all names correctly")
	}
}

func TestPluckFunc(t *testing.T) {
	m := map[string]Person{
		"p1": {Name: "Alice", Age: 30},
		"p2": {Name: "Bob", Age: 25},
		"p3": {Name: "Charlie", Age: 35},
	}
	c := map_collection.NewCollection(m)

	ages := c.PluckFunc(func(p Person) any {
		return p.Age
	})

	if len(ages) != 3 {
		t.Errorf("PluckFunc returned wrong number of items: %d", len(ages))
	}

	// 验证是否包含所有年龄
	ageMap := make(map[int]bool)
	for _, age := range ages {
		if age != nil {
			ageMap[age.(int)] = true
		}
	}

	if !ageMap[30] || !ageMap[25] || !ageMap[35] {
		t.Error("PluckFunc did not extract all ages correctly")
	}
}
