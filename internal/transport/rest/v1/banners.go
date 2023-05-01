package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initBannersRoutes(api *gin.RouterGroup) {
	{
		api.POST("banners", h.BannerCreate)
		api.PUT("banners/:id", h.BannerUpdate)
		api.GET("banners", h.BannersGet)
		api.GET("banners/:id", h.BannerGetByID)
		api.DELETE("banners/:id", h.BannerDelete)
	}
}

// BannerCreate
//	@Summary		Create a banner.
//	@Description	This API to create a banner.
//	@Tags			Banners
//	@Accept			json
//	@Produce		json
//	@Param			data		body		models.CreateBannerInput	true	"data body"
//	@Param			fileImage	formData	file						true	"fileImage"
//	@Success		201			{object}	models.Banners
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/banners [POST]
func (h *Handler) BannerCreate(c *gin.Context) {
	var body models.CreateBannerInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		h.log.Error("error while bind to json BannerCreate", logger.Error(err))
		return
	}

	banner, err := h.services.Banners.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		h.log.Error("error while create banner", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, banner)
}

// BannerUpdate
//	@Summary		Update a banner.
//	@Description	This API to update a banner.
//	@Tags			Banners
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int							true	"id for update Banner"	Format(id)
//	@Param			fileImage	formData	file						false	"fileImage"
//	@Param			data		body		models.UpdateBannerInput	true	"data body"
//	@Success		200			{object}	models.Banners
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/banners/{id} [PUT]
func (h *Handler) BannerUpdate(c *gin.Context) {
	var body models.UpdateBannerInput

	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil || inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json BannerUpdate", logger.Error(err))
		return
	}

	body.ID = uint32(ID)

	banner, err := h.services.Banners.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while update banner", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &banner)
}

// BannersGet
//	@Summary		Get all Banners
//	@Description	This api for get Banners
//	@Tags			Banners
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Banners
//	@Failure		500	{object}	Response
//	@Router			/v1/banners [GET]
func (h *Handler) BannersGet(c *gin.Context) {

	banners, err := h.services.Banners.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrGetAll.Error())
		h.log.Error("error while get all BannersGet", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, banners)
}

// BannerGetByID
//	@Summary		Get banner by id.
//	@Description	This API to get banner by id.
//	@Tags			Banners
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id for get Banner"	Format(id)
//	@Success		200		{object}	models.Banners
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/banners/{id} [GET]
func (h *Handler) BannerGetByID(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	banner, err := h.services.Banners.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, banner)
}

// BannerDelete this api deletes Banner
//	@Summary		Delete a banner.
//	@Description	This API to delete a banner.
//	@Tags			Banners
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id for delete Banner"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/banners/{id} [DELETE]
func (h *Handler) BannerDelete(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = h.services.Banners.DeleteByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
