// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTUser = "t_user"

// TUser mapped from table <t_user>
type TUser struct {
	ID            int64      `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`    // 主键id
	Name          string     `gorm:"column:name;type:varchar(255);not null" json:"name"`                     // 用户名
	FollowCount   int64      `gorm:"column:follow_count;type:int unsigned;not null" json:"follow_count"`     // 关注数
	FollowerCount int64      `gorm:"column:follower_count;type:int unsigned;not null" json:"follower_count"` // 粉丝数
	Password      string     `gorm:"column:password;type:varchar(255);not null" json:"password"`             // 密码
	UpdateDate    *time.Time `gorm:"column:update_date;type:datetime" json:"update_date"`                    // 用户更新日期，格式为mm-dd
	DeleteDate    *time.Time `gorm:"column:delete_date;type:datetime" json:"delete_date"`                    // 用户删除日期，格式为mm-dd
}

// TableName TUser's table name
func (*TUser) TableName() string {
	return TableNameTUser
}
