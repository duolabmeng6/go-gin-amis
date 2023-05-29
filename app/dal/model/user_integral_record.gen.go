// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserIntegralRecord = "user_integral_record"

// UserIntegralRecord mapped from table <user_integral_record>
type UserIntegralRecord struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`
	UserID         int64          `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`
	BeforeIntegral int64          `gorm:"column:before_integral;not null;comment:之前货币" json:"before_integral"`
	AfterIntegral  int64          `gorm:"column:after_integral;not null;comment:之后货币" json:"after_integral"`
	Remarks        string         `gorm:"column:remarks;not null;comment:备注" json:"remarks"`
	IntegralType   int32          `gorm:"column:integral_type;not null;comment:类型 0余额 1金币 2api积分" json:"integral_type"`
	ChangeValue    int64          `gorm:"column:change_value;not null;comment:变更的数值" json:"change_value"`
	OrderID        int64          `gorm:"column:order_id;comment:这条流水对应的订单号" json:"order_id"`
	OrderType      int64          `gorm:"column:order_type;comment:0无 1 系统订单 2充值表 3提现表 4api消费表" json:"order_type"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserIntegralRecord's table name
func (*UserIntegralRecord) TableName() string {
	return TableNameUserIntegralRecord
}
