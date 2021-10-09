package router

import (
	"github.com/UniqueStudio/UniqueSSO/controller"
	"github.com/UniqueStudio/UniqueSSO/middleware"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// global middleware
	r.Use(middleware.TracingMiddleware())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// traefik validate
	r.GET("/gateway/validate/traefik", controller.TraefikAuthValidate)

	// sms
	smsrouter := r.Group("/sms")
	smsrouter.POST("code", controller.SendSmsCode)

	// normal login
	router := r.Group("")
	router.Use(sessions.Sessions("UserSystem", middleware.RedisSessionStore))
	router.Use(middleware.SessionRedirect())
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)

	// oauth login
	oauth := r.Group("/login/oauth")
	oauth.GET("/lark", controller.LarkOauthCallbackHandler)
}
