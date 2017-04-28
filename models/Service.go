package api

//Service основно класс
type Service struct {
	ID 		int32 		`json:"id"`
	Name 	string		`json:"name"`
	Enable	bool		`json:"enable"`
}