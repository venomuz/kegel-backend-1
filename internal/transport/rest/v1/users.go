package v1

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	{
		api.POST("users/login", h.UserLogin)
		api.POST("users/logout", h.UserLogout)
		api.GET("users/auth-check", h.UserAuthCheck)
	}
}

// UserLogin
//	@Summary		Login user
//	@Description	This API to login user.
//	@Tags			Users
//	@Accept			mpfd
//	@Produce		json
//	@Param			data		formData	models.LoginUserInput	true	"data body"
//	@Success		202			{object}	models.Users
//	@Failure		400,401,404	{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/users/login [POST]
func (h *Handler) UserLogin(c *gin.Context) {

	session := sessions.Default(c)

	var body models.LoginUserInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json UserLogin", logger.Error(err))
		return
	}

	user, err := h.services.Users.Login(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, models.ErrUsernameOrPassword.Error())
		h.log.Error("error while authentication UserLogin", logger.Error(err))
		return
	}

	session.Set(h.cfg.UsersKey, user.ID)
	if err = session.Save(); err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrSaveSession.Error())
		return
	}

	c.JSON(http.StatusAccepted, user)
}

// UserLogout
//	@Summary		Logout user
//	@Description	This API to logout users.
//	@Tags			Users
//	@Accept			mpfd
//	@Produce		json
//	@Success		200		{object}	Response
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/users/logout [POST]
func (h *Handler) UserLogout(c *gin.Context) {

	session := sessions.Default(c)

	user := session.Get(h.cfg.UsersKey)
	fmt.Println(user)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(h.cfg.UsersKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// UserAuthCheck
//	@Summary		Check user authorization
//	@Description	This API to check user authorization.
//	@Tags			Users
//	@Produce		json
//	@Success		200		{object}	Response
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/users/auth-check [GET]
func (h *Handler) UserAuthCheck(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(h.cfg.UsersKey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "authorized"})
}
