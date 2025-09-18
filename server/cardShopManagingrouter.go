package server

import (
	_cardManagingController "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/controller"
	_cardManagingRepository "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/repository"
	_cardManagingService 	"github.com/tehdev/summoner-rift-api/pkg/cardShopMange/service"
	_cardShopRepository		"github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"

)

func (s *echoServer) initCardManagingRouter(){
	router := s.app.Group("/v1/card-managing")

	cardShopRepository := _cardShopRepository.NewCardShpRepositoryImpl(s.db, s.app.Logger)  
	cardManagingRepository := _cardManagingRepository.NewCardManagingRepositorympl(s.db, s.app.Logger)
	cardManagingService := _cardManagingService.NewCardManagingServicempl(
		cardManagingRepository, 
		cardShopRepository,
	)
	cardManagingController := _cardManagingController.NewCardManagingControllermpl(cardManagingService)

	router.POST("", cardManagingController.Creating)
	router.PATCH("/:cardID",cardManagingController.Editing)
	router.DELETE("/:cardID",cardManagingController.Archiving)



}