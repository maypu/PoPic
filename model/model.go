package model

import (
	"gorm.io/gorm"
)

type Common struct {
	Remarks   string
	Status    string `gorm:"default:1"`
	IpAddress string
}

// User 用户
type User struct {
	gorm.Model
	Nickname string
	Mail     string
	Password string
	Common
}

// Platform 上传平台
type Platform struct {
	gorm.Model
	Name string
	Common
}

type Upload struct {
	gorm.Model
	User      string
	Platform  int
	ImgUrl    string
	DeleteUrl string
	Common
}
