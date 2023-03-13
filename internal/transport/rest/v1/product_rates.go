package v1

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
)

func (h *Handler) initProductRatesRoutes(api *gin.RouterGroup) {
	{
		api.POST("product-rates", h.ProductRateCreate)
	}
}

// ProductRateCreate
//	@Summary		create a productRate
//	@Description	This API to create a productRate.
//	@Tags			ProductRates
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.CreateProductRateInput	true	"data body"
//	@Success		201		{object}	models.ProductRates
//	@Failure		400,401	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/product-rates [POST]
func (h *Handler) ProductRateCreate(c *gin.Context) {

	session := sessions.Default(c)

	var body models.CreateProductRateInput
	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductRateCreate", logger.Error(err))
		return
	}

	id := session.Get(h.cfg.AccountsKey).(uint32)

	body.AccountID = id

	productRate, err := h.services.ProductRates.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while create ProductRates", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, productRate)
}
