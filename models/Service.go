package models

//Service основно класс
type Service struct {
	ID 		int32 		
	Name 	string		`gorm:"type:varchar(100);unique_index"`
	Enable	bool		
}