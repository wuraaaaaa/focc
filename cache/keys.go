package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 日榜
	DailyRankKey = "rank:daily"
)

// VideoViewKey 视频点击数的key view:video:id
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
