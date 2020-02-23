package service

import (
	"focc/model"
	"focc/serializer"
)

//CreatVideoService 视频投稿服务
type CreatVideoService struct {
	Title string `json:"title,omitempty"`
	Info  string `json:"info,omitempty"`
}

// Creat 视频创建函数
func (service *CreatVideoService) Creat() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
