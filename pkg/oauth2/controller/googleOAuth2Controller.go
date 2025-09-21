package controller

import (
	"math/rand"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/config"
	_oauth2Service "github.com/tehdev/summoner-rift-api/pkg/oauth2/service"
	"golang.org/x/oauth2"
)


type googleOAuth2Controller struct{
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf		*config.OAuth2
	logger		echo.Logger	
}

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rft"
	stateCookieName        = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)


func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2conf *config.OAuth2,
	logger 		echo.Logger,

) OAuth2Controller {
	once.Do(func()  {
		setGoogleOAuth2Config(oauth2conf)
	})

	return &googleOAuth2Controller{
		oauth2Service,
		oauth2conf,
		logger,
	}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error{
	state := c.randomState()

	c.setCookie(pctx, stateCookieName, state)
 
	return  pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error{
	state := c.randomState()

	c.setCookie(pctx, stateCookieName, state)

	return  pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error{
	panic("Implement Me")
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	panic("Implement Me")
}


func (c *googleOAuth2Controller) Logout(pctx echo.Context) error{
	c.removeCookie(pctx, accessTokenCookieName)
	c.removeCookie(pctx, refreshTokenCookieName)
	c.removeCookie(pctx, stateCookieName)

	return pctx.NoContent(http.StatusNoContent)
}


func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) error{
	cookie := &http.Cookie{
		Name: name,
		Value: value,
		Path: "/",
		HttpOnly: true,
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name: name,
		HttpOnly: true,
		MaxAge: -1,
	}

	pctx.SetCookie(cookie)
}







//คอยเช็คว่า ตัวที่ Login เข้ามาน่าไว้ใจให้ไปขอโทเค็นมากมั้ย
func (c *googleOAuth2Controller) randomState() string{
	b := make([]byte,16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return  string(b)
}



