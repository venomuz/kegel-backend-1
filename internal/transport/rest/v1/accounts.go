package v1

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initAccountsRoutes(api *gin.RouterGroup) {
	{

		api.PUT("accounts/:id", h.AccountUpdate)
		api.GET("accounts", h.AccountsGet)
		api.GET("accounts/:id", h.CheckAccountRole, h.AccountGetByID)

		api.POST("accounts-registration", h.AccountRegistration)
		api.POST("accounts-verification", h.AccountSendVerification)

		api.POST("accounts/login", h.AccountLogin)
		api.POST("accounts/logout", h.AccountLogout)
	}
}

// AccountUpdate
//	@Summary		Update an Account
//	@Description	This api is for create Account
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint						true	"id for update Account"	Format(id)
//	@Param			data	body		models.UpdateAccountInput	true	"data body"
//	@Success		200		{object}	models.Accounts
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/{id} [PUT]
func (h *Handler) AccountUpdate(c *gin.Context) {
	var body models.UpdateAccountInput

	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url AccountUpdate", logger.Error(err))
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json AccountUpdate", logger.Error(err))
		return
	}

	body.ID = uint32(ID)

	account, err := h.services.Accounts.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, models.ErrPhoneAlreadyExists.Error())
		h.log.Error("error while update account", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &account)
}

// AccountsGet
//	@Summary		Get all Accounts
//	@Description	This api for get Accounts
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Accounts
//	@Failure		500	{object}	Response
//	@Router			/v1/accounts [GET]
func (h *Handler) AccountsGet(c *gin.Context) {
	accounts, err := h.services.Accounts.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrGetAll.Error())
		h.log.Error("error while get all AccountsGet", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// AccountGetByID
//	@Summary		Gets an Account by id
//	@Description	this api is for get Account by id
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID for get Account"	Format(id)
//	@Success		200	{object}	models.Accounts
//	@Failure		400	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/v1/accounts/{id} [GET]
func (h *Handler) AccountGetByID(c *gin.Context) {
	inputID := c.Param("id")

	id, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		h.log.Error("error while get query from url AccountGetByIDWithOrders", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.GetByID(c.Request.Context(), uint32(id))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		h.log.Error("error while get accounts AccountGetByIDWithOrders", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

// AccountRegistration
//	@Summary		Registration an account.
//	@Description	this API to registration account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.RegistrationAccountInput	true	"data body"
//	@Success		201		{object}	models.Accounts
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts-registration [POST]
func (h *Handler) AccountRegistration(c *gin.Context) {
	var body models.RegistrationAccountInput

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json AccountCreate", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.Registration(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while create account", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, account)
}

// AccountSendVerification
//	@Summary		Send verification code.
//	@Description	this API to send account's verification code to phone number.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.AccountSendVerificationInput	true	"data body"
//	@Success		200		{object}	Response
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts-verification [POST]
func (h *Handler) AccountSendVerification(c *gin.Context) {
	var body models.AccountSendVerificationInput

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json AccountSendVerification", logger.Error(err))
		return
	}

	err = h.services.Accounts.SendVerificationCode(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("error while create account", logger.Error(err))
		return
	}

	newResponse(c, http.StatusOK, "success")
}

// AccountLogin
//	@Summary		Login account
//	@Description	This API to login account.
//	@Tags			Accounts
//	@Accept			mpfd
//	@Produce		json
//	@Param			data		formData	models.LoginAccountInput	true	"data body"
//	@Success		202			{object}	models.Accounts
//	@Failure		400,401,404	{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/v1/accounts/login [POST]
func (h *Handler) AccountLogin(c *gin.Context) {
	var body models.LoginAccountInput

	session := sessions.Default(c)

	err := c.ShouldBind(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		h.log.Error("error while bind to json AccountLogin", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.Login(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, models.ErrPhoneOrPasswordWrong.Error())
		h.log.Error("error while authentication AccountLogin", logger.Error(err))
		return
	}

	session.Set(h.cfg.AccountsKey, account.ID)
	if err := session.Save(); err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrSaveSession.Error())
		return
	}

	c.JSON(http.StatusAccepted, account)
}

// AccountLogout
//	@Summary		Logout account
//	@Description	This API to logout accounts.
//	@Tags			Accounts
//	@Accept			mpfd
//	@Produce		json
//	@Success		200		{object}	Response
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/logout [POST]
func (h *Handler) AccountLogout(c *gin.Context) {
	session := sessions.Default(c)

	account := session.Get(h.cfg.AccountsKey)
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(h.cfg.AccountsKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out hello"})
}
