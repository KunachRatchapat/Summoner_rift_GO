package service

import (
							"github.com/tehdev/summoner-rift-api/entities"
	_cardShopModel 			"github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
	_cardManagingModel 		"github.com/tehdev/summoner-rift-api/pkg/cardShopMange/model"
	_cardManagingRepository "github.com/tehdev/summoner-rift-api/pkg/cardShopMange/repository"
	_cardShopRepository 	"github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"
)

type cardManagingServicempl struct{
	cardManagingRepository _cardManagingRepository.CardManagingRepository
	cardShopRepository	  _cardShopRepository.CardshopRepository

}

func NewCardManagingServicempl (
	cardManagingRepository _cardManagingRepository.CardManagingRepository,
	cardShopRepository   _cardShopRepository.CardshopRepository,

) CardManagingService{
	return  &cardManagingServicempl{
		cardManagingRepository,
		cardShopRepository,
	}

}

//รับ Request การ์ดสร้างการ์ด
func (s *cardManagingServicempl) Creating(cardCreatingReq *_cardManagingModel.CardCreatingReq) (*_cardShopModel.Card, error) {

	//แปลงข้อมูลเพื่อให้เข้าไปสู่ Repository
	cardEntity := &entities.Card{
		Name: 					cardCreatingReq.Name,
		Description: 			cardCreatingReq.Description,
		Picture: 				cardCreatingReq.Picture,
		Price:  				cardCreatingReq.Price,						
	}
	//ส่งไปให้ Repo จัดการสร้าง
	cardEntityResult, err := s.cardManagingRepository.Creating(cardEntity)
	if err != nil{
		return nil, err
	}

	return  cardEntityResult.ToCardModel(), nil
}


//แก้การ์ด
func (s *cardManagingServicempl) Editing(cardID  uint64, cardEditingReq *_cardManagingModel.CardEditingReq)(*_cardShopModel.Card, error) {
	_ , err := s.cardManagingRepository.Editing(cardID,cardEditingReq)
	if err != nil{
		return  nil, err
	}

	cardEntityResult, err := s.cardShopRepository.FindByID(cardID)
	if err != nil{
		return  nil, err
	}

	return cardEntityResult.ToCardModel(), nil
}

func (s *cardManagingServicempl) Archiving(cardID uint64) error{
	return  s.cardManagingRepository.Archiving(cardID)
}


