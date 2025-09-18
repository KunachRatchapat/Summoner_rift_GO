package model	

//เอาไว้แจ้งเตือนกรณี Login หรือ Logout ได้มั้ย

type (
	LoginResponse struct {
		Message string `json:"message"`
	}

	LogoutResponse struct {
		Message string `json:"message"`
	}
)