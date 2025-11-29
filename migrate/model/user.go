package model

import "github.com/gofrs/uuid/v5"

type User struct {
	ID   uuid.UUID `gorm:"column:user_id;type:uuid;primaryKey;default:uuid_generate_v7()"` // UUIDを主キーとして使用
	Name string    `gorm:"column:user_name;type:varchar(255);not null"`
	Age  int       `gorm:"column:age;type:int;not null"`
}

func (User) TableName() string {
	return "user"
}
