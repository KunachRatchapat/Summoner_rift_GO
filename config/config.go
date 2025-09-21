package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config 	struct{
		Database *Database `mapstructure:"database" validate:"required"`
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
	}

	Server struct{
		Port 			int 				`mapstructure:"port"  validata:"required"`
		AllowOrigins	[]string			`mapstructure:"allowOrigins"  validate:"required"` // API (สำหรับ CORS)
		BodyLimit		string				`mapstructure:"bodylimit"  validate:"required"`
		Timeout			time.Duration		`mapstructure:"timeout"  validate:"required"`

	}

	OAuth2 struct{
		  PlayerRedirectUrl		string			`mapstructure:"playerRedirectUrl" validate:"required"`
		  AdminRedirectUrl 		string			`mapstructure:"adminRedirectUrl" validate:"required"`
		  ClientID				string			`mapstructure:"clientId" validate:"required"` // ข้อมูลประจำตัว (credentials)
		  ClientSecret			string			`mapstructure:"clientSecret" validate:"required"`
		  Endpoints				endpoint		`mapstructure:"endpoints" validate:"required"` //URL ของ Endpoint ต่างๆ ของ OAuth2
		  Scopes				[]string		`mapstructure:"scopes" validate:"required"`
		  UserInfoUrl			string			`mapsturcture:"userInfoUrl" validate"required"`
		  RevokeUrl				string			`mapsturcture:"revokeUrl" validate"required"`
	}

	endpoint struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)	


var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()

		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}



