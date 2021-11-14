package data

import "geektimepractice/practice3/biz"

// NewOrderRepo 新建一个 orderRepo
func NewOrderRepo() biz.OrderRepo {
	return new(orderRepo)
}

// OrderRepo 接口的实现
type orderRepo struct {}

// SaveOrder 有点类似 java 中 mapper，主要做了 do -> po 的工作，是 biz 层中具体数据库操作接口方法的实现
func (or *orderRepo) SaveOrder(o *biz.Order) {
	// todo：DO -> PO，数据持久化处理
}


