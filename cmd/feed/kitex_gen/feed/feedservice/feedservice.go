// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	"context"
	feed "douyinProject/feed/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FeedMethod": kitex.NewMethodInfo(feedMethodHandler, newFeedServiceFeedMethodArgs, newFeedServiceFeedMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "feed",
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

func feedMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feed.FeedServiceFeedMethodArgs)
	realResult := result.(*feed.FeedServiceFeedMethodResult)
	success, err := handler.(feed.FeedService).FeedMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedServiceFeedMethodArgs() interface{} {
	return feed.NewFeedServiceFeedMethodArgs()
}

func newFeedServiceFeedMethodResult() interface{} {
	return feed.NewFeedServiceFeedMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FeedMethod(ctx context.Context, request *feed.FeedReq) (r *feed.FeedResp, err error) {
	var _args feed.FeedServiceFeedMethodArgs
	_args.Request = request
	var _result feed.FeedServiceFeedMethodResult
	if err = p.c.Call(ctx, "FeedMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
