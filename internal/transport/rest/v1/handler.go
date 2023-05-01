package v1

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/kegel-backend/config"
	"github.com/venomuz/kegel-backend/internal/service"
	"github.com/venomuz/kegel-backend/internal/storage/rdb"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"net/http"
)

type Handler struct {
	services *service.Services
	rdbRepos rdb.Repository
	log      logger.Logger
	cfg      config.Config
}

func NewHandler(services *service.Services, rdbRepos rdb.Repository, log logger.Logger, cfg config.Config) *Handler {
	return &Handler{
		services: services,
		rdbRepos: rdbRepos,
		log:      log,
		cfg:      cfg,
	}
}
func (h *Handler) Init(v1 *gin.RouterGroup) {
	{
		h.initAccountsRoutes(v1)
		h.initBannersRoutes(v1)
		h.initGroupsRoutes(v1)
		h.initOrdersRoutes(v1)
		h.initProductImagesRoutes(v1)
		h.initProductRatesRoutes(v1)
		h.initProductsRoutes(v1)
		h.initSettingsRoutes(v1)
		h.initUsersRoutes(v1)
	}

}

func (h *Handler) CheckUserRole(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(h.cfg.UsersKey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}

func (h *Handler) CheckAccountRole(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get(h.cfg.AccountsKey)
	if account == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}
