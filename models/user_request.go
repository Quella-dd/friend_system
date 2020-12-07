package models

type RequestManager struct {}

type Request struct {
	UserID string
	RequestID string
	Content string
}

func (m *RequestManager) ListUserRequest(id string) ([]Request, error) {
	return nil, nil
}

func (m *RequestManager) DeleteRequest(id string) error {
	return nil
}
func NewRequestManager() *RequestManager {
	return  &RequestManager{}
}