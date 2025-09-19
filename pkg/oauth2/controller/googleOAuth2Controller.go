package controller

import (


	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/config"
	_oauth2Service "github.com/tehdev/summoner-rift-api/pkg/oauth2/service"
)



type googleOAuth2Controller struct{
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf		*config.OAuth2
	logger		echo.Logger	
}

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2conf *config.OAuth2,
	logger 		echo.Logger,

) OAuth2Controller {
	return &googleOAuth2Controller{
		oauth2Service,
		oauth2conf,
		logger,
	}
}

