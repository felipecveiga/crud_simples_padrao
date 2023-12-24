package model

type User struct {
	ID int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Idade int `json:"idade"`
	Email string `json:"email"`
}