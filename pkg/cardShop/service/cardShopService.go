package service
import(
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
)
type CardShopService interface {
	Listing(cardFilter *_cardShopModel.CardFilter) ([]*_cardShopModel.Card, error)
}