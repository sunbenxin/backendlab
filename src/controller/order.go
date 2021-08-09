package controller

import (
	"net/http"

	"com.github.sunbenxin/dao"
	"com.github.sunbenxin/rest"
	"com.github.sunbenxin/service"
	"github.com/gin-gonic/gin"
)

type OrderCtl struct {
}

var Order *OrderCtl = &OrderCtl{}

// GetOrder
func (ctl *OrderCtl) GetOrder(c *gin.Context) {
	order, err := service.Order.GetOrder(c.Param("order_id"))
	if err != nil {
		//c.Error(&rest.CError{InnerErr: err, Code: rest.InternalServiceError})
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

// CreateOrder
func (ctl *Controller) CreateOrder(c *gin.Context) {

	param := service.OrderParam{}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.RequestParameterInvalid})
		return
	}

	logUser := GetLoginUser(c)
	order, err := ctl.Service.CreateOrder(param, logUser)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.InternalServiceError})
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: &rest.Meta{
			Code: rest.CodeSuccess,
		},
		Data: struct {
			Order *db.Order `json:"order"`
		}{
			Order: order,
		},
	})

}

// UpdateOrder
func (ctl *Controller) UpdateOrder(c *gin.Context) {

	orderId, err := strconv.ParseUint(c.Param("order_id"), 10, 64)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.RequestParameterInvalid})
		return
	}

	param := service.OrderParam{}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.RequestParameterInvalid})
		return
	}

	logUser := GetLoginUser(c)

	order, err := ctl.Service.UpdateOrder(orderId, param, logUser)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.InternalServiceError})
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: &rest.Meta{
			Code: rest.CodeSuccess,
		},
		Data: struct {
			Order *db.Order `json:"order"`
		}{
			Order: order,
		},
	})

}

// DeleteOrder
func (ctl *Controller) DeleteOrder(c *gin.Context) {
	// params
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 64)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.RequestParameterInvalid})
		return
	}

	order, err := ctl.Service.DeleteOrder(orderID)
	if err != nil {
		c.Error(&rest.CError{InnerErr: err, Code: rest.InternalServiceError})
		return
	}

	c.JSON(http.StatusOK, &rest.BaseResp{
		Meta: &rest.Meta{
			Code: rest.CodeSuccess,
		},
		Data: struct {
			Order *db.Order `json:"order"`
		}{
			Order: order,
		},
	})
}

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
