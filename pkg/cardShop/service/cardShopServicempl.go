package service

import (
	
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
	_cardShopRepository "github.com/tehdev/summoner-rift-api/pkg/cardShop/repository"
)

type cardShopServiceImpl struct {
	cardShopRepository _cardShopRepository.CardshopRepository
}

func NewCardShopServiceImpl(
	cardShopRepository _cardShopRepository.CardshopRepository,
) CardShopService {
	return &cardShopServiceImpl{cardShopRepository}
}


func (s *cardShopServiceImpl) Listing(cardFilter *_cardShopModel.CardFilter) (*_cardShopModel.CardResult, error) {
	// 1. ดึงรายการการ์ด (entities) จาก Repository
	cardEntityList, err := s.cardShopRepository.Listing(cardFilter)
	if err != nil {
		return nil, err
	}

	// 2. นับจำนวนการ์ดทั้งหมดที่ตรงตามเงื่อนไข
	totalCards, err := s.cardShopRepository.Counting(cardFilter)
	if err != nil {
		return nil, err
	}

	// 3. คำนวณจำนวนหน้าทั้งหมด
	totalPage := s.totalPageCalculation(totalCards, cardFilter.Size)

	// 4. แปลง Entity List เป็น Model List
	cardModelList := make([]*_cardShopModel.Card, 0)
	for _, cardEntity := range cardEntityList {
		cardModelList = append(cardModelList, cardEntity.ToCardModel())
	}

	// 5. สร้างผลลัพธ์สุดท้ายที่สมบูรณ์
	return &_cardShopModel.CardResult{
		Card	: cardModelList, // **หมายเหตุ:** ตรวจสอบให้แน่ใจว่าใน struct CardResult ของคุณ field นี้ชื่อ Cards (พหูพจน์)
		Paginate: _cardShopModel.PaginateResult{
			Page:      cardFilter.Page,
			TotalPage: totalPage,
		},
	}, nil
}

// totalPageCalculation คำนวณจำนวนหน้าทั้งหมด
func (s *cardShopServiceImpl) totalPageCalculation(totalCards int64, size int64) int64 {
	// ป้องกันการหารด้วยศูนย์
    if size <= 0 {
        return 1
    }
    
    totalPage := totalCards / size
	if totalCards%size != 0 {
		totalPage++
	}
	return totalPage
}

// --- ไม่ต้องใช้ฟังก์ชัน toCardResultResponse แล้ว ---
// เราได้รวม Logic การแปลงข้อมูลเข้าไปในฟังก์ชัน Listing โดยตรงแล้ว
// ซึ่งเป็นวิธีที่สะอาดและเข้าใจง่ายกว่า