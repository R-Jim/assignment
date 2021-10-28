package dao

type Like struct {
	TweetID uint  `gorm:"primaryKey"`
	UserID  uint  `gorm:"primaryKey"`
	Tweet   Tweet `gorm:"foreignKey:TweetID"`
	User    User  `gorm:"foreignKey:UserID" json:"-"`
}
