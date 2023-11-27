package service

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"go-blog/global"
	"go-blog/models"
	"go-blog/util"
	"gorm.io/gorm"
)

type UserService struct{}

func (us UserService) Register(u models.User) (newUser models.User, err error) {
	var user models.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return newUser, errors.New("用户名已注册")
	}
	u.Password = util.BcryptHash(u.Password)
	u.Uuid = uuid.Must(uuid.NewV4())
	err = global.DB.Create(&u).Error
	return u, err
}

func (us UserService) Login(u *models.User) (newUser *models.User, err error) {
	global.Logger.Info("service login")
	if nil == global.DB {
		return nil, fmt.Errorf("DB has not been connected: %v", global.Config.Mysql.Host)
	}
	var user models.User
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		global.Logger.Info(err)
		if ok := util.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("wrong password, please try again")
		}
	}
	return &user, err
}
