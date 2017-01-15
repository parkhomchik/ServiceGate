package database

//Parameter
type Parameter struct {
	ID        int
	RequestID int `gorm:"index"`
	Name      string
	Value     string
}
