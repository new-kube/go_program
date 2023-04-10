package study

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// test group by, order by, sum(*) etc.

// model

// select name, sum(num) as num from test_group where status = 1 group by name order by num desc;

type TestGroup struct {
	ID     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Phone  string `gorm:"column:phone" json:"phone"`
	Status int    `gorm:"column:status" json:"status"`
	Num    int    `gorm:"column:num" json:"num"`
}

type Result struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

/*
create table `test_group` (
	id int(11) NOT NULL auto_increment,
	name varchar(255) NOT NULL default '',
	phone varchar(255) NOT NULL default '',
	status int(10) unsigned NOT NULL default '0',
	num int(10) unsigned NOT NULL default '0',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 COMMENT='test_group_by';
*/

func GroupBy(ctx context.Context, db *gorm.DB) ([]*Result, error) {
	results := make([]*Result, 0)

	xdb := db.WithContext(ctx)
	xdb = xdb.Table("test_group").Select("name, SUM(num) as total").Where("status = 1").Group("name").Order("total desc").Find(&results)

	if err := xdb.Error; err != nil {
		return results, err
	}

	return results, nil
}

func AddTestRecords(ctx context.Context, db *gorm.DB) error {

	data := make([]*TestGroup, 0)
	for i := 0; i < 100; i++ {
		t := &TestGroup{
			Name:   fmt.Sprintf("name_%d", i%10),
			Phone:  fmt.Sprintf("138111122%02d", i),
			Status: 1,
			Num:    i,
		}
		data = append(data, t)
	}

	xdb := db.WithContext(ctx).Create(&data)
	if err := xdb.Error; err != nil {
		return err
	}

	return nil

}
