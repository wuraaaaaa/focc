package service

import (
	"focc/model"
	"focc/serializer"
)

//ShowVideoService 视频详情服务
type ShowVideoService struct {
}

// Show 展示视频信息
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频查询错误",
			Error: err.Error(),
		}
	}
	//观看数
	video.AddView()
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
