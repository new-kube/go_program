package study

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() (*gorm.DB, error) {

	user := "root"
	pass := "123456"
	host := "localhost"
	port := 3306
	database := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, database)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			//Logger:                 log.Default(),
			NamingStrategy: schema.NamingStrategy{ // 禁用表名复数。
				SingularTable: true,
			},
		},
	)
	if err != nil {
		fmt.Printf("connect db=%s failed, err=%v\n", database, err)
		return nil, err
	}

	return db, nil
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) TableCreate(ctx context.Context) error {
	// 创建表，如果存在表，则会根据user定义修改表。

	if r.db.Migrator().HasTable(TableUser) {
		return nil
	}
	return r.db.Table(TableUser).AutoMigrate(&User{})
}

/*
CREATE TABLE `user` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `name` longtext,
   `phone` longtext,
   `created` datetime(3) DEFAULT NULL,
   `updated,autoUpdateTime` datetime(3) DEFAULT NULL,
   `created_at` datetime(3) DEFAULT NULL,
   `updated_at` bigint(20) DEFAULT NULL,
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8\



 CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(500) NOT NULL DEFAULT '' COMMENT '姓名',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '电话',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='user';
*/

func (r *UserRepo) Create(ctx context.Context, user *User) error {
	xdb := r.db.WithContext(ctx)
	return xdb.Create(user).Error
}

func (r *UserRepo) Update(ctx context.Context, user *User) error {

	xdb := r.db.WithContext(ctx)

	return xdb.Table(TableUser).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepo) ForceUpdate(ctx context.Context, user *User) error {

	xdb := r.db.WithContext(ctx)

	// 这种更新能强制更新。
	updated := map[string]any{
		"name":    user.Name,
		"phone":   user.Phone,
		"updated": user.Updated, // 虽然这个有表定义的约束：CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP，但是当指定了特定值，就能更新。
	}

	return xdb.Table(TableUser).Where("id = ?", user.ID).Updates(&updated).Error
}

func (r *UserRepo) UpdateSome(ctx context.Context, user *User) error {

	xdb := r.db.WithContext(ctx)

	// 这种更新能强制更新。
	updated := map[string]any{
		"name":  user.Name,
		"phone": user.Phone,
	}

	return xdb.Table(TableUser).Where("id = ?", user.ID).Updates(&updated).Error
}
