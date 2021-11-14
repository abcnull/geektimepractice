package biz

// Order 把 do 也定义在了业务逻辑处理层，有点类似 java 中的 bean
type Order struct {
	OrderId string
}

// OrderRepo 有点类似 java 中的 dao
type OrderRepo interface {
	SaveOrder(*Order)
}

// OrderBiz 业务层对象
type OrderBiz struct {
	repo OrderRepo
}

// Buy 业务层中的购买逻辑
func (ob *OrderBiz) Buy(o *Order) error {

	// todo: 这里写具体的业务逻辑，调用 OrderRepo 接口操作（被 data 所实现），这些操作实际上就是数据库 crud

	return nil
}
