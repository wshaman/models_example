package model


var errMockMethodNotImplemented = errors.New("mock method not implemented yet")

type modelMock struct{}

// NewMockModel returns mocked model
func NewMockModel() Model {
	return &mockModel{}
}
