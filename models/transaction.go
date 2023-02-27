package models

type Transaction struct {
	Base
	Amount  float64 `json:"amount" gorm:"Column:amount"`
	Type    string  `json:"type" gorm:"Column:type"`
	Balance float64 `json:"balance" gorm:"Column:balance"`
}
