package controller

import (
	"net/http"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TraefikAuthValidate(ctx *gin.Context) {
	sess := sessions.Default(ctx)

	uid, ok := sess.Get(common.SESSION_NAME_UID).(string)
	if !ok {
		ctx.Redirect(http.StatusFound, conf.SSOConf.Application.TraefikRedirectURI)
		return
	}

	ctx.Writer.Header().Set("X-UID", uid)
	ctx.Status(http.StatusOK)
	return
}
