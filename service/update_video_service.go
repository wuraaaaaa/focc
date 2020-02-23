package service

import (
	"focc/model"
	"focc/serializer"
)

//UpdateVideoService 视频详情服务
type UpdateVideoService struct {
	Title string `json:"title,omitempty"`
	Info  string `json:"info,omitempty"`
}

// Update 更新视频信息
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
