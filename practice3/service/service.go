package service

import (
	"context"
	"geektimepractice/practice3/biz"
)

type ItemService struct {
	ob *biz.OrderBiz
}

// NewItemService 创建一个新 service
func NewItemService() interface{}{
	return &ItemService{}
}

// DtoToDo 路由到这里，主要做到事情是 dto -> do
func (is *ItemService) DtoToDo(ctx context.Context, interfaceReq interface{}) (interfaceRes interface{}, err error) {
	// 先拿到业务中 Order 对象
	o := new(biz.Order)

	// todo: dto -> do，从 interfaceReq 拿到数据并且组装 biz 层中到 DO 对象

	// 组装完成 do 后直接往 biz 层中传入 do
	is.ob.Buy(o)

	// todo: 返回时候需要带一些信息
	return interfaceRes, nil
}
