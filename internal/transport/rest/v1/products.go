package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	{
		api.POST("products", h.ProductCreate)
		api.PUT("products/:id", h.ProductUpdate)
		api.GET("products-by-filter", h.ProductsWithImagesGetByFilter)
		api.POST("products-by-ids", h.ProductsGetByIDs)
		api.GET("products/:id", h.ProductWithImagesGetByID)
		api.GET("products-url/:url", h.ProductGetByUrl)
		//api.GET("products/groups/:url", h.ProductsWithImagesGetByFilterAndGroupUrl)
		api.DELETE("products/:id", h.ProductDelete)
	}
}

// ProductCreate
//	@Summary		Create a product.
//	@Description	This API to create a product.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.CreateProductInput	true	"data body"
//	@Success		201		{object}	models.Products
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/products [POST]
func (h *Handler) ProductCreate(c *gin.Context) {
	var body models.CreateProductInput

	err := c.ShouldBindJSON(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		h.log.Error("error while bind to json ProductCreate", logger.Error(err))
		return
	}

	product, err := h.services.Products.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		h.log.Error("error while create product", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, product)
}

// ProductUpdate
//	@Summary		Update a product.
//	@Description	This API to update a product.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"id for update Product"	Format(id)
//	@Param			data	body		models.UpdateProductInput	true	"data body"
//	@Success		200		{object}	models.Products
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/products/{id} [PUT]
func (h *Handler) ProductUpdate(c *gin.Context) {
	var body models.UpdateProductInput

	inputID := c.Param("id")
	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url ProductUpdate", logger.Error(models.ErrNotFoundId))
		return
	}

	err := c.ShouldBindJSON(&body)
	if err != nil || inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductUpdate", logger.Error(err))
		return
	}

	body.ID = inputID

	product, err := h.services.Products.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while update product", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &product)
}

//// ProductsWithImagesGetByFilterAndGroupUrl
//// @Summary     Get product with images by group url.
//// @Description This API to get product with images by group url.
//// @Tags        Products
//// @Accept      json
//// @Produce     json
//// @Param       url      path     string true "url for get Product" Format(url)
//// @Success     200     {object} []models.ProductWithImages
//// @Failure     400,404 {object} Response
//// @Failure     500     {object} Response
//// @Router      /v1/products/groups/{url} [GET]
//func (h *Handler) ProductsWithImagesGetByFilterAndGroupUrl(c *gin.Context) {
//	inputID := c.Param("url")
//	fmt.Println(inputID)
//	if inputID == "" {
//		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
//		h.log.Error("error while get param from url ProductsWithImagesGetByFilterAndGroupUrl", logger.Error(models.ErrNotFoundId))
//		return
//	}
//
//	product, err := h.services.Products.GetAllByGroupUrlWithImages(c.Request.Context(), inputID)
//	if err != nil {
//		c.JSON(http.StatusOK, models.EmptyStruct{})
//		return
//	}
//
//	c.JSON(http.StatusOK, product)
//}

// ProductsWithImagesGetByFilter
//	@Summary		Get products with images by filter.
//	@Description	This API to get products with images by filter.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			data	query		models.GetProductsByFilterInput	true	"data body"
//	@Success		200		{object}	models.ProductsWithImagesAndPagination
//	@Failure		500		{object}	Response
//	@Router			/v1/products-by-filter [GET]
func (h *Handler) ProductsWithImagesGetByFilter(c *gin.Context) {

	var body models.GetProductsByFilterInput

	err := c.ShouldBindQuery(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductsGetWithFilter", logger.Error(err))
		return
	}

	products, err := h.services.Products.GetAllWithImagesByFilter(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrGetAll.Error())
		h.log.Error("error while get all products", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, products)
}

// ProductsGetByIDs
//	@Summary		Get products with images by ids.
//	@Description	This API to get products with images by ids.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.GetProductsByIDsInput	true	"data body"
//	@Success		200		{object}	[]models.ProductWithImages
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/products-by-ids [POST]
func (h *Handler) ProductsGetByIDs(c *gin.Context) {
	var body models.GetProductsByIDsInput

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json ProductsGetByIDs", logger.Error(err))
		return
	}

	product, err := h.services.Products.GetAllByIDs(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, product)
}

// ProductWithImagesGetByID
//	@Summary		Get product with images by id.
//	@Description	This API to get product with images by id.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"id for get Product"	Format(id)
//	@Success		200		{object}	models.ProductWithImages
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/products/{id} [GET]
func (h *Handler) ProductWithImagesGetByID(c *gin.Context) {
	inputID := c.Param("id")

	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url ProductWithImagesGetByID", logger.Error(models.ErrNotFoundId))
		return
	}

	product, err := h.services.Products.GetByIDWithImagesAndRates(c.Request.Context(), inputID)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, product)
}

// ProductGetByUrl
//	@Summary		Get product by url.
//	@Description	This API to get product by url.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			url	path		string	true	"url for get Product"	Format(id)
//	@Success		200	{object}	models.ProductWithImages
//	@Failure		400	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/v1/products-url/{url} [GET]
func (h *Handler) ProductGetByUrl(c *gin.Context) {
	url := c.Param("url")

	if url == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url ProductGetByUrl", logger.Error(models.ErrNotFoundId))
		return
	}

	product, err := h.services.Products.GetByUrlWithImagesAndRates(c.Request.Context(), url)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, product)
}

// ProductDelete this api deletes Product
//	@Summary		Delete a product.
//	@Description	This API to delete a product.
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id for delete Product"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/products/{id} [DELETE]
func (h *Handler) ProductDelete(c *gin.Context) {
	inputID := c.Param("id")

	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url ProductUpdate", logger.Error(models.ErrNotFoundId))
		return
	}

	err := h.services.Products.DeleteByID(c.Request.Context(), inputID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while delete product with images", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
