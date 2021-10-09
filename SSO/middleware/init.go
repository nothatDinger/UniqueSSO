package middleware

import (
	"github.com/UniqueStudio/UniqueSSO/conf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

var (
	RedisSessionStore redis.Store
)

func SetupMiddleware() (err error) {
	if err = initRedisSessionStore(); err != nil {
		return err
	}

	return nil
}

func initRedisSessionStore() (err error) {
	RedisSessionStore, err = redis.NewStore(
		10, "tcp",
		conf.SSOConf.Redis.Addr,
		conf.SSOConf.Redis.Password,
		[]byte(conf.SSOConf.Application.SessionSecret),
	)
	if err != nil {
		return err
	}
	RedisSessionStore.Options(sessions.Options{
		Path:     "/",
		Domain:   conf.SSOConf.Application.SessionDomain,
		HttpOnly: true,
	})
	return nil
}
