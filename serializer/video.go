package serializer

import "focc/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Info      string `json:"info,omitempty"`
	URL       string `json:"url,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.URL,
		Avatar:    item.AvatarURL(),
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {

	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
