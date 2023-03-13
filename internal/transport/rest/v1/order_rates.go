package v1

//func (h *Handler) initRatesRoutes(api *gin.RouterGroup) {
//	{
//		api.POST("rates", h.RateCreate)
//	}
//}
//
//// RateCreate
//// @Summary     Create a rate.
//// @Description This API to create an rate.
//// @Tags        Rates
//// @Accept      json
//// @Produce     json
//// @Param       data      body     models.CreateRateInput true  "data body"
//// @Success     201       {object} models.Rates
//// @Failure     400,409   {object} Response
//// @Failure     500       {object} Response
//// @Router      /v1/rates [POST]
//func (h *Handler) RateCreate(c *gin.Context) {
//	var body models.CreateRateInput
//
//	err := c.ShouldBindJSON(&body)
//
//	if err != nil {
//		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
//		h.log.Error("error while bind to json RateCreate", logger.Error(err))
//		return
//	}
//
//	rate, err := h.services.Rates.Create(c.Request.Context(), body)
//	if err != nil {
//		newResponse(c, http.StatusConflict, err.Error())
//		h.log.Error("error while create rate", logger.Error(err))
//		return
//	}
//
//	c.JSON(http.StatusCreated, rate)
//}
