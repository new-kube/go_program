package study

import (
	"fmt"

	"gorm.io/gorm"
)

func Start(db *gorm.DB) error {

	err := db.Transaction(func(tx *gorm.DB) error {

		user := User{
			ID:    1,
			Name:  "hello",
			Phone: "12345678901",
		}

		xdb := tx.Where("id = ?", user.ID).Updates(&user)
		if err := xdb.Error; err != nil {
			fmt.Printf("update failed, err=%v\n", err)
			return err
		}

		fmt.Printf("updated first: %t\n", xdb.RowsAffected > 0)

		user.ID = 2
		user.Name = "newbee"

		xdb = tx.Where("id = ?", user.ID).Updates(&user)
		if err := xdb.Error; err != nil {
			fmt.Printf("update failed, err=%v\n", err)
			return err
		}

		fmt.Printf("updated first: %t\n", xdb.RowsAffected > 0)

		return nil
	})

	return err
}
