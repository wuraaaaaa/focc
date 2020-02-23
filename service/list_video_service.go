package service

import (
	"focc/model"
	"focc/serializer"
)

//ListVideoService 视频列表服务
type ListVideoService struct {
	Title string `json:"title,omitempty"`
	Info  string `json:"info,omitempty"`
}

// List 视频列表函数
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频列表查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
