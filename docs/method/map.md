# Map

适用集合：slice_collcection

`Map(func(item T, key int) T) *Collection[T]`

对集合中的每个元素进行函数映射，返回同类型的新集合。若需“映射并过滤”，请使用 `MapFilter`。

```go
c := collection.NewSliceCollect([]int{1, 2, 3, 4})
newC := c.Map(func(item int, _ int) int { return item * 2 })
newC.DD()

// 结果：2,4,6,8
```