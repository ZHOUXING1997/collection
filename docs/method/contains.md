# Contains

适用集合：slice_collcection

`Contains(obj T) (bool, error)`

判断元素是否在集合中；需先设置比较函数。

```go
c := collection.NewSliceCollect([]int{1, 2, 2, 3})
c.SetCompare(func(a, b any) int {
    ai, bi := a.(int), b.(int)
    switch {
    case ai < bi: return -1
    case ai > bi: return 1
    default: return 0
    }
})

ok, _ := c.Contains(1)   // true
ok2, _ := c.Contains(5)  // false
_ = ok; _ = ok2
```