package v1

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/venomuz/kegel-backend/internal/models"
//	"github.com/venomuz/kegel-backend/pkg/logger"
//	"net/http"
//	"strconv"
//)
//
//func (h *Handler) initOrderProductsRoutes(api *gin.RouterGroup) {
//	{
//		api.POST("order-products", h.OrderProductCreate)
//		//api.PUT("order-products/:id", h.OrderProductUpdate)
//		api.GET("order-products/:id", h.OrderProductGetByID)
//		//api.DELETE("order-products/:id", h.OrderProductDelete)
//	}
//}
//
//// OrderProductCreate
////	@Summary		Create a orderProduct.
////	@Description	This API to create an orderProduct.
////	@Tags			OrderProducts
////	@Accept			json
////	@Produce		json
////	@Param			data	body		models.CreateOrderProductsInput	true	"data body"
////	@Success		201		{object}	models.OrderProducts
////	@Failure		400,409	{object}	Response
////	@Failure		500		{object}	Response
////	@Router			/v1/order-products [POST]
//func (h *Handler) OrderProductCreate(c *gin.Context) {
//	var body models.CreateOrderProductsInput
//
//	err := c.ShouldBindJSON(&body)
//
//	if err != nil {
//		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
//		h.log.Error("error while bind to json OrderProductCreate", logger.Error(err))
//		return
//	}
//
//	orderProduct, err := h.services.OrderProducts.Create(c.Request.Context(), body)
//	if err != nil {
//		newResponse(c, http.StatusConflict, err.Error())
//		h.log.Error("error while create orderProduct", logger.Error(err))
//		return
//	}
//
//	c.JSON(http.StatusCreated, orderProduct)
//}
//
//// OrderProductUpdate
////	@Summary		Update a orderProduct.
////	@Description	This API to update an orderProduct.
////	@Tags			OrderProducts
////	@Accept			json
////	@Produce		json
////	@Param			id		path		int								true	"id for update OrderProduct"	Format(id)
////	@Param			data	body		models.UpdateOrderProductsInput	true	"data body"
////	@Success		200		{object}	models.OrderProducts
////	@Failure		400,409	{object}	Response
////	@Failure		500		{object}	Response
////	@Router			/v1/order-products/{id} [PUT]
//func (h *Handler) OrderProductUpdate(c *gin.Context) {
//	var body models.UpdateOrderProductsInput
//
//	inputID := c.Param("id")
//
//	ID, err := strconv.ParseUint(inputID, 10, 64)
//	if inputID == "" || err != nil {
//		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
//		h.log.Error("error while get query from url OrderProductUpdate", logger.Error(err))
//		return
//	}
//
//	err = c.ShouldBindJSON(&body)
//	if err != nil || inputID == "" {
//		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
//		h.log.Error("error while bind to json OrderProductUpdate", logger.Error(err))
//		return
//	}
//
//	body.ID = ID
//
//	orderProduct, err := h.services.OrderProducts.Update(c.Request.Context(), body)
//	if err != nil {
//		newResponse(c, http.StatusInternalServerError, err.Error())
//		h.log.Error("error while update orderProduct", logger.Error(err))
//		return
//	}
//
//	c.JSON(http.StatusOK, &orderProduct)
//}
//
//// OrderProductGetByID
////	@Summary		Get orderProduct by id.
////	@Description	This API to get orderProduct by id.
////	@Tags			OrderProducts
////	@Accept			json
////	@Produce		json
////	@Param			id		path		int	true	"id for get OrderProduct"	Format(id)
////	@Success		200		{object}	models.OrderProducts
////	@Failure		400,404	{object}	Response
////	@Failure		500		{object}	Response
////	@Router			/v1/order-products/{id} [GET]
//func (h *Handler) OrderProductGetByID(c *gin.Context) {
//	inputID := c.Param("id")
//
//	ID, err := strconv.ParseUint(inputID, 10, 64)
//	if inputID == "" || err != nil {
//		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
//		h.log.Error("error while get query from url OrderProductGetByID", logger.Error(err))
//		return
//	}
//
//	orderProduct, err := h.services.OrderProducts.GetByID(c.Request.Context(), ID)
//	if err != nil {
//		c.JSON(http.StatusOK, models.EmptyStruct{})
//		return
//	}
//
//	c.JSON(http.StatusOK, orderProduct)
//}
//
//// OrderProductDelete this api deletes OrderProduct
////	@Summary		Delete a orderProduct.
////	@Description	This API to delete an orderProduct.
////	@Tags			OrderProducts
////	@Accept			json
////	@Produce		json
////	@Param			id	path		int	true	"id for delete OrderProduct"	Format(id)
////	@Success		200	{object}	Response
////	@Router			/v1/order-products/{id} [DELETE]
//func (h *Handler) OrderProductDelete(c *gin.Context) {
//	inputID := c.Param("id")
//
//	ID, err := strconv.ParseUint(inputID, 10, 64)
//	if inputID == "" || err != nil {
//		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
//		h.log.Error("error while get query from url OrderProductDelete", logger.Error(err))
//		return
//	}
//
//	err = h.services.OrderProducts.DeleteByID(c.Request.Context(), ID)
//	if err != nil {
//		newResponse(c, http.StatusInternalServerError, err.Error())
//		h.log.Error("error while delete OrderProductDelete by id", logger.Error(err))
//		return
//	}
//
//	c.JSON(http.StatusOK, Response{Message: "success"})
//}
