package model

type(
	CardCreatingReq struct{
		AdminID 		string
		Name            string 		`json:"name" validate:"required,max=64"`
		Description		string		`json:"description" validate:"required,max=128"`
		Picture			string		`json:"picture" validate:"required"`
		Price			uint		`json:"price" validate:"required"`
	}

	CardEditingReq struct{
		AdminID 		string
		Name            string 		`json:"name" validate:"required,max=64"`
		Description		string		`json:"description" validate:"required,max=128"`
		Picture			string		`json:"picture"  validate:"required"`
		Price			uint		`json:"price" validate:"required"`
	}
)