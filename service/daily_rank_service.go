package service

import (
	"fmt"
	"focc/cache"
	"focc/model"
	"focc/serializer"
	"strings"
)

//DailyRankService  视频每日排行服务
type DailyRankService struct {
	Title string `json:"title,omitempty"`
	Info  string `json:"info,omitempty"`
}

// Get 视频每日排行函数
func (service *DailyRankService) Get() serializer.Response {
	var videos []model.Video

	// 从redis读取点击前十的视频
	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
