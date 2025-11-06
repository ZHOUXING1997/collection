# Avg

适用集合：slice_collcection

`Avg() (float64, error)`

返回集合的数值平均数，支持整型与浮点型元素，返回 `float64`。

```go
c := collection.NewSliceCollect([]int{1, 2, 2, 3})
avg, err := c.Avg()
if err != nil {
    panic(err)
}
if avg != 2.0 {
    panic("Avg error")
}
```