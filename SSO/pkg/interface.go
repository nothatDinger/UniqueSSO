package pkg

import (
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
)

type UserKicker interface {
	RemoveUser(*sso.User) error
}

type UserMaintainer interface {
	RegisterKicker(UserKicker) error
	SaveUser(*sso.User) error
	SaveCustomUser(*sso.User, *sso.Permission) error
	RemoveUser(*sso.User) error
}

type PermissionController interface {
	AddPermission(*sso.User, []*sso.Permission) error
	DeletePermission(*sso.User, []*sso.Permission) error
}
