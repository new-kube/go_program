package study

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DmAdapterCompany struct {
	ID       int    `gorm:"column:id" json:"id"`             //自增id
	Company  string `gorm:"column:company" json:"company"`   //快递公司名称
	Code     string `gorm:"column:code" json:"code"`         //快递公司编码
	Supplier int    `gorm:"column:supplier" json:"supplier"` //履约方
	Created  int    `gorm:"column:created" json:"created"`   //创建时间
	Updated  int    `gorm:"column:updated" json:"updated"`   //修改时间
}

func SaveCompany(ctx context.Context, db *gorm.DB) (int, error) {
	now := int(time.Now().Unix())

	data := DmAdapterCompany{
		Company:  "中国邮政",
		Code:     "YZPY",
		Supplier: 2,
		Created:  now,
		Updated:  now,
	}

	xdb := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "company"}, {Name: "supplier"}},
		Where: clause.Where{Exprs: []clause.Expression{
			clause.Neq{Column: clause.Column{Name: "code"}, Value: "YZPY"},
		}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "updated"}),
	}).Create(&data)

	if err := xdb.Error; err != nil {
		fmt.Printf("update record: %v failed, err=%v\n", data, err)
		return 0, err
	}

	return int(xdb.RowsAffected), nil
}
