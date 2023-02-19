package main

import (
	"context"
	user "douyinProject/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegistMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegistMethod(ctx context.Context, request *user.RegistReq) (resp *user.RegistResp, err error) {
	// TODO: Your code here...
	return
}

// LoginMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginMethod(ctx context.Context, request *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// InfoMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoMethod(ctx context.Context, request *user.InfoReq) (resp *user.InfoResp, err error) {
	// TODO: Your code here...
	return
}

// AuthMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) AuthMethod(ctx context.Context, request *user.AuthReq) (resp *user.AuthResp, err error) {
	// TODO: Your code here...
	return
}
