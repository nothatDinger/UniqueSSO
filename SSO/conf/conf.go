package conf

import (
	"net/url"

	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

type Conf struct {
	Application  ApplicationConf  `mapstructure:"application"`
	Database     DatabaseConf     `mapstructure:"database"`
	Redis        RedisConf        `mapstructure:"redis"`
	Sms          []SMSOptions     `mapstructure:"sms"`
	OpenPlatform OpenPlatformConf `mapstructure:"openplat_form"`
	APM          APMConf          `mapstructure:"apm"`
	Lark         LarkConf         `mapstructure:"lark"`
}
type ApplicationConf struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	RPCPort            string `mapstructure:"rpc_port"`
	RPCCertFile        string `mapstructure:"rpc_cert_file"`
	RPCKeyFile         string `mapstructure:"rpc_key_file"`
	Name               string `mapstructure:"name"`
	Mode               string `mapstructure:"mode"`
	ReadTimeout        int    `mapstructure:"read_timeout"`
	WriteTimeout       int    `mapstructure:"write_timeout"`
	SessionSecret      string `mapstructure:"session_secret"`
	SessionDomain      string `mapstructure:"session_domain"`
	TraefikRedirectURI string `mapstructure:"traefik_redirect_uri"`
}

type DatabaseConf struct {
	PostgresDSN string `mapstructure:"postgres_dsn"`
}

type RedisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type SMSOptions struct {
	Name       string `mapstructure:"name" validator:"oneof='verificationCode'"`
	TemplateId string `mapstructure:"template_id"`
	SignName   string `mapstructure:"sign_name"`
}

type OpenPlatformConf struct {
	GrpcAddr       string `mapstructure:"grpc_addr"`
	GrpcCert       string `mapstructure:"grpc_cert"`
	GrpcServerName string `mapstructure:"grpc_server_name"`
}

type APMConf struct {
	ReporterBackground string `mapstructure:"reporter_backend"`
}

type LarkConf struct {
	AppId       string `mapstructure:"app_id"`
	AppSecret   string `mapstructure:"app_secret"`
	RedirectUri string `mapstructure:"redirect_uri"`
	GroupId     struct {
		Web     string `mapstructure:"web"`
		Lab     string `mapstructure:"lab"`
		PM      string `mapstructure:"pm"`
		Design  string `mapstructure:"design"`
		Android string `mapstructure:"android"`
		IOS     string `mapstructure:"ios"`
		Game    string `mapstructure:"game"`
		AI      string `mapstructure:"ai"`
	} `mapstructure:"group_id"`
	GroupIdNameMap map[string]sso.Group `mapstructure:"-"`
	LarkUserType   struct {
		Intern    int32 `mapstructure:"intern"`
		Regular   int32 `mapstructure:"regular"`
		Graduated int32 `mapstructure:"graduated"`
	} `mapstructure:"lark_user_type"`
	LarkEnumTypeMap map[int32]sso.UserType `mapstructure:"-"`
}

var (
	SSOConf = &Conf{}
)

func InitConf(confFilepath string) error {
	viper.SetConfigFile(confFilepath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(SSOConf)
	if err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(SSOConf); err != nil {
		return err
	}

	if SSOConf.Application.Mode == "debug" {
		zapx.Info("run mode", zap.String("mode", SSOConf.Application.Mode))
	}

	SSOConf.Lark.RedirectUri = url.PathEscape(SSOConf.Lark.RedirectUri)

	// FIXME: ugly.
	SSOConf.Lark.GroupIdNameMap = make(map[string]sso.Group)
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.Web] = sso.Group_WEB
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.Lab] = sso.Group_LAB
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.PM] = sso.Group_PM
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.Design] = sso.Group_DESIGN
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.Android] = sso.Group_ANDROID
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.IOS] = sso.Group_IOS
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.Game] = sso.Group_GAME
	SSOConf.Lark.GroupIdNameMap[SSOConf.Lark.GroupId.AI] = sso.Group_AI

	SSOConf.Lark.LarkEnumTypeMap = make(map[int32]sso.UserType)
	SSOConf.Lark.LarkEnumTypeMap[SSOConf.Lark.LarkUserType.Intern] = sso.UserType_Intern
	SSOConf.Lark.LarkEnumTypeMap[SSOConf.Lark.LarkUserType.Regular] = sso.UserType_Regular
	SSOConf.Lark.LarkEnumTypeMap[SSOConf.Lark.LarkUserType.Graduated] = sso.UserType_Graduated

	return nil
}
