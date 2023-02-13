package main

import (
	"context"
	"fmt"
	"mini-min-tiktok/cmd/user/utils"
	userservice "mini-min-tiktok/kitex_gen/userservice"
	"mini-min-tiktok/pkg/dal/model"
	"mini-min-tiktok/pkg/dal/query"
	jwt "mini-min-tiktok/pkg/utils"
	"strconv"
)

// UserserviceImpl implements the last service interface defined in the IDL.
type UserserviceImpl struct{}

// Register implements the UserserviceImpl interface.
func (s *UserserviceImpl) Register(ctx context.Context, req *userservice.UserRegisterReq) (resp *userservice.UserRegisterResp, err error) {
	// 检查用户是否存在
	q := query.Q
	checkRes, _ := utils.CheckUser(q, req.Username, req.Password)
	if checkRes != nil {
		err = fmt.Errorf("注册失败：用户已存在 %w", err)
		return
	}
	// 不存在，密码加密存入数据库
	pwd := utils.ScryptPwd(req.Password)
	newUser := &model.TUser{Name: req.Username, Password: pwd}
	err = q.WithContext(context.Background()).TUser.Create(newUser)
	if err != nil {
		err = fmt.Errorf("注册失败: %w", err)
		return
	}
	// 生成 token
	token, err := jwt.CreateToken(newUser.ID)
	if err != nil {
		err = fmt.Errorf("token生成失败: %w", err)
		return
	}
	// 返回数据
	resp = &userservice.UserRegisterResp{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserId:     newUser.ID,
		Token:      token,
	}
	return
}

// Login implements the UserserviceImpl interface.
func (s *UserserviceImpl) Login(ctx context.Context, req *userservice.UserLoginReq) (resp *userservice.UserLoginResp, err error) {
	q := query.Q
	checkRes, err := utils.CheckUser(q, req.Username, req.Password)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return
	}
	// 生成 token
	token, err := jwt.CreateToken(checkRes.ID)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return
	}
	// 成功返回
	resp = &userservice.UserLoginResp{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		Token:      token,
		UserId:     checkRes.ID,
	}
	return
}

// FollowerList implements the UserserviceImpl interface.
func (s *UserserviceImpl) FollowerList(ctx context.Context, req *userservice.RelationFollowerListReq) (resp *userservice.RelationFollowerListResp, err error) {
	// 关注列表
	queryFollow := query.Q.TFollow
	// 进行查询用户
	queryUser := query.Q.TUser
	// 格式转换
	id, _ := strconv.ParseInt(req.UserId, 10, 64)
	// 只进行查询关注用户的 id
	follows, err := queryFollow.WithContext(ctx).Select(queryFollow.FollowerID).Where(queryFollow.UserID.Eq(id)).Find()
	if err != nil {
		return
	}
	followerIds := make([]int64, len(follows))
	// 将关注用户的 id 进行提取出来
	for i, follow := range follows {
		followerIds[i] = follow.FollowerID
	}
	// 使用 select 进行规范查询的数据，使得不查询密码
	// 根据关注用户 id 查询到所有的 关注用户信息
	users, err := queryUser.WithContext(ctx).Select(queryUser.ID, queryUser.Name,
		queryUser.FollowCount, queryUser.FollowerCount).Where(queryUser.ID.In(followerIds...)).Find()
	if err != nil {
		return
	}
	// 将所有的关注用户信息进行添加到 返回值中
	var user userservice.User
	var list []*userservice.User
	for _, tUser := range users {
		user.Id = tUser.ID
		user.Name = tUser.Name
		user.FollowerCount = tUser.FollowerCount
		user.FollowCount = tUser.FollowCount
		user.IsFollow = true
		var us userservice.User
		us = user
		list = append(list, &us)
	}
	resp = &userservice.RelationFollowerListResp{
		StatusCode: 0,
		StatusMsg:  "the request succeeded",
		UserList:   list,
	}
	return
}

// FriendList implements the UserserviceImpl interface.
func (s *UserserviceImpl) FriendList(ctx context.Context, req *userservice.RelationFriendListReq) (resp *userservice.RelationFriendListResp, err error) {
	resp = &userservice.RelationFriendListResp{}
	qFriend := query.Q.TFriend
	qUser := query.Q.TUser
	qFollow := query.Q.TFollow
	// 格式转换
	id, _ := strconv.ParseInt(req.UserId, 10, 64)
	// 查询 查看用户的好友
	friendUsers, err := qFriend.WithContext(ctx).Select(qFriend.FriendID).Where(qFriend.UserID.Eq(id)).Find()
	if err != nil {
		if err.Error() == "record not found" {
			resp.StatusCode = 0
			resp.StatusMsg = "用户没有好友"
			resp.UserList = nil
			return resp, nil
		}
		return nil, err
	}
	userIds := make([]int64, len(friendUsers))
	// 抽离出粉丝的用户 id
	for i, user := range friendUsers {
		userIds[i] = user.FriendID
	}
	// 对关注的用户进行查询
	queryUsers, _ := qUser.WithContext(ctx).Where(qUser.ID.In(userIds...)).Find()
	users := make([]userservice.User, len(queryUsers))
	claims, _ := jwt.CheckToken(req.Token)
	// 如果查看用户与当前登录用户是好友，不需要返回自身的数据
	// 如果这个数大于 -1 ，说明登陆用户与查看用户是好友，将此数据进行剔除
	whetherExistCurrentUser := -1
	for i, queryUser := range queryUsers {

		if queryUser.ID == claims.UserId {
			whetherExistCurrentUser = i
			continue
		}
		users[i].Id = queryUser.ID
		users[i].Name = queryUser.Name
		users[i].FollowerCount = queryUser.FollowerCount
		users[i].FollowCount = queryUser.FollowCount
	}
	// 进行剔除登录用户的数据
	if whetherExistCurrentUser >= 0 {
		users = append(users[:whetherExistCurrentUser], users[whetherExistCurrentUser+1:]...)
	}
	var lists []*userservice.User
	// 如果查看的用户是自己，就不需要查询是否已经关注
	if id == claims.UserId {
		for i := 0; i < len(users); i++ {
			users[i].IsFollow = true
			lists = append(lists, &users[i])
		}
	} else {
		for i := 0; i < len(users); i++ {
			whetherToCare, err := qFollow.WithContext(ctx).
				Where(qFollow.UserID.Eq(claims.UserId), qFollow.FollowerID.Eq(users[i].Id)).First()
			if err == nil && whetherToCare != nil {
				users[i].IsFollow = true
			} else {
				users[i].IsFollow = false
			}
			lists = append(lists, &users[i])
		}
	}
	resp = &userservice.RelationFriendListResp{
		UserList:   lists,
		StatusCode: 0,
		StatusMsg:  "查询成功",
	}
	return resp, nil
}
