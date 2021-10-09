package service

import (
	"fmt"

	"github.com/UniqueStudio/UniqueSSO/common"
	"github.com/UniqueStudio/UniqueSSO/conf"
)

func GeneateLarkRedirectUrl(service string) string {
	return fmt.Sprintf(common.LARK_OAUTH,
		conf.SSOConf.Lark.RedirectUri,
		conf.SSOConf.Lark.AppId,
		service,
	)
}
