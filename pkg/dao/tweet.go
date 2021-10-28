package dao

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Content    string
	AuthorId   uint
	Author     User `gorm:"foreignKey:AuthorId" json:"-"`
	ReplyCount int
	LikeCount  int
}
