package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductID int `json:"productId"`
	Quantity int `json:"quantity"`
}