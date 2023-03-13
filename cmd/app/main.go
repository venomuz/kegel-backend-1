package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/venomuz/kegel-backend/config"
	"github.com/venomuz/kegel-backend/internal/migration"
	"github.com/venomuz/kegel-backend/internal/service"
	mysqlrepo "github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/internal/storage/rdb"
	"github.com/venomuz/kegel-backend/internal/transport/rest"
	mysqldb "github.com/venomuz/kegel-backend/pkg/database/mysql"
	"github.com/venomuz/kegel-backend/pkg/gen"
	"github.com/venomuz/kegel-backend/pkg/hash"
	"github.com/venomuz/kegel-backend/pkg/humanizer"
	"github.com/venomuz/kegel-backend/pkg/logger"
)

// main
//	@title			Golang CRM Swagger Documentation
//	@version		1.0
//	@description	This is a sample server CRM server.
//	@contact.name	API Support
//	@contact.url	https://t.me/xalmatoff
//	@contact.email	venom.uzz@mail.ru
func main() {

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "kegel-service")

	humanizerUrl := humanizer.NewUrlHumanizer(map[string]string{
		"ь": "",
		"ъ": "",
	})

	hasherPassword := hash.NewPasswordHasher()

	randomManager := gen.NewRandomManager()

	db, err := mysqldb.NewClient(cfg)
	if err != nil {
		log.Error("mysqldb connection error", logger.Error(err))
		return
	}

	rdbC := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + fmt.Sprintf(":%d", cfg.RedisPort),
		Password: cfg.RedisPassword, // no password set
	})

	err = migration.AutoMigrate(db)
	if err != nil {
		log.Error("error while auto migrate", logger.Error(err))
	}

	rdbRepos := rdb.NewRedisRepo(rdbC)

	psqlRepos := mysqlrepo.NewRepositories(db)

	services := service.NewServices(service.Deps{
		MysqlRepos:   psqlRepos,
		RdbRepos:     rdbRepos,
		Log:          log,
		Cfg:          cfg,
		HumanizerUrl: humanizerUrl,
		Hasher:       hasherPassword,
		Generator:    randomManager,
	})

	handlers := rest.NewHandler(services, rdbRepos, log, cfg)

	srv := handlers.Init()
	fmt.Println("HELLO behruz")
	err = srv.Run(":" + cfg.HTTPPort)
	if err != nil {
		log.Error("rest api router running error", logger.Error(err))
		return
	}

}
