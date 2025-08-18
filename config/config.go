package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type (
	Config struct{
		Server 		*Server			`mapstructure:"server" validate:"required"`
		OAuth2		*OAuth2			`mapstructure:"OAuth2" validate:"required"`
		State		*State			`mapstructure:"state" validate:"required"`
		Database	*Database		`mapstructure:"database" validate:"required"`
	}

	Server struct{
		Port			int			`mapstructure:"port" validate:"required"`
		AllowedOrigins	[]string 	`mapstructure:"allowOrigins" validate:"required"`
		BodyLimit		string		`mapstructure:"bodyLimit" validate:"required"`
		TimeOut			int			`mapstructure:"timeout" validate:"required"`
	}

	OAuth2 struct{
		PlayerRedirectUrl		string		`mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl		string		`mapstructure:"adminRedirectUrl" validate:"required"`
		ClientID				string		`mapstructure:"clientId" validate:"required"`
		ClientSecret			string		`mapstructure:"clientSecret" validate:"required"`
		Endpoints				endpoint	`mapstructure:"endpoints" validate:"required"`
		Scopes					[]string	`mapstructure:"scopes" validate:"required"`
		UserInfoUrl				string		`mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl				string		`mapstructure:"revokeUrl" validate:"required"`
	}

	endpoint struct{
		AuthUrl					string		`mapstructure:"authUrl" validate:"required"`
		TokenUrl				string		`mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl			string		`mapstructure:"deviceAuthUrl" validate:"required"`
	}

	State struct{
		Secret		string			`mapstructure:"secret" validate:"required"`
		ExpiresAt	time.Duration	`mapstructure:"expiresAt" validate:"required"`
		Issuer		string			`mapstructure:"issuer" validate:"required"`	
	}
	
	Database struct{
		Host		string		`mapstructure:"host" validate:"required"`
		Port		int			`mapstructure:"port" validate:"required"`
		User        string		`mapstructure:"user" validate:"required"`
		Password	int			`mapstructure:"password" validate:"required"`
		DBname		string		`mapstructure:"dbname" validate:"required"`
		SSLMode		string		`mapstructure:"sslmode" validate:"required"`
		Schema		string		`mapstructure:"schema" validate:"required"`

	}
)

//Instance
var (
	once 		sync.Once 	//utility ของ Go เอาไว้ให้ฟังก์ชันบางอย่างรันครั้งเดียวเท่านั้น 
	configInstance *Config	//ไว้เก็บค่า config ที่เก็บมา
)

//Function For Config
func ConfigGetting()  *Config {

	//ให้ code run ครั้งเดียวไม่โหลด Config ซ้ำ
	once.Do(func() {   

		//คือบอก viper ให้ใช้ไฟล์ config ชื่อ config.(yaml/json/toml)
		viper.SetConfigName("config") 
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		// ให้ viper อ่านค่าจาก ENV ด้วย (เช่นถ้า deploy ขึ้น server)
		viper.AutomaticEnv()  //อ่านไฟล์ ENV พวกดาต้าเบสไรงี้
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// อ่านไฟล์ config.yaml
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		// แปลงค่าที่อ่านได้จาก viper ไปเก็บใน struct Config
		if err := viper.Unmarshal(&configInstance); err != nil{ //จะ map key → field ของ struct Config ที่เราประกาศไว้
			panic(err)
		}

		// ใช้ validator เช็คว่าค่า config ถูกต้องตาม struct ไหม
		//ใช้ validator เพื่อเช็คว่า field ที่ต้องมี (เช่น required) มีค่าครบไหม
		validating := validator.New()
		if err := validating.Struct(configInstance); err != nil{
			panic(err)
		}

	})
	return configInstance
}
