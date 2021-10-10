package repo

import (
	"context"
	"errors"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/database"
	"github.com/UniqueStudio/UniqueSSO/model"
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"github.com/UniqueStudio/UniqueSSO/util"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/xylonx/zapx"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetBasicUserById(ctx context.Context, uid string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "GetBasicUserById")
	defer span.End()

	users := make([]model.BasicUserInfo, 0, 1)
	tx := database.DB.Where(&model.BasicUserInfo{
		User: sso.User{
			UID: uid,
		},
	}).Find(&users)
	if tx.RowsAffected == 0 {
		err := tx.Error
		if err == nil {
			err = gorm.ErrRecordNotFound
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get basic user by uid failed", zap.Error(err))
		return nil, err
	}

	return &users[0].User, nil
}

func GetBasicUserByEmail(ctx context.Context, email string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "GetBasicUserByEmail")
	defer span.End()

	users := make([]model.BasicUserInfo, 0, 1)
	tx := database.DB.Where(&model.BasicUserInfo{
		User: sso.User{
			Email: email,
		},
	}).Find(&users)
	if tx.RowsAffected == 0 {
		err := tx.Error
		if err == nil {
			err = gorm.ErrRecordNotFound
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get basic user by email failed", zap.Error(err))
		return nil, err
	}

	return &users[0].User, nil
}

func GetBasicUserByPhone(ctx context.Context, phone string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "GetBasicUserByPhone")
	defer span.End()

	users := make([]model.BasicUserInfo, 0, 1)
	tx := database.DB.Where(&model.BasicUserInfo{
		User: sso.User{
			Phone: phone,
		},
	}).Find(&users)
	if tx.RowsAffected == 0 {
		err := tx.Error
		if err == nil {
			err = gorm.ErrRecordNotFound
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get basic user by phone failed", zap.Error(err))
		return nil, err
	}

	return &users[0].User, nil
}

func GetBasicUserWithGroupById(ctx context.Context, uid string) (*sso.User, error) {
	apmCtx, span := util.Tracer.Start(ctx, "GetBasicUserWithGroupById")
	defer span.End()

	tx := database.DB.Begin()
	users := make([]model.BasicUserInfo, 0, 1)
	newTx := tx.Where(&model.BasicUserInfo{
		User: sso.User{
			UID: uid,
		},
	}).Find(&users)
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = gorm.ErrRecordNotFound
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get basic user by uid failed", zap.Error(err))
		tx.Rollback()
		return nil, err
	}

	groups := make([]model.UserGroup, 0, 1)
	newTx = tx.Where(&model.UserGroup{
		UserID: uid,
	}).Find(&groups)
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = gorm.ErrRecordNotFound
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get user groups by uid failed", zap.Error(err))
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	ssogroups := make([]sso.Group, len(groups[0].Groups))
	for i := range ssogroups {
		ssogroups[i] = sso.Group(groups[0].Groups[i])
	}
	users[0].Group = ssogroups

	return &users[0].User, nil
}

func SaveUser(ctx context.Context, user *sso.User) error {
	apmCtx, span := util.Tracer.Start(ctx, "SaveUser")
	defer span.End()

	uid, err := uuid.NewUUID()
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("genrate uuid failed", zap.Error(err))
		return err
	}
	user.UID = uid.String()

	// group table data
	groups := make(pq.Int64Array, len(user.Group))
	for i := range groups {
		groups[i] = int64(user.GetGroup()[i])
	}

	// permission table data
	perms := make([]model.UserPermission, len(user.Permissions))
	for i := range perms {
		perms[i] = model.UserPermission{
			UserID:     user.UID,
			Permission: *user.Permissions[i],
		}
	}

	// lark external user info data
	// ! handle external info
	// ! FIXME: define external name in IDL
	eids := make([]string, len(user.ExternalInfos))
	extInfos := make([]interface{}, len(user.ExternalInfos))
	for i := range user.ExternalInfos {
		eids[i] = user.ExternalInfos[i].EID
		switch user.ExternalInfos[i].EName {
		case common.EXTERNAL_NAME_LARK:
			larkExtInfo, err := util.UnmarshalLarkExternalInfo(*user.ExternalInfos[i].Detail)
			if err != nil {
				span.RecordError(err)
				zapx.WithContext(apmCtx).Error("unmarshal larkExtInfo failed", zap.Error(err))
				return err
			}
			extInfos[i] = larkExtInfo

		default:
			err := errors.New("unsupported external name")
			span.RecordError(err)
			zapx.WithContext(apmCtx).Error("unsupported external name")
			return err
		}
	}

	tx := database.DB.Begin()

	//insert basic user info
	newTx := tx.Create(&model.BasicUserInfo{User: *user})
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = errors.New("no basic user info records were saved")
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("no records were saved")
		tx.Rollback()
		return err
	}

	// insert user group information
	newTx = tx.Create(&model.UserGroup{
		UserID: user.UID,
		Groups: groups,
	})
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = errors.New("no user group records were saved")
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("no records were saved")
		tx.Rollback()
		return err
	}

	// insert user permission information
	newTx = tx.Create(&perms)
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = errors.New("no user permissions records were saved")
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("no records were saved")
		tx.Rollback()
		return err
	}

	// insert external information
	for i := range extInfos {
		newTx = tx.Create(i)
		if newTx.RowsAffected == 0 {
			err := newTx.Error
			if err == nil {
				err = errors.New("no external info records were saved")
			}
			span.RecordError(err)
			zapx.WithContext(apmCtx).Error("no records were saved")
			tx.Rollback()
			return err
		}
	}

	newTx = tx.Create(&model.UserExternalInfo{
		UserID:      user.UID,
		ExternalIDs: eids,
	})
	if newTx.RowsAffected == 0 {
		err := newTx.Error
		if err == nil {
			err = errors.New("no user external id relation records were saved")
		}
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("no records were saved")
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func IsPermissionExist(ctx context.Context, uid string, perm *sso.Permission) (bool, error) {
	apmCtx, span := util.Tracer.Start(ctx, "IsPermissionExist")
	defer span.End()
	span.SetAttributes(attribute.Any("permission", perm))

	perms := make([]model.UserPermission, 0, 1)
	newTx := database.DB.Where(&model.UserPermission{
		UserID:     uid,
		Permission: *perm,
	}).Find(&perm)
	if newTx.Error != nil {
		zapx.WithContext(apmCtx).Error("database failed", zap.Error(newTx.Error))
		return false, newTx.Error
	}

	if newTx.RowsAffected == 0 {
		zapx.WithContext(apmCtx).Info("no permission in database")
		return false, nil
	}

	zapx.WithContext(apmCtx).Info("database record", zap.Any("permissions", perms))

	return true, nil
}
