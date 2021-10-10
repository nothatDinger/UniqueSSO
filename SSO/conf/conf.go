package conf

import (
	"net/url"

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
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	ReadTimeout   int    `mapstructure:"read_timeout"`
	WriteTimeout  int    `mapstructure:"write_timeout"`
	SessionSecret string `mapstructure:"session_secret"`
	SessionDomain string `mapstructure:"session_domain"`
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

	return nil
}
