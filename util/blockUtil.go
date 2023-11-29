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
	count, err := global.Redis.Get(c, b.LimitKey).Result()
	if err != nil {
		global.Logger.Infof("can not get set key: %v", b.LimitKey)
	}
	IntCount, _ := strconv.Atoi(count)
	return IntCount
}

func (b *BlockUtil) IncrFailedCount(c *gin.Context) error {
	_, err := global.Redis.Get(c, b.LimitKey).Result()
	if err != nil { // no value for limitKey
		global.Logger.Infof("can not get value for key: %v", b.LimitKey)
		e := global.Redis.Set(c, b.LimitKey, 1, time.Minute*time.Duration(b.KeyTTL)).Err()
		if e != nil {
			global.Logger.Errorf("can not set value for key: %v", b.LimitKey)
			return e
		}
	} else { // already a value for limitKey
		afterIncr, err := global.Redis.Incr(c, b.LimitKey).Result()
		if err != nil {
			global.Logger.Errorf("can not increat value for key: %s\n", err)
			return err
		}
		if int(afterIncr) >= global.Config.System.LoginLimitCount {
			e := global.Redis.Set(c, b.BlockKey, "true", time.Minute*time.Duration(b.KeyTTL)).Err()
			if e != nil {
				global.Logger.Infof("can not set key: %v", b.BlockKey)
				return e
			}
			//return e
		}
	}
	return nil
	//IntCount, _ := strconv.Atoi(count)
	//IntCount++
	//if IntCount >= global.Config.System.LoginLimitCount {
	//	_, err := global.Redis.Set(c, b.BlockKey, "true", time.Minute*time.Duration(b.KeyTTL)).Result()
	//	if err != nil {
	//		global.Logger.Infof("can not set key: %v", b.BlockKey)
	//		return err
	//	}
	//} else {
	//	_, err := global.Redis.Set(c, b.LimitKey, IntCount, time.Minute*time.Duration(b.KeyTTL)).Result()
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil
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

func IsUserBlocked(username string, c *gin.Context) bool {
	blockKey := "_BLOCK_KEY_" + strings.ToUpper(username)
	v, _ := global.Redis.Get(c, blockKey).Result()
	if v == "true" {
		return true
	}
	return false
}
