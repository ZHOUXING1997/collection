# 切片集合使用（slice_collcection）

## 初始化

```go
c := collection.NewSliceCollect([]int{3, 1, 2})
empty := collection.NewEmptyCollection[string]()
```

对于需要比较能力的方法（如 `Sort`/`Max`/`Contains` 等），需提供比较函数：

```go
c.SetCompare(func(a, b any) int {
    ai := a.(int)
    bi := b.(int)
    switch {
    case ai < bi:
        return -1
    case ai > bi:
        return 1
    default:
        return 0
    }
})
```

## 常用操作

- 过滤：`Filter` / 排除：`Reject`
- 查找：`Search`、`First`、`Last`
- 排序：`Sort`、`SortDesc`、`SortBy`、`SortByDesc`、`SortFloatBy`
- 聚合：`Sum`、`Avg`、`Median`、`Mode`
- 其它：`Map`、`MapFilter`、`GroupBy`、`Split`、`ForPage`、`Nth` 等

```go
filtered := c.Filter(func(item int, _ int) bool { return item > 1 })
max, _ := c.Max()
joined := c.Join(",")
```

更多 API 说明见 pkg 文档：
- [`github.com/ZHOUXING1997/collection/slice_collcection`](https://pkg.go.dev/github.com/ZHOUXING1997/collection/slice_collcection)


