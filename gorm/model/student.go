package model

type student struct {
	ID   int `gorm:"primaryKey"`	
	Name string `gorm:"nvarchar(100);not null"`
	Age  int `gorm:"not null"`
}