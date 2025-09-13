package server 

import(
	_cardShopRepository "github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"
	_cardShopService "github.com/tehdev/summoner-rift-api/pkg/cardShop/service"
	_cardShopController "github.com/tehdev/summoner-rift-api/pkg/cardShop/controller"
 )

func(s *echoServer) initCardShopRouter() {
	router := s.app.Group("/v1/card-shop")

	cardShopRepository := _cardShopRepository.NewCardShpRepositoryImpl(s.db, s.app.Logger)
	cardShopservice := _cardShopService.NewCardShopServiceImpl(cardShopRepository)
	cardShopController := _cardShopController.NewCardShopControllerImpl(cardShopservice)

	router.GET("",cardShopController.Listing) //อย่าลืมประกาศให้ Server รู้ว่าประกาศ router นี้ไว้
}