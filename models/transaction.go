package models

type Transaction struct {
	Base
	Amount string `json:"amount" gorm:"Column:amount"`
	Type   string `json:"type" gorm:"Column:type"`
}
