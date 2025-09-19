package server

import (
	_playerCoinController "github.com/tehdev/summoner-rift-api/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/tehdev/summoner-rift-api/pkg/playerCoin/repository"
	_playerCoinService 	  "github.com/tehdev/summoner-rift-api/pkg/playerCoin/service"
)
func (s *echoServer) initPlayerCoinRouter() {
	router := s.app.Group("v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositorympl(s.db,s.app.Logger)
	playerCoinService 	 := _playerCoinService.NewPlayerCoinServicempl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllermpl(playerCoinService)

	router.POST("",playerCoinController.CoinAdding)
	router.GET("",playerCoinController.Showing)
}