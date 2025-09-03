// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
// 和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， 
// to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
// 需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID int `gorm:"primarykey"`
	Balance float64
}

type Transaction struct {
	ID int `gorm:"primarykey"`
	FromAccountID int
	ToAccountID int
	Amount float64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	// 查询账户 A 的余额 账户A的ID为1
	var account Account
	if err := tx.First(&account, 1).Error; err!= nil {
		tx.Rollback()
		fmt.Println("查询账户 A 失败：", err)
		return
	}

		// 余额不足，回滚事务
	if account.Balance < 100 {
		tx.Rollback()			
			fmt.Println("账户 A 余额不足，转账失败")
			return		
	}

	// 扣除 100 元
	account.Balance -= 100
	if err := tx.Save(&account).Error; err!= nil {
		tx.Rollback()
		fmt.Println("扣除 100 元失败：", err)
		return
	}

	// 向账户 B 增加 100 元 账户B的ID为2
	var toAccount Account
	if err := tx.First(&toAccount, 2).Error; err!= nil {
		tx.Rollback()
		fmt.Println("查询账户 B 失败：", err)
		return
	}
	toAccount.Balance += 100
	if err := tx.Save(&toAccount).Error; err!= nil {
		tx.Rollback()
		fmt.Println("向账户 B 增加 100 元失败：", err)
		return
	}

	// 记录转账信息
	transaction := Transaction{
		FromAccountID: 1,
		ToAccountID: 2,
		Amount: 100,
	}
	if err := tx.Create(&transaction).Error; err!= nil {
		tx.Rollback()
		fmt.Println("记录转账信息失败：", err)
		return
	}

	// 提交事务
	tx.Commit()
	fmt.Println("转账成功")		
}