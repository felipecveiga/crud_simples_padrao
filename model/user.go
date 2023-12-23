package model

type User struct {
	ID int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}