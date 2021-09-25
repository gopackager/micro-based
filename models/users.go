package models

type UsersCreateInput struct {
	Fullname        string `json:"fullname" form:"fullname" binding:"required"`
	NewPassword     string `json:"new_password" form:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required,email"`
}

type UsersCreateOutput struct {
	RowID    int64  `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UsersReadOutput struct {
	Users []UsersReadOutputList `json:"users"`
}

type UsersReadOutputList struct {
	RowID    int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UsersUpdateInput struct {
	RowID    int
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
}
