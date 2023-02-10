package mw

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"log"
)

var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware 打印rpc的信息，请求和响应
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		log.Printf("real request: %+v\n", req)
		// get remote service information
		log.Printf("remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		log.Printf("real response: %+v\n", resp)
		return nil
	}
}
