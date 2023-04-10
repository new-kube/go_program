package study

import "time"

type User struct {
	ID      int       `gorm:"column:id" json:"id"`
	Name    string    `gorm:"column:name;unique" json:"name"`
	Phone   string    `gorm:"column:phone" json:"phone"`
	Created time.Time `gorm:"column:created" json:"created"`
	Updated time.Time `gorm:"column:updated" json:"updated"`
}

const (
	TableUser = "user"
)

func (u User) TableName() string {
	return TableUser
}
