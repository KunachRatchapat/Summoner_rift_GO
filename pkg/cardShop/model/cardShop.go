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
	}
)