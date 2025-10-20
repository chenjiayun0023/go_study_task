package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//SQL语句练习

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return db
}

type account struct {
	gorm.Model
	Balance float64 `gorm:"not null;default:0"`
}

type transaction struct {
	gorm.Model
	FromAccountId uint    `gorm:"not null"`
	ToAccountId   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}

func main() {
	db := InitDB()
	//db.AutoMigrate(&account{})
	//db.AutoMigrate(&transaction{})

	// 方案1：闭包捕获变量
	transferAmount := 100.00
	senderID := uint(1111)
	receiverID := uint(2222)

	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		var accountA account
		//指定锁的强度为 UPDATE 锁（行级排他锁）
		res1 := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account{Model: gorm.Model{ID: senderID}}).Scan(&accountA)
		fmt.Println("A的账户信息：", accountA)
		if res1.Error != nil {
			if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("转出账户不存在")
			}
			return res1.Error
		}
		fmt.Printf("转出账户 %d 加锁成功，当前余额: %.2f\n", senderID, accountA.Balance)

		if accountA.Balance < transferAmount {
			return fmt.Errorf("转出账户余额不足，当前余额: %.2f，需要金额: %.2f", accountA.Balance, transferAmount)
		}

		var accountB account
		res2 := tx.First(&account{Model: gorm.Model{ID: receiverID}}).Scan(&accountB)
		fmt.Println("B的账户信息：", accountB)
		if res2.Error != nil {
			if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("转入账户不存在")
			}
			return res2.Error
		}

		// 模拟业务处理时间
		time.Sleep(2 * time.Second)

		updateAccountA := map[string]interface{}{"balance": gorm.Expr("balance - ?", transferAmount), "updated_at": time.Now()}
		res3 := tx.Model(&accountA).Where("(balance - ?) >= 0", transferAmount).Updates(updateAccountA)
		if res3.Error != nil {
			return res3.Error
		}

		//return errors.New("模拟异常事务回滚")

		updateAccountB := map[string]interface{}{"balance": gorm.Expr("balance + ?", transferAmount), "updated_at": time.Now()}
		res4 := tx.Model(&accountB).Updates(updateAccountB)
		if res4.Error != nil {
			return res4.Error
		}

		if err := tx.Create(&transaction{FromAccountId: senderID, ToAccountId: receiverID, Amount: transferAmount}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}
