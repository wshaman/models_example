package model

func (m *modelMock) ListUsers() ([]User, error) {
	return nil, errMockMethodNotImplemented
}
