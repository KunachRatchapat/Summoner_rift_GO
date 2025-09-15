package model

type(
	Card struct {
		ID 				uint64 			`json:"id"`
		Name			string			`json:"name"`
		Description		string			`json:"description"`
		Picture			string			`json:"picture"`
		Price			int 			`json:"price"`
		
	}

	CardFilter struct{
		//ใช้ omitempty เพื่อไม่เอาค่าว่างไปแสดงงับ
		Name 					string 			`query:"name" validate:"omitempty,max=64"` 
		Description 			string 			`query:"description" validate:"omitempty,max=128"`
		Paginate
	}

	Paginate struct{
		Page 				int64  				`query:"page" validate:"required,min=1"`
		Size 				int64				`query:"size" validate:"required,min=1,max=20"`	
	}

	CardResult struct{
		Card 				[]*Card				`json:"cards"`
		Paginate			PaginateResult		`json:"paginate"`
	}

	PaginateResult struct{
		Page			int64					`json:"page"`
		TotalPage		int64					`json:"totalPage"`
	}
)