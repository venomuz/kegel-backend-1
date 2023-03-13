package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	Environment    string // develop, staging, production
	LogLevel       string
	CtxTimeout     int
	StaticFilePath string
	HTTPHost       string
	HTTPPort       string
	MysqlHost      string
	MysqlPort      int
	MysqlUser      string
	MysqlPassword  string
	MysqlDatabase  string
	RedisHost      string
	RedisPort      int
	RedisPassword  string
	UserSecret     string
	UsersKey       string
	AccountSecret  string
	AccountsKey    string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.StaticFilePath = cast.ToString(getOrReturnDefault("STATIC_FILE_PATH", "/root/public/"))

	c.HTTPHost = cast.ToString(getOrReturnDefault("HTTP_HOST", "142.93.237.244"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "9090"))

	c.MysqlHost = cast.ToString(getOrReturnDefault("MARIADB_HOST", "127.0.0.1"))
	c.MysqlPort = cast.ToInt(getOrReturnDefault("MARIADB_PORT", 3306))
	c.MysqlUser = cast.ToString(getOrReturnDefault("MARIADB_USER", "kilogram"))
	c.MysqlPassword = cast.ToString(getOrReturnDefault("MARIADB_PASSWORD", "112233"))
	c.MysqlDatabase = cast.ToString(getOrReturnDefault("MARIADB_DATABASE", "kegel"))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "127.0.0.1"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "112233"))

	c.UserSecret = cast.ToString(getOrReturnDefault("USER_SECRET", "s5v8y/B?E(H+MbQeThWmZq3t6w9z$C&F)J@NcRfUjXn2r5u7x!A%D*G-KaPdSgVk"))
	c.UsersKey = cast.ToString(getOrReturnDefault("USERS_KEY", "user"))

	c.AccountSecret = cast.ToString(getOrReturnDefault("ACCOUNT_SECRET", "36o81m_IL&hTA_DEZ`|C(zk=G^(E@.xnI%dpepH&dLv?e3m5%oK8v}It8M{]69n"))
	c.AccountsKey = cast.ToString(getOrReturnDefault("ACCOUNTS_KEY", "account"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
