package service

import (
	"context"
	"errors"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/UniqueStudio/UniqueSSO/repo"
	"github.com/UniqueStudio/UniqueSSO/util"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

func VerifyUser(ctx context.Context, login *pkg.LoginUser, signType string) (*sso.User, error) {
	switch signType {
	case common.SignTypeEmailPassword:
		return VerifyUserByEmail(ctx, login.Email, login.Password, login.TOTPPasscode)
	case common.SignTypePhonePassword:
		return VerifyUserByPhone(ctx, login.Phone, login.Password, login.TOTPPasscode)
	case common.SignTypePhoneSms:
		return VerifyUserBySMS(ctx, login.Phone, login.Code)
	default:
		return nil, errors.New("Invalid sign type")
	}
}

func VerifyUserByEmail(ctx context.Context, email, password, totpPasscode string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "VerifyUserByEmail")
	defer span.End()

	user, err := repo.GetBasicUserByEmail(apmCtx, email)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get user by email failed", zap.Error(err))
		return nil, err
	}

	if !util.ValidateTOTPPasscode(totpPasscode, user.TOTPSecret) {
		err := errors.New("validate totp passcode failed")
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("validate totp passcode failed")
		return nil, err
	}

	if err := util.ValidatePassword(password, user.Phone); err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("validate password failed", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func VerifyUserByPhone(ctx context.Context, phone, password, totpPasscode string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "VerifyUserByPhone")
	defer span.End()

	user, err := repo.GetBasicUserByPhone(apmCtx, phone)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get user by phone failed", zap.Error(err))
		return nil, err
	}

	if !util.ValidateTOTPPasscode(totpPasscode, user.TOTPSecret) {
		err := errors.New("validate totp passcode failed")
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("validate totp passcode failed")
		return nil, err
	}

	if err := util.ValidatePassword(password, user.Password); err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("validate password failed", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func VerifyUserBySMS(ctx context.Context, phone, sms string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "VerifyUserBySMS")
	defer span.End()

	code, err := util.GetSMSCodeByPhone(ctx, phone)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get sms code by phone failed", zap.Error(err))
		return nil, err
	}
	if code != sms {
		zapx.WithContext(apmCtx).Error("wrong sms code")
		return nil, errors.New("sms code is wrong")
	}

	user, err := repo.GetBasicUserByPhone(apmCtx, phone)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get user by phone failed")
		return nil, err
	}

	return user, nil
}
