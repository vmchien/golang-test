package model

type Info struct {
	Phone    string  `json:"phone" gorm:"column:phone;primary_key;not_null"`
	UserName string  `json:"userName" gorm:"column:user_name"`
	Gender   string  `json:"gender" gorm:"column:gender"`
	Email    string  `json:"email" gorm:"column:email"`
	Image    string  `json:"image" gorm:"column:image"`
	Status   int     `json:"status" gorm:"column:status"`
	Amount   float64 `json:"amount" gorm:"column:amount"`
}

func (Info) TableName() string {
	return "info_user"
}
