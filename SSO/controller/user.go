package controller

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/UniqueStudio/UniqueSSO/service"
	"github.com/UniqueStudio/UniqueSSO/util"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

/*
	query param:
	type: phone / sms / email / wechat
	service[option]

	1. phone number with password
    body:
    {
        "phone": "",
        "password": ""
		"totp_token": ""
    }

2. phone sms
    body:
    {
        "phone": "",
        "code": ""
    }

3. email address with password
    body:
    {
        "email": "",
        "password": ""
    }
*/

func Login(ctx *gin.Context) {
	apmCtx, span := util.Tracer.Start(ctx.Request.Context(), "Login")
	defer span.End()

	signType, ok := ctx.GetQuery("type")
	if !ok {
		zapx.WithContext(apmCtx).Error("sign type unsupported", zap.String("type", signType))
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("unsupported login type: "+signType)))
		return
	}

	redirectURI := url.PathEscape(ctx.Query("service"))

	// judge oauth type first
	switch signType {
	case common.SignTypeLark:
		ctx.Redirect(http.StatusFound, service.GeneateLarkRedirectUrl(redirectURI))
		return
	}

	data := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(data); err != nil {
		zapx.WithContext(apmCtx).Error("post body format incorroct", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(err))
		return
	}

	// validate user
	user, err := service.VerifyUser(ctx.Request.Context(), data, signType)
	if err != nil {
		zapx.WithContext(apmCtx).Error("validate user failed", zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, pkg.InvalidTicketSpec(err))
		return
	}

	// issue session
	sess := sessions.Default(ctx)
	sess.Set(common.SESSION_NAME_UID, user.UID)
	sess.Save()

	if redirectURI != "" {
		ctx.Redirect(http.StatusFound, redirectURI)
		return
	}

	ctx.JSON(http.StatusOK, pkg.AuthSuccess(user))
}

func Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Delete(common.SESSION_NAME_UID)
}
