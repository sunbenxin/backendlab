// Package model provides ...
package dao

import (
	"time"

	"com.github.sunbenxin/common"
)

// Order represents the model for an order
type Order struct {
	ID          uint64 `gorm:"primary_key" json:"orderId" example:"1"`
	Description string `json:"desc" example:"Leo Messi"`
	//Items        []Item    `json:"items"`

	CreatedBy uint64     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy uint64     `json:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedBy *uint64    `json:"deleted_by"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `gorm:"primary_key" json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`

	CreatedBy uint64       `json:"created_by"`
	CreatedAt common.Time  `json:"created_at"`
	UpdatedBy uint64       `json:"updated_by"`
	UpdatedAt common.Time  `json:"updated_at"`
	DeletedBy *uint64      `json:"deleted_by"`
	DeletedAt *common.Time `sql:"index" json:"deleted_at"`
}
