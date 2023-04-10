package main

import (
	"context"
	"fmt"
	"time"

	"orm/study"

	"gorm.io/gorm"
)

func testUpdate(db *gorm.DB) {
	repo := study.NewUserRepo(db)

	ctx := context.Background()

	if err := repo.TableCreate(ctx); err != nil {
		fmt.Printf("table create failed, err=%v\n", err)
		//return
	}

	user := study.User{
		Name:  "yangxingya",
		Phone: "13312345678",
	}

	if err := repo.Create(ctx, &user); err != nil {
		fmt.Printf("create failed, err=%v\n", err)
		return
	}

	now := time.Now().Unix()
	yes := now - int64(time.Hour/time.Second*24)
	tm := time.Unix(yes, 0)

	user.Updated = tm
	if err := repo.Update(ctx, &user); err != nil {
		fmt.Printf("update failed, err=%v\n", err)
		return
	}

	user2 := study.User{
		ID:    2,
		Name:  "yangxingya",
		Phone: "12345678901",
	}
	if err := repo.Create(ctx, &user2); err != nil {
		fmt.Printf("update some failed, err=%v\n", err)
		return
	}
}

func main() {
	// 连接数据库
	db, err := study.Connect()
	if err != nil {
		fmt.Printf("connect failed, err=%v\n", err)
		return
	}
	fmt.Printf("connect succeed...\n")

	testUpdate(db)
	// study.Start(db)
}
