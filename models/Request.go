package models

//Request запрос на выполнение
type Request struct {
	ID          int64
	ExternalID  int64
	PortalName  string
	ServiceName string
	CommandName string
	Parameters  []Parameter
}
