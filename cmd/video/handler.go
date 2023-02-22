package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/nanakura/go-ramda"
	"mini-min-tiktok/cmd/video/mw/cos"
	"mini-min-tiktok/cmd/video/utils"
	videoservice "mini-min-tiktok/kitex_gen/videoservice"
	"mini-min-tiktok/pkg/cache"
	"mini-min-tiktok/pkg/configs/config"
	"mini-min-tiktok/pkg/consts"
	"mini-min-tiktok/pkg/dal/model"
	"mini-min-tiktok/pkg/dal/query"
	jwt "mini-min-tiktok/pkg/utils"
	"strconv"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

type queryVideoListRes struct {
	ID            int64 // 视频id
	AuthorID      int64 `sql:"author_id"`
	Name          string
	FollowCount   int64
	FollowerCount int64
	Password      string
	PlayURL       string    // 视频链接
	CoverURL      string    // 视频封面链接
	FavoriteCount int64     // 点赞数
	CommentCount  int64     // 评论数
	IsFavorite    bool      // 是否已点赞(0为未点赞, 1为已点赞)
	Title         string    // 视频标题
	CreateDate    time.Time // 视频上传时间
}

func CastQueryVideoListtoVideoServiceVideo(from []queryVideoListRes) []*videoservice.Video {
	return ramda.Map(func(model queryVideoListRes) *videoservice.Video {
		return &videoservice.Video{
			Id: model.ID,
			Author: &videoservice.User{
				Id:            model.AuthorID,
				Name:          model.Name,
				FollowCount:   model.FollowCount,
				FollowerCount: model.FollowerCount,
			},
			PlayUrl:       model.PlayURL,
			CoverUrl:      model.CoverURL,
			FavoriteCount: model.FavoriteCount,
			CommentCount:  model.CommentCount,
			IsFavorite:    model.IsFavorite,
			Title:         model.Title,
		}
	})(from)
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.FeedReq) (resp *videoservice.FeedResp, err error) {
	latestTime, _ := strconv.ParseInt(*req.LatestTime, 10, 64)
	// 值为0（默认值）则说明不限制最新时间
	tv := query.Q.TVideo.As("v")
	tu := query.Q.TUser.As("u")
	var resList []queryVideoListRes
	if latestTime == 0 {
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.ALL,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.ID,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Order(tv.ID.Desc()).
			Limit(10).Scan(&resList)

		if err != nil {
			return
		}
	} else {
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.Name,
				tu.Password,
				tu.FollowCount,
				tu.FollowerCount,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.ID,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Order(tv.ID.Desc()).
			Limit(10).
			Scan(&resList)
		if err != nil {
			return
		}
	}
	if resList == nil {
		resList = []queryVideoListRes{}
	}
	resp = &videoservice.FeedResp{
		StatusCode: 0,
		VideoList:  CastQueryVideoListtoVideoServiceVideo(resList),
	}
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

	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.FavoriteActionReq) (resp *videoservice.FavoriteActionResp, err error) {
	claims, flag := jwt.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("token is expired")
	}
	redis := cache.RedisCache.RedisClient
	// 查看是否点赞
	result, err := redis.SIsMember(context.Background(), "post_set"+":"+consts.FavoriteActionPrefix+req.VideoId,
		strconv.FormatInt(claims.UserId, 10)).Result()
	if err != nil {
		err = fmt.Errorf("redis访问失败")
	}
	// 如果已经点赞
	id, err := strconv.ParseInt(req.VideoId, 10, 64)
	if result {
		// 取消关联
		_, err := redis.SRem(context.Background(), "post_set"+":"+consts.FavoriteActionPrefix+req.VideoId,
			strconv.FormatInt(claims.UserId, 10)).Result()
		if err != nil {
			err = fmt.Errorf("取消点赞失败")
			return
		}
		// 将视频总点赞减一
		utils.LikeNumDel(id)
		resp = &videoservice.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}
	// 在数据库中查询点赞信息
	q := query.Q
	favorite := q.TFavorite
	first, _ := q.WithContext(context.Background()).TFavorite.Where(favorite.UserID.Eq(claims.UserId), favorite.VideoID.Eq(id)).First()

	// 查询为空
	if first == nil {
		// 将点赞存入redis
		redis.SAdd(context.Background(), "post_set"+":"+consts.FavoriteActionPrefix+req.VideoId, strconv.FormatInt(claims.UserId, 10), 0)
		resp = &videoservice.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "已成功点赞",
		}
		// 将视频总点赞数加一
		_ = utils.LikeNumAdd(id)
		return
	}

	// 查询数据库，数据库为已点赞，取消点赞
	if first.Status {
		q.WithContext(context.Background()).TFavorite.Update(favorite.Status, false)
		if err != nil {
			err = fmt.Errorf("更新数据库")
		}
		resp = &videoservice.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}
	resp = &videoservice.FavoriteActionResp{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	}
	return
}
