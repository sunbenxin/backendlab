// Package database provides ...
package dal

import (
	"errors"

	"com.github.sunbenxin/dao"
	"gorm.io/gorm"
)

var Order *OrderDB = &OrderDB{}

type OrderDB struct {
}

func (da *OrderDB) Get(orderID string) (*dao.Order, error) {
	order := dao.Order{}
	if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

func (da *OrderDB) Create(order *dao.Order) (*dao.Order, error) {
	if err := db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (da *OrderDB) Update(order *dao.Order) (*dao.Order, error) {
	if err := db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (da *OrderDB) Delete(order *dao.Order) error {
	if err := db.Delete(order).Error; err != nil {
		return err
	}
	return nil
}
