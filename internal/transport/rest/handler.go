package rest

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/venomuz/kegel-backend/config"
	"github.com/venomuz/kegel-backend/docs"
	"github.com/venomuz/kegel-backend/internal/service"
	"github.com/venomuz/kegel-backend/internal/storage/rdb"
	v1 "github.com/venomuz/kegel-backend/internal/transport/rest/v1"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"github.com/venomuz/kegel-backend/pkg/middleware"
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

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.yaml",
		SpecPath:    "/redoc/swagger.yaml",
		DocsPath:    "/redoc",
	}

	store := sessions.NewCookieStore([]byte(h.cfg.UserSecret))
	store.Options(sessions.Options{
		Path:     "/v1",
		Domain:   "",
		MaxAge:   3600 * 3,
		Secure:   false,
		HttpOnly: false,
	})

	router.Use(
		ginredoc.New(doc),
		gin.Recovery(),
		gin.Logger(),
		middleware.New(GinCorsMiddleware()),
		sessions.Sessions("kegel", store),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", h.cfg.HTTPHost, h.cfg.HTTPPort)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.rdbRepos, h.log, h.cfg)

	{
		api := router.Group("/v1")
		{
			handlerV1.Init(api)
		}
		//api.Static("/public/settings", "./public/settings")
		api.StaticFS("/public/", http.Dir(h.cfg.StaticFilePath))
	}

}
