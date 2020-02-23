package service

import (
	"focc/model"
	"focc/serializer"
)

//DeleteVideoService 视频投稿服务
type DeleteVideoService struct {
}

// Delete 视频删除函数
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "视频删除成功",
	}
}
