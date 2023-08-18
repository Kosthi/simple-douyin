package service

import (
	"bytes"
	"context"
	"fmt"
	servLog "github.com/prometheus/common/log"
	"github.com/upyun/go-sdk/v3/upyun"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/client"
	"simple-douyin/service/publish/dal"
	"time"
)

func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// 根据userId向db写入视频数据
	// 需要生成PlayUrl和CoverUrl
	servLog.Info(req.UserId)
	servLog.Info(req.Title)
	servLog.Info(len(req.Data))

	resp = publish.NewPublishActionResponse()

	// upload video data to oss
	path := fmt.Sprintf("/videos/%d/%s.mp4", req.UserId, time.Now().Format("2006-01-02 15:04:05"))
	err = client.UpyClient.Put(&upyun.PutObjectConfig{
		Path:   path,
		Reader: bytes.NewReader(req.Data),
	})
	if err != nil {
		servLog.Info("Upload failed.")
		return resp, err
	}
	servLog.Info("Upload success.")

	// write video info into db
	err = dal.WriteVideoInfoIntoDB(ctx, req.UserId, req.Title, path)
	if err != nil {
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, err
	}

	// update work count when redis cache exist
	err = dal.UpdateWorkCount(ctx, req.UserId)
	if err != nil {
		servLog.Error(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, err
	}
	return resp, nil
}
