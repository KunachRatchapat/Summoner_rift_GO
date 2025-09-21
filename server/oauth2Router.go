package server

import(
	_oauth2Service "github.com/tehdev/summoner-rift-api/pkg/oauth2/service"
	_oauth2Controller "github.com/tehdev/summoner-rift-api/pkg/oauth2/controller"
	_playerRepository "github.com/tehdev/summoner-rift-api/pkg/player/repository"
	_adminRepository "github.com/tehdev/summoner-rift-api/pkg/admin/repository"
)
func (s *echoServer)initOAuth2Router(){
	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _playerRepository.NewplayerRepositorympl(s.db,s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositorympl(s.db,s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		s.app.Logger,
	)

	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin) 
	router.GET("player/login/callback",	oauth2Controller.PlayerLoginCallback)
	router.GET("admin/login/callback",oauth2Controller.AdminLoginCallback)
	router.DELETE("logout",oauth2Controller.Logout)
	

}