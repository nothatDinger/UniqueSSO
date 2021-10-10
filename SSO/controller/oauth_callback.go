package controller

import (
	"errors"
	"net/http"

	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/UniqueStudio/UniqueSSO/util"
	"github.com/gin-gonic/gin"
)

// for first login, append user info. otherwise just login
func LarkOauthCallbackHandler(ctx *gin.Context) {
	apmCtx, span := util.Tracer.Start(ctx.Request.Context(), "LarkOauthCallbackHandler")
	defer span.End()

	code, ok := ctx.GetQuery("code")
	if !ok {
		ctx.JSON(http.StatusBadGateway, pkg.InvalidRequest(errors.New("no code in redirect url")))
		return
	}
	// service := ctx.Query("state")

	userToken, err := util.GetLarkUserToken(apmCtx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}

	larkInfo, err := util.GetLarkUserInfo(apmCtx, userToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}


}
