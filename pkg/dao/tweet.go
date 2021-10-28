package dao

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Content    string
	AuthorID   uint
	Author     User `gorm:"foreignKey:AuthorID" json:"-"`
	ReplyCount int
	LikeCount  int
}
