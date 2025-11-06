# Sum

适用集合：slice_collcection

`Sum() (float64, error)`

返回集合元素的和（整型、浮点型），以 `float64` 表示。

```go
c := collection.NewSliceCollect([]int{1, 2, 2, 3})
sum, err := c.Sum()
if err != nil {
    panic(err)
}
if sum != 8 {
    panic("sum 错误")
}
```