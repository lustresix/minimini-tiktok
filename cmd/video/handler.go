package main

import (
	"context"
	videoservice "mini-min-tiktok/kitex_gen/videoservice"
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
	// TODO: Your code here...
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoservice.CommentActionReq) (resp *videoservice.CommentActionResp, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) RelationAction(ctx context.Context, req *videoservice.RelationActionReq) (resp *videoservice.RelationActionResp, err error) {
	// TODO: Your code here...
	return
}
