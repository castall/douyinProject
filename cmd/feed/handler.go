package main

import (
	"context"
	feed "douyinProject/feed/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// FeedMethod implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) FeedMethod(ctx context.Context, request *feed.FeedReq) (resp *feed.FeedResp, err error) {
	// TODO: Your code here...
	return
}
