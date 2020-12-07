package models

type AccountManager struct {}

func NewAccountManager() *AccountManager{
	return &AccountManager{}
}

type User struct {
	 Name string
	 Password string
	 Email string
}

func (m *AccountManager) Login() error {
	return nil
}

func (m *AccountManager) Logout() error {
	return nil
}

func (m *AccountManager) ListFriend(id string) ([]User, error) {
	return nil, nil
}

func (m *AccountManager) AddFriend(id, friendID string) error {
	return nil
}

func (m *AccountManager) DeleteFriend(id, friendID string) error {
	return nil
}

func (m *AccountManager) SearchUsers(name string) ([]User, error) {
	return nil, nil
}
