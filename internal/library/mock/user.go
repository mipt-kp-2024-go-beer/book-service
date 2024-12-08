package mock

type MockUserServiceClient struct{}

func NewMockUserServiceClient() *MockUserServiceClient {
	return &MockUserServiceClient{}
}

func (client *MockUserServiceClient) CheckPermissions(token string, mask uint) (bool, error) {
	return true, nil
}
