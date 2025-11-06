package map_collection

// CollectionOption 是用于配置 Collection 的函数式选项
type CollectionOption[K comparable, V any] func(*Collection[K, V])
