package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {
	{
		api.POST("orders", h.OrderCreate)
		//api.PUT("orders/:id", h.OrderUpdate)
		api.GET("orders-by-filter", h.OrdersGetByFilter)
		api.GET("orders/:id", h.OrderGetByID)
		api.DELETE("orders/:id", h.OrderDelete)
	}
}

// OrderCreate
//	@Summary		Create an order.
//	@Description	This API to create an order.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.CreateOrderWithProductsInput	true	"data body"
//	@Success		201		{object}	models.Orders
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/orders [POST]
func (h *Handler) OrderCreate(c *gin.Context) {
	var body models.CreateOrderWithProductsInput

	err := c.ShouldBindJSON(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json OrderCreate", logger.Error(err))
		return
	}

	order, err := h.services.Orders.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		h.log.Error("error while create order", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, order)
}

// OrderUpdate
//	@Summary		Update an order.
//	@Description	This API to update an order.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"id for update Order"	Format(id)
//	@Param			data	body		models.UpdateOrderInput	true	"data body"
//	@Success		200		{object}	models.Orders
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/orders/{id} [PUT]
func (h *Handler) OrderUpdate(c *gin.Context) {
	var body models.UpdateOrderInput

	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 64)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil || inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json OrderUpdate", logger.Error(err))
		return
	}

	body.ID = ID

	order, err := h.services.Orders.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while update order", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &order)
}

// OrdersGetByFilter
//	@Summary		Get orders by filter.
//	@Description	This API to get orders by filter.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			data	query		models.GetOrdersByFilterInput	true	"data body"
//	@Success		200		{object}	[]models.Orders
//	@Failure		500		{object}	Response
//	@Router			/v1/orders-by-filter [GET]
func (h *Handler) OrdersGetByFilter(c *gin.Context) {

	var body models.GetOrdersByFilterInput

	err := c.ShouldBindQuery(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json OrdersGetWithFilter", logger.Error(err))
		return
	}

	orders, err := h.services.Orders.GetAllByFilter(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while get all orders", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, orders)
}

// OrderGetByID
//	@Summary		Get order by id.
//	@Description	This API to get order by id.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id for get Order"	Format(id)
//	@Success		200		{object}	models.Orders
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/orders/{id} [GET]
func (h *Handler) OrderGetByID(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 64)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderGetByID", logger.Error(err))
		return
	}

	order, err := h.services.Orders.GetByID(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, order)
}

// OrderDelete this api deletes Order
//	@Summary		Delete an order.
//	@Description	This API to delete an order.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id for delete Order"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/orders/{id} [DELETE]
func (h *Handler) OrderDelete(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 64)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderDelete", logger.Error(err))
		return
	}

	err = h.services.Orders.DeleteByID(c.Request.Context(), ID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while delete orders by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
