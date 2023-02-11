// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTFriend = "t_friend"

// TFriend mapped from table <t_friend>
type TFriend struct {
	ID       int64 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"` // 主键id
	UserID   int64 `gorm:"column:user_id;type:int unsigned;not null" json:"user_id"`            // 用户id
	FriendID int64 `gorm:"column:friend_id;type:int unsigned;not null" json:"friend_id"`        // 好友id
}

// TableName TFriend's table name
func (*TFriend) TableName() string {
	return TableNameTFriend
}
