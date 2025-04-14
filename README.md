# Collection

Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求极致性能，追求业务开发效能的场景。

[![Go Reference](https://pkg.go.dev/badge/github.com/ZHOUXING1997/collection.svg)](https://pkg.go.dev/github.com/ZHOUXING1997/collection)

> 本项目 forked from [jianfengye/collection](https://github.com/jianfengye/collection)，并修改部分内容以适配自己的项目。

## 简介

Collection包封装了对切片（Slice）的各种操作，使其更符合业务开发的语义，提高开发效率。该库支持多种数据类型，包括基本类型和结构体类型，并提供了丰富的方法来操作和转换这些数据。

## 安装

```bash
go get github.com/ZHOUXING1997/collection@v0.0.5
```

## 支持的类型

Collection包目前支持的元素类型：
- int32, int, int64
- uint32, uint, uint64
- float32, float64
- string
- object (结构体)
- objectPoint (结构体指针)

## 使用示例

### 初始化Collection

```go
// 初始化整数集合
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})

// 初始化字符串集合
strColl := NewStrCollection([]string{"a", "b", "c"})

// 初始化结构体集合
type User struct {
    Name string
    Age  int
}
users := []User{{Name: "张三", Age: 18}, {Name: "李四", Age: 20}}
objColl := NewObjCollection(users)
```

### 常用方法示例

```go
// 过滤集合
filtered := intColl.Filter(func(item interface{}, key int) bool {
    return item.(int) > 2
})

// 映射集合
mapped := intColl.Map(func(item interface{}, key int) interface{} {
    return item.(int) * 2
})

// 排序集合
sorted := intColl.Sort()

// 查找元素
index := intColl.Search(3) // 返回元素3的索引

// 获取第一个元素
first, _ := intColl.First().ToInt()

// 获取最后一个元素
last, _ := intColl.Last().ToInt()

// 判断是否包含某元素
contains := intColl.Contains(3)
```

## 更多示例

更多使用示例请参考 [使用手册](http://collection.funaio.cn/)

## 版本历史

| 版本     | 说明                                  |
|--------|-------------------------------------|
| v0.0.6 | 重新打tag，无功能变更                        |
| v0.0.5 | 添加 pkg.go.dev 文档支持，完善包和主要类型的文档注释    |
| v0.0.4 | 修复部分问题，优化代码结构                       |
| v0.0.3 | 增加新的功能和测试用例                         |
| v0.0.2 | 修复初始版本的问题，改进性能                      |
| v0.0.1 | 初始版本，从 jianfengye/collection 项目分支而来 |

## 许可证

`collection` 使用 [Apache License 2.0](LICENSE) 许可证。
