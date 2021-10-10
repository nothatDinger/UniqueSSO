package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/conf"
	"github.com/UniqueStudio/UniqueSSO/database"
	"github.com/UniqueStudio/UniqueSSO/model"
	"github.com/UniqueStudio/UniqueSSO/pb/lark"
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

const (
	LarkWebFmt    = "https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal"                  //获取app_token的网址
	LarkStateFmt  = "https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri=%s&app_id=%s&state=%s" //获取auth_code的网址
	LarkUserIdFmt = "https://open.larksuite.com/open-apis/authen/v1/access_token"                             //获取用户信息的网址
)

func GetLarkUserToken(ctx context.Context, authCode string) (string, error) {
	apmCtx, span := Tracer.Start(ctx, "GetLarkUserToken")
	defer span.End()

	bs, err := json.Marshal(pkg.LarkCode2TokenReq{
		GrantType: "authorization_code",
		Code:      authCode,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		apmCtx,
		http.MethodPost,
		common.LARK_AUTH_CODE2TOKEN,
		bytes.NewBuffer(bs),
	)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("new get lark user token request failed", zap.Error(err))
		return "", err
	}

	tt, err := GetLarkTenantToken(apmCtx)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+tt)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("do get lark_user_token failed", zap.Error(err))
		return "", err
	}

	data := new(pkg.LarkCode2TokenResp)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("unmarshal lark_user_token failed", zap.Error(err))
		return "", err
	}

	if data.Data.AccessToken == "" {
		err := errors.New(data.Message)
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("lark_user_token is missing", zap.Error(err))
		return "", err
	}

	return data.Data.AccessToken, nil
}

func GetLarkTenantToken(ctx context.Context) (string, error) {
	apmCtx, span := Tracer.Start(ctx, "GetLarkTenantToken")
	defer span.End()

	token, err := database.RedisClient.Get(apmCtx, common.REDIS_LARK_TENANT_TOKEN_KEY).Result()
	// meet cache
	if err == nil && token != "" {
		return token, nil
	}

	bs, err := json.Marshal(pkg.LarkTenantTokenReq{
		AppId:     conf.SSOConf.Lark.AppId,
		AppSecret: conf.SSOConf.Lark.AppSecret,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		common.LARK_TENANT_TOKEN,
		bytes.NewBuffer(bs),
	)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("new accessing lark_tenant_token request failed", zap.Error(err))
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("do accessing lark_tenant_token failed", zap.Error(err))
		return "", err
	}

	data := new(pkg.LarkTenantTokenResp)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("unmarshal lark_tenant_token response failed", zap.Error(err))
		return "", err
	}

	if data.TenantAccessToken == "" {
		err := errors.New(data.Message)
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get tenant access token failed", zap.Error(err))
		return "", err
	}

	database.RedisClient.Set(apmCtx, common.REDIS_LARK_TENANT_TOKEN_KEY, data.TenantAccessToken, time.Duration(data.Expire)*time.Second)

	return data.TenantAccessToken, nil
}

func GetLarkUserInfo(ctx context.Context, userToken string) (*lark.LarkUserInfo, error) {
	apmCtx, span := Tracer.Start(ctx, "GetLarkUserInfo")
	defer span.End()

	req, err := http.NewRequestWithContext(
		apmCtx,
		http.MethodGet,
		common.LARK_FETCH_USER_INFO,
		nil,
	)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("build fetch lark user info request failed", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+userToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("send fetch lark user info failed", zap.Error(err))
		return nil, err
	}

	user := new(lark.LarkUserInfo)
	if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("unmarshal response failed", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func UnmarshalLarkExternalInfo(buf any.Any) (*model.LarkExternalInfo, error) {
	larkExtInfo := &lark.LarkUserInfo{}
	if err := buf.UnmarshalTo(larkExtInfo); err != nil {
		return nil, err
	}
	return &model.LarkExternalInfo{
		ExternalInfo: sso.ExternalInfo{
			EName: common.EXTERNAL_NAME_LARK,
			EID:   larkExtInfo.UnionID,
		},
		LarkUserInfo: *larkExtInfo,
	}, nil
}
