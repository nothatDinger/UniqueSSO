package controller

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/UniqueStudio/UniqueSSO/service"
	"github.com/UniqueStudio/UniqueSSO/util"

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

	target := &url.URL{
		Path: "/",
	}
	if redirectUrl, ok := ctx.GetQuery("service"); ok && redirectUrl != "" {
		if service.VerifyService(redirectUrl) != nil {
			ctx.JSON(http.StatusUnauthorized, pkg.InvalidService(errors.New("unsupported service: "+redirectUrl)))
			return
		}
		ru, err := url.Parse(redirectUrl)
		if err != nil {
			zapx.WithContext(apmCtx).Error("failed to parse redirect url", zap.String("service", redirectUrl))
			ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("service格式错误")))
			return
		}
		target = ru
	}

	// judge oauth type first
	switch signType {
	case common.SignTypeLark:
		ctx.Redirect(http.StatusFound, service.GeneateLarkRedirectUrl(target.String()))
		return
	}

	data := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(data); err != nil {
		zapx.WithContext(apmCtx).Error("post body format incorroct", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("上传数据格式错误")))
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

	ctx.Redirect(http.StatusFound, target.String())
}

// TODO: construct a watcher to implement logout function
func Logout(ctx *gin.Context) {

}
