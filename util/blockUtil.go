package util

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"strconv"
	"strings"
	"time"
)

type BlockUtil struct {
	Username string
	LimitKey string
	BlockKey string
	KeyTTL   int
}

func NewBlockUtil(username string) *BlockUtil {
	return &BlockUtil{
		Username: username,
		LimitKey: "_LIMIT_KEY_" + strings.ToUpper(username),
		BlockKey: "_BLOCK_KEY_" + strings.ToUpper(username),
		KeyTTL:   global.Config.System.LoginLimitTime,
	}
}

func (b *BlockUtil) GetRemainderCount(c *gin.Context) int {
	limitCount := global.Config.System.LoginLimitCount
	failedCount := b.GetFailedCount(c)
	remainder := limitCount - failedCount
	return remainder
}

func (b *BlockUtil) GetFailedCount(c *gin.Context) int {
	count, err := global.Redis.GetSet(c, b.LimitKey, 0).Result()
	if err != nil {
		global.Logger.Infof("can not get / set key: %v", b.LimitKey)
	}
	IntCount, _ := strconv.Atoi(count)
	return IntCount
}

func (b *BlockUtil) IncrFailedCount(c *gin.Context) {
	count, err := global.Redis.GetSet(c, b.LimitKey, 0).Result()
	if err != nil {
		global.Logger.Infof("can not get / set key: %v", b.LimitKey)
	}
	IntCount, _ := strconv.Atoi(count)
	IntCount++
	if IntCount >= global.Config.System.LoginLimitCount {
		_, err := global.Redis.Set(c, b.BlockKey, "true", time.Minute*time.Duration(b.KeyTTL)).Result()
		if err != nil {
			global.Logger.Infof("can not get / set key: %v", b.BlockKey)
		}
	}
}

func (b *BlockUtil) UnblockUser(c *gin.Context) {
	global.Redis.Del(c, b.LimitKey).Result()
	global.Redis.Del(c, b.BlockKey).Result()
}

func (b *BlockUtil) IsBlocked(c *gin.Context) bool {
	v, _ := global.Redis.Get(c, b.BlockKey).Result()
	if v == "true" {
		return true
	}
	return false
}
