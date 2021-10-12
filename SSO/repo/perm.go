package repo

import (
	"context"

	"github.com/UniqueStudio/UniqueSSO/database"
	"github.com/UniqueStudio/UniqueSSO/model"
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"github.com/UniqueStudio/UniqueSSO/util"
	"github.com/xylonx/zapx"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

func IsPermissionExist(ctx context.Context, uid string, perm *sso.Permission) (bool, error) {
	apmCtx, span := util.Tracer.Start(ctx, "IsPermissionExist")
	defer span.End()
	span.SetAttributes(attribute.Any("permission", perm))

	dbperm := model.UserPermission{}
	// FIXME: scope should user larger
	newTx := database.DB.Where(
		"uid = ? AND action = ? AND resource = ? AND scope >= ?",
		uid, perm.Action, perm.Resource, perm.Scope).
		First(&dbperm)
	if newTx.Error != nil {
		zapx.WithContext(apmCtx).Error("database failed", zap.Error(newTx.Error))
		return false, newTx.Error
	}

	if newTx.RowsAffected == 0 {
		zapx.WithContext(apmCtx).Info("no permission in database")
		return false, nil
	}

	zapx.WithContext(apmCtx).Info("database record", zap.Any("permissions", perm))

	return true, nil
}

func DeleteAllNoReadPermission(ctx context.Context, uid string) error {
	apmCtx, span := util.Tracer.Start(ctx, "DeleteAllNoReadPermission")
	defer span.End()

	newTx := database.DB.Where("uid = ? AND action <> ?", uid, sso.Action_READ).Delete(model.UserPermission{})
	if newTx.Error != nil {
		zapx.WithContext(apmCtx).Error("database failed", zap.Error(newTx.Error))
		return newTx.Error
	}

	zapx.WithContext(apmCtx).Info("delete all no-read permission")
	return nil
}
