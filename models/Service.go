package models

//Service Сервис
type Service struct {
	ID   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}
