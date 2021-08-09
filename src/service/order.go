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
func (svc *OrderSrv) Get(orderID string) (*dao.Order, error) {
	order, err := dal.Order.Get(orderID)
	if err != nil {
		//return nil, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
		return nil, err
	}

	return order, nil
}

type OrderParam struct {
	Description *string `json:"description"`
}

// Create
func (svc *OrderSrv) Create(param OrderParam) (*dao.Order, error) {
	now := time.Now()
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

/*
// UpdateOrder
func (svc *OrderSrv) UpdateOrder(orderId uint64, param OrderParam) (*dao.Order, error) {

	order, err := dal.Order.Get(orderId)
	if err != nil {
		return nil, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
	}
	if order == nil {
		return nil, &rest.CError{Message: "order does not exist", Code: rest.RequestDataNotExisted}
	}

	now := common.Now()

	//order.Xxx     = param.Xxx
	order.UpdatedAt = now
	order.UpdatedBy = login.ID

	updated, err := dal.Order.Update(order)
	if err != nil {
		return nil, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
	}

	updated, err = svc.FillOrderWith(updated, true)
	if err != nil {
		return nil, &rest.CError{InnerErr: err, Code: rest.InternalError}
	}

	return updated, nil
}

// DeleteOrder
func (svc *OrderSrv) DeleteOrder(orderId uint64) (*dao.Order, error) {

	order, err := dal.Order.Get(orderId)

	if err != nil {
		return nil, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
	}

	if order == nil {
		return nil, &rest.CError{Message: "order does not exist", Code: rest.RequestDataNotExisted}
	}

	order, err = svc.FillOrderWith(order, true)
	if err != nil {
		return nil, &rest.CError{InnerErr: err, Code: rest.InternalError}
	}

	return order, dal.Order.Delete(order)

}

// FillOrderWith
func (svc *OrderSrv) FillOrderWith(order *dao.Order, withUser bool) (*dao.Order, error) {
	if order == nil || !withUser {
		return order, nil
	}

	orders := db.OrderArray{}
	orders = append(orders, *order)
	err := svc.FillOrdersWith(&orders, withUser)
	if err != nil {
		return order, err
	}

	return &orders[0], nil
}

// FillOrdersWith
func (svc *OrderSrv) FillOrdersWith(orders *dao.OrderArray, withUser bool) error {
	if orders == nil {
		return nil
	}

	if withUser {
		userIDs := orders.GetUIDs()
		if len(userIDs) > 0 {
			users, err := svc.BatchGetUserMap(userIDs)
			if err != nil {
				return err
			}
			orders.FillUsers(users)
		}
	}

	return nil
}

func (svc *OrderSrv) ListOrders(req *dao.ListOrdersReq) (*dao.OrderArray, uint64, error) {
	orders, total, err := dal.Order.List(req)
	if err != nil {
		return nil, total, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
	}

	if orders != nil {
		err = svc.FillOrdersWith(orders, true)
		if err != nil {
			return nil, total, &rest.CError{InnerErr: err, Code: rest.DatabaseError}
		}
	}
	return orders, total, nil
}
*/
