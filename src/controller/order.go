package controller

import (
	"net/http"

	"com.github.sunbenxin/dao"
	"com.github.sunbenxin/rest"
	"com.github.sunbenxin/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OrderCtl struct {
}

var Order *OrderCtl = &OrderCtl{}

// GetOrder
func (ctl *OrderCtl) Get(c *gin.Context) {
	order, err := service.Order.Get(c.Param("order_id"))
	if err != nil {
		logrus.Errorf("Get order err:%v", err)
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: rest.Meta{
			Code: 0,
		},
		Data: struct {
			Order *dao.Order `json:"order"`
		}{
			Order: order,
		},
	})

}

// CreateOrder
func (ctl *OrderCtl) Create(c *gin.Context) {

	param := service.OrderParam{}
	if err := c.ShouldBindJSON(&param); err != nil {
		logrus.Errorf("bind err:%+v", err)
	}

	order, err := service.Order.Create(param)
	if err != nil {
		logrus.Errorf("create order err:%+v", err)
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: rest.Meta{
			Code: 0,
		},
		Data: struct {
			Order *dao.Order `json:"order"`
		}{
			Order: order,
		},
	})

}

// UpdateOrder
func (ctl *OrderCtl) Update(c *gin.Context) {

	param := service.OrderParam{}

	if err := c.ShouldBindJSON(&param); err != nil {
		logrus.Errorf("bind err:%+v", err)
		return
	}

	order, err := service.Order.Update(c.Param("order_id"), param)
	if err != nil {
		logrus.Errorf("update order err:%+v", err)
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: rest.Meta{
			Code: 0,
		},
		Data: struct {
			Order *dao.Order `json:"order"`
		}{
			Order: order,
		},
	})

}

// DeleteOrder
func (ctl *OrderCtl) Delete(c *gin.Context) {
	order, err := service.Order.Delete(c.Param("order_id"))
	if err != nil {
		logrus.Errorf("delete order err:%+v", err)
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: rest.Meta{
			Code: 0,
		},
		Data: struct {
			Order *dao.Order `json:"order"`
		}{
			Order: order,
		},
	})
}

/*
// ListOrders
func (ctl *Controller) ListOrders(c *gin.Context) {
	req := db.ListOrdersReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.RequestParameterInvalid})
		return
	}

	Orders, total, err := ctl.Service.ListOrders(&req)

	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.InternalServiceError})
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: &rest.Meta{
			Code: rest.CodeSuccess,
		},
		Data: struct {
			Orders *db.OrderArray `json:"rows"`
			Total  uint64         `json:"total"`
		}{
			Orders: Orders,
			Total:  total,
		},
	})

}
*/
