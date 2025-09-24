package model

import (
	"time"
)

// Demo 数据模型
type Demo struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Field1     int        `json:"field1" gorm:"column:field1"`
	Field2     string     `json:"field2" gorm:"type:varchar(255);column:field2"`
	IsDeleted  string     `json:"is_deleted" gorm:"column:is_deleted;default:'N'"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

// TableName 指定表名
func (*Demo) TableName() string {
	return "demo"
}
