package repository

type  cardShopRepositoryImpl struct{}   //เป็น Struct ที่เป็น private ไว้อิมพริเม้นตัว Cardshop

func NewCardShpRepositoryImpl() CardshopRepository{
	return &cardShopRepositoryImpl{} //return this for implement
}