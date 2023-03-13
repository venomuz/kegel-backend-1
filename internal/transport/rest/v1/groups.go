package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
)

func (h *Handler) initGroupsRoutes(api *gin.RouterGroup) {
	{
		api.POST("groups", h.GroupCreate)
		api.PUT("groups/:id", h.GroupUpdate)
		api.GET("groups-by-filter", h.GroupsGetByFilter)
		api.GET("groups/:id", h.GroupGetByID)
		api.GET("groups-url/:url", h.GroupGetByUrl)
		api.DELETE("groups/:id", h.GroupDelete)
	}
}

// GroupCreate
//	@Summary		Create a group.
//	@Description	This API to create a group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			data		formData	models.CreateGroupInput	true	"data body"
//	@Param			fileImage	formData	file					false	"fileImage"
//	@Success		201			{object}	models.Groups
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/groups [POST]
func (h *Handler) GroupCreate(c *gin.Context) {
	var body models.CreateGroupInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json GroupCreate", logger.Error(err))
		return
	}

	group, err := h.services.Groups.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		h.log.Error("error while create group", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, group)
}

// GroupUpdate
//	@Summary		Update a group.
//	@Description	This API to update a group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string					true	"id for update Group"	Format(id)
//	@Param			data		formData	models.UpdateGroupInput	true	"data body"
//	@Param			file_image	formData	file					false	"file_image"
//	@Success		200			{object}	models.Groups
//	@Failure		400,409		{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/groups/{id} [PUT]
func (h *Handler) GroupUpdate(c *gin.Context) {
	var body models.UpdateGroupInput

	inputID := c.Param("id")
	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url GroupUpdate", logger.Error(models.ErrNotFoundId))
		return
	}

	err := c.ShouldBind(&body)
	if err != nil || inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json GroupUpdate", logger.Error(err))
		return
	}

	body.ID = inputID

	group, err := h.services.Groups.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while update group", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &group)
}

// GroupsGetByFilter
//	@Summary		Get groups by filter.
//	@Description	This API to get groups by filter.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			data	query		models.GetGroupsByFilterInput	true	"data body"
//	@Success		200		{object}	[]models.GroupsWithChild
//	@Failure		500		{object}	Response
//	@Router			/v1/groups-by-filter [GET]
func (h *Handler) GroupsGetByFilter(c *gin.Context) {

	var body models.GetGroupsByFilterInput

	err := c.ShouldBindQuery(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json GroupsGetWithFilter", logger.Error(err))
		return
	}

	groups, err := h.services.Groups.GetAllByFilterWithChild(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while get all groups", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GroupGetByID
//	@Summary		Get group by id.
//	@Description	This API to get group by id.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"id for get Group"	Format(id)
//	@Success		200		{object}	models.Groups
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/groups/{id} [GET]
func (h *Handler) GroupGetByID(c *gin.Context) {
	inputID := c.Param("id")

	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url GroupGetByID", logger.Error(models.ErrNotFoundId))
		return
	}

	group, err := h.services.Groups.GetByID(c.Request.Context(), inputID)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GroupGetByUrl
//	@Summary		Get group by url.
//	@Description	This API to get group by url.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			url	path		string	true	"url for get Group"	Format(id)
//	@Success		200	{object}	models.Groups
//	@Failure		400	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/v1/groups-url/{url} [GET]
func (h *Handler) GroupGetByUrl(c *gin.Context) {
	inputID := c.Param("url")

	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url GroupGetByUrl", logger.Error(models.ErrNotFoundId))
		return
	}

	group, err := h.services.Groups.GetByUrl(c.Request.Context(), inputID)
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GroupDelete this api deletes Group
//	@Summary		Delete a group.
//	@Description	This API to delete a group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id for delete Group"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/groups/{id} [DELETE]
func (h *Handler) GroupDelete(c *gin.Context) {
	inputID := c.Param("id")

	if inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get param from url GroupUpdate", logger.Error(models.ErrNotFoundId))
		return
	}

	err := h.services.Groups.DeleteByID(c.Request.Context(), inputID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while delete groups by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
