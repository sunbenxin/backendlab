package service

import (
	"time"

	"com.github.sunbenxin/dal"
	"com.github.sunbenxin/dao"
)

type OrderSrv struct {
}

var Order *OrderSrv = &OrderSrv{}

// Get
func (svc *OrderSrv) Get(ID string) (*dao.Order, error) {
	order, err := dal.Order.Get(ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

type OrderParam struct {
	Description *string `json:"description" binding:"required"`
}

// Create
func (svc *OrderSrv) Create(param OrderParam) (*dao.Order, error) {
	now := time.Now().Unix()
	order := dao.Order{
		Description: *param.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   1,
		UpdatedBy:   1,
	}

	created, err := dal.Order.Create(&order)
	if err != nil {
		return nil, err
	}

	return created, nil
}

// UpdateOrder
func (svc *OrderSrv) Update(ID string, param OrderParam) (*dao.Order, error) {

	order, err := dal.Order.Get(ID)
	if err != nil {
		/*
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
		*/
		return nil, err
	}

	now := time.Now().Unix()

	order.Description = *param.Description
	order.UpdatedAt = now
	order.UpdatedBy = 1

	updated, err := dal.Order.Update(order)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// DeleteOrder
func (svc *OrderSrv) Delete(ID string) (*dao.Order, error) {

	order, err := dal.Order.Get(ID)

	if err != nil {
		return nil, err
	}

	return order, dal.Order.Delete(order)

}

/*

func (svc *OrderSrv) ListOrders(req *dao.ListOrdersReq) ([]dao.Order, uint64, error) {
	orders, total, err := dal.Order.List(req)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}
*/
