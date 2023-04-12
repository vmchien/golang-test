package model

type Purchase struct {
	Id     string  `json:"id" gorm:"column:phone;primary_key;auto_increment:1;not_null"`
	Phone  string  `json:"phone" gorm:"column:phone;"`
	Amount float64 `json:"amount" gorm:"column:amount"`
	Extra  string  `json:"extra" gorm:"column:extra"`
}

func (Purchase) TableName() string {
	return "purchase"
}
