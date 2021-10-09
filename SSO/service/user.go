package service

import (
	"context"
	"errors"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/database"
	"github.com/UniqueStudio/UniqueSSO/pkg"
	"github.com/UniqueStudio/UniqueSSO/util"

	"github.com/sirupsen/logrus"
)

func VerifyUser(ctx context.Context, login *pkg.LoginUser, signType string) (*database.User, error) {
	switch signType {
	case common.SignTypeEmailPassword:
		return VerifyUserByEmail(login.Email, login.Password)
	case common.SignTypePhonePassword:
		return VerifyUserByPhone(login.Phone, login.Password)
	case common.SignTypePhoneSms:
		return VerifyUserBySMS(ctx, login.Phone, login.Code)
	default:
		return nil, errors.New("Invalid sign type")
	}
}

func GetUserById(uid string) (*database.User, error) {
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("uid = ?", uid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserByEmail(email, password string) (*database.User, error) {
	lg := logrus.WithField("type", "email")
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("email = ?", email).Scan(user).Error
	if err != nil {
		lg.WithError(err).Info("get user by email error")
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		lg.WithError(err).Info("validate password error")
		return nil, err
	}
	return user, nil
}

func VerifyUserByPhone(phone, password string) (*database.User, error) {
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserBySMS(ctx context.Context, phone, sms string) (*database.User, error) {
	user := new(database.User)
	code, err := util.GetSMSCodeByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if code != sms {
		return nil, errors.New("sms code is wrong")
	}
	err = database.DB.Table((user.TableName())).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
