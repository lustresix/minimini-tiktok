package main

import (
	"context"
	userservice "mini-min-tiktok/kitex_gen/userservice"
)

// UserserviceImpl implements the last service interface defined in the IDL.
type UserserviceImpl struct{}

// Register implements the UserserviceImpl interface.
func (s *UserserviceImpl) Register(ctx context.Context, req *userservice.UserRegisterReq) (resp *userservice.UserRegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserserviceImpl interface.
func (s *UserserviceImpl) Login(ctx context.Context, req *userservice.UserLoginReq) (resp *userservice.UserLoginResp, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserserviceImpl interface.
func (s *UserserviceImpl) FollowerList(ctx context.Context, req *userservice.RelationFollowerListReq) (resp *userservice.RelationFollowerListResp, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserserviceImpl interface.
func (s *UserserviceImpl) FriendList(ctx context.Context, req *userservice.RelationFriendListReq) (resp *userservice.RelationFriendListResp, err error) {
	// TODO: Your code here...
	return
}
