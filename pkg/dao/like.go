package dao

type Like struct {
	TweetId uint  `gorm:"primaryKey"`
	UserId  uint  `gorm:"primaryKey"`
	Tweet   Tweet `gorm:"foreignKey:TweetId"`
	User    User  `gorm:"foreignKey:UserId" json:"-"`
}
