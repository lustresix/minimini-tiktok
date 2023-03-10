// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	videoservice "mini-min-tiktok/kitex_gen/videoservice"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*videoservice.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":           kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
		"PublishAction":  kitex.NewMethodInfo(publishActionHandler, newVideoServicePublishActionArgs, newVideoServicePublishActionResult, false),
		"CommentAction":  kitex.NewMethodInfo(commentActionHandler, newVideoServiceCommentActionArgs, newVideoServiceCommentActionResult, false),
		"RelationAction": kitex.NewMethodInfo(relationActionHandler, newVideoServiceRelationActionArgs, newVideoServiceRelationActionResult, false),
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newVideoServiceFavoriteActionArgs, newVideoServiceFavoriteActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "videoservice",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videoservice.VideoServiceFeedArgs)
	realResult := result.(*videoservice.VideoServiceFeedResult)
	success, err := handler.(videoservice.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return videoservice.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return videoservice.NewVideoServiceFeedResult()
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videoservice.VideoServicePublishActionArgs)
	realResult := result.(*videoservice.VideoServicePublishActionResult)
	success, err := handler.(videoservice.VideoService).PublishAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishActionArgs() interface{} {
	return videoservice.NewVideoServicePublishActionArgs()
}

func newVideoServicePublishActionResult() interface{} {
	return videoservice.NewVideoServicePublishActionResult()
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videoservice.VideoServiceCommentActionArgs)
	realResult := result.(*videoservice.VideoServiceCommentActionResult)
	success, err := handler.(videoservice.VideoService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceCommentActionArgs() interface{} {
	return videoservice.NewVideoServiceCommentActionArgs()
}

func newVideoServiceCommentActionResult() interface{} {
	return videoservice.NewVideoServiceCommentActionResult()
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videoservice.VideoServiceRelationActionArgs)
	realResult := result.(*videoservice.VideoServiceRelationActionResult)
	success, err := handler.(videoservice.VideoService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceRelationActionArgs() interface{} {
	return videoservice.NewVideoServiceRelationActionArgs()
}

func newVideoServiceRelationActionResult() interface{} {
	return videoservice.NewVideoServiceRelationActionResult()
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videoservice.VideoServiceFavoriteActionArgs)
	realResult := result.(*videoservice.VideoServiceFavoriteActionResult)
	success, err := handler.(videoservice.VideoService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFavoriteActionArgs() interface{} {
	return videoservice.NewVideoServiceFavoriteActionArgs()
}

func newVideoServiceFavoriteActionResult() interface{} {
	return videoservice.NewVideoServiceFavoriteActionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *videoservice.FeedReq) (r *videoservice.FeedResp, err error) {
	var _args videoservice.VideoServiceFeedArgs
	_args.Req = req
	var _result videoservice.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishAction(ctx context.Context, req *videoservice.PublishActionReq) (r *videoservice.PublishActionResp, err error) {
	var _args videoservice.VideoServicePublishActionArgs
	_args.Req = req
	var _result videoservice.VideoServicePublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, req *videoservice.CommentActionReq) (r *videoservice.CommentActionResp, err error) {
	var _args videoservice.VideoServiceCommentActionArgs
	_args.Req = req
	var _result videoservice.VideoServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationAction(ctx context.Context, req *videoservice.RelationActionReq) (r *videoservice.RelationActionResp, err error) {
	var _args videoservice.VideoServiceRelationActionArgs
	_args.Req = req
	var _result videoservice.VideoServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, req *videoservice.FavoriteActionReq) (r *videoservice.FavoriteActionResp, err error) {
	var _args videoservice.VideoServiceFavoriteActionArgs
	_args.Req = req
	var _result videoservice.VideoServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
