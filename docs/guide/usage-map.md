# Map 集合使用（map_collection）
## 初始化
```go
m := map[string]int{"a": 2, "b": 1}
mc := collection.NewMapCollect(m)
```
## Key/Value 排序与遍历
- 按 key 排序遍历：
```go
mc.SetKeyCompare(func(a, b string) int {
    if a < b { return -1 }
    if a > b { return 1 }
    return 0
}).OrderKey()
mc.Foreach(func(v int, k string) { /* ... */ })
```
- 按 value 排序（更新有序键列表）：
```go
mc.SetValCompare(func(a, b int) int {
    switch {
    case a < b: return -1
    case a > b: return 1
    default: return 0
    }
}).OrderValue()
```
## 常用操作
- 提取：`Keys`、`Values`、`Pluck`、`PluckFunc`
- 修改：`Set`（就地）、`Put`（返回新集合）、`Merge`/`MergeInPlace`
- 过滤：`Filter`、`Only`、`Except`
- 定位与聚合：`First`、`Last`、`FirstWhere`、`LastWhere`、`Reduce`
- 序列化：`ToJSON`
```go
k, v, ok := mc.First()
jsonStr, _ := mc.ToJSON()
```
更多 API 说明见 pkg 文档：
- [`github.com/ZHOUXING1997/collection/map_collection`](https://pkg.go.dev/github.com/ZHOUXING1997/collection/map_collection)

