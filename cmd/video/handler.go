package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"mini-min-tiktok/cmd/video/mw/cos"
	videoservice "mini-min-tiktok/kitex_gen/videoservice"
	"mini-min-tiktok/pkg/configs/config"
	"mini-min-tiktok/pkg/dal/model"
	"mini-min-tiktok/pkg/dal/query"
	jwt "mini-min-tiktok/pkg/utils"
	"strconv"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.FeedReq) (resp *videoservice.FeedResp, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.PublishActionReq) (resp *videoservice.PublishActionResp, err error) {
	resp = &videoservice.PublishActionResp{}
	l := len(req.Data)
	klog.Infof("视频长度：%d", l)
	// mb不知道为什么thrift的byte生成出来的int8啊啊啊
	bytes := make([]byte, l)
	for i, _ := range req.Data {
		bytes[i] = byte(req.Data[i])
	}
	// 生成唯一通识码
	uuidv4, _ := uuid.NewUUID()
	uuidname := uuidv4.String()
	path := fmt.Sprintf("%s.mp4", uuidname)

	tv := query.Q.TVideo
	cliams, _ := jwt.CheckToken(req.Token)
	userId := cliams.UserId
	// 将视频保存到cos里
	videoPath, photoPath, err := cos.SaveUploadedFile(ctx, bytes, path)
	if err != nil {
		return
	}
	// 将元数据存入数据库
	url := config.GlobalConfigs.CosConfig.Url
	err = tv.WithContext(context.Background()).
		Create(&model.TVideo{
			AuthorID:      userId,
			PlayURL:       fmt.Sprintf("%s%s", url, videoPath),
			CoverURL:      fmt.Sprintf("%s%s", url, photoPath),
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         req.Title,
			//CreateDate:    time.Now(),
		})
	if err != nil {
		klog.Error("Error uploading file:", err)
		err = fmt.Errorf("视频保存失败：%w", err)
		return
	}
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoservice.CommentActionReq) (resp *videoservice.CommentActionResp, err error) {
	// 评论操作
	queryUser := query.Q.TUser
	queryVideo := query.Q.TVideo
	queryComment := query.Q.TComment
	timeLayoutStr := "2006-01-02 15:04:05"
	// 解析 token 拿取用户id
	claims, flag := jwt.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("token is expired")
	}
	id, _ := strconv.ParseInt(req.VideoId, 10, 64)
	// 判断视频是否存在
	_, err = queryVideo.WithContext(ctx).Where(queryVideo.ID.Eq(id)).First()
	if err != nil {
		return nil, errors.New("video does not exist")
	}
	a, _ := strconv.ParseInt(req.ActionType, 10, 64)
	// 发布评论
	if a == 1 {
		comment := &model.TComment{
			UserID:     claims.UserId,
			Content:    *req.CommentText,
			VideoID:    id,
			CreateDate: time.Now(),
		}

		err := queryComment.WithContext(ctx).Create(comment)
		user, _ := queryUser.WithContext(ctx).Select(queryUser.ID, queryUser.Name).
			Where(queryUser.ID.Eq(claims.UserId)).First()
		if err != nil {
			return nil, errors.New("add failure")
		}
		resp = &videoservice.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			Comment: &videoservice.Comment{
				Id: comment.ID,
				User: &videoservice.User{
					Id:            user.ID,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
				},
				Content:    comment.Content,
				CreateDate: comment.CreateDate.Format(timeLayoutStr),
			},
		}
		// 删除评论
	} else if a == 2 {
		com, _ := strconv.ParseInt(*req.CommentId, 10, 64)
		// 用户是否有此条评论
		_, err := queryComment.WithContext(ctx).Where(queryComment.ID.Eq(com)).
			Where(queryComment.UserID.Eq(claims.UserId)).Delete()
		if err != nil {
			return nil, errors.New("comment does not exist")
		}
		resp = &videoservice.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "删除成功",
		}
	} else {
		return nil, errors.New("operation error")
	}
	return
}

// RelationAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) RelationAction(ctx context.Context, req *videoservice.RelationActionReq) (resp *videoservice.RelationActionResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.FavoriteActionReq) (resp *videoservice.FavoriteActionResp, err error) {
	// TODO: Your code here...
	return
}
