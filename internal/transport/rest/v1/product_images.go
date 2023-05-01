package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initProductImagesRoutes(api *gin.RouterGroup) {
	{
		api.POST("product-images", h.ProductImageCreate)
		api.PUT("product-images/:id", h.ProductImageUpdate)
		api.GET("product-images/:id", h.ProductImageGetByID)
		api.DELETE("product-images/:id", h.ProductImageDelete)
	}
}

// ProductImageCreate
//	@Summary		Create a productImage.
//	@Description	This API to create a productImage.
//	@Tags			ProductImages
//	@Accept			json
//	@Produce		json
//	@Param			data		formData	models.CreateProductImageInput	true	"data body"
//	@Param			fileImage	formData	file							true	"fileImage"
//	@Success		201			{object}	models.ProductImages
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/product-images [POST]
func (h *Handler) ProductImageCreate(c *gin.Context) {
	var body models.CreateProductImageInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductImageCreate", logger.Error(err))
		return
	}

	productImage, err := h.services.ProductImages.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		h.log.Error("error while create productImage", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, productImage)
}

// ProductImageUpdate
//	@Summary		Update a productImage.
//	@Description	This API to update a productImage.
//	@Tags			ProductImages
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int								true	"id for update ProductImage"	Format(id)
//	@Param			data		formData	models.UpdateProductImageInput	true	"data body"
//	@Param			file_image	formData	file							true	"file_image"
//	@Success		200			{object}	models.ProductImages
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/product-images/{id} [PUT]
func (h *Handler) ProductImageUpdate(c *gin.Context) {
	var body models.UpdateProductImageInput

	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url AccountUpdate", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductImageUpdate", logger.Error(err))
		return
	}

	body.ID = uint32(ID)

	productImage, err := h.services.ProductImages.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while update productImage", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &productImage)
}

// ProductImageGetByID
//	@Summary		Get productImage by id.
//	@Description	This API to get productImage by id.
//	@Tags			ProductImages
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id for get ProductImage"	Format(id)
//	@Success		200		{object}	models.ProductImages
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/product-images/{id} [GET]
func (h *Handler) ProductImageGetByID(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url AccountUpdate", logger.Error(err))
		return
	}

	productImage, err := h.services.ProductImages.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, productImage)
}

// ProductImageDelete this api deletes ProductImage
//	@Summary		Delete a productImage.
//	@Description	This API to delete a productImage.
//	@Tags			ProductImages
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id for delete ProductImage"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/product-images/{id} [DELETE]
func (h *Handler) ProductImageDelete(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url AccountUpdate", logger.Error(err))
		return
	}

	err = h.services.ProductImages.DeleteByID(c.Request.Context(), uint32(ID))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while delete productImages by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
