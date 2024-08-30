package user

import (
    "github.com/hd2yao/ecom/types"
    "testing"
)

func TestUserServiceHandlers(t *testing.T) {
    userStore := &mockUserStore{}
    handler := NewHandler(userStore)
}

// 模拟(mock)实现 UserStore 接口

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
    return &types.User{}, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
    return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
    return nil
}
