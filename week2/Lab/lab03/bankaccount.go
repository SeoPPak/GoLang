package main

import (
	"fmt"
)

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) (int, error) {
	if amount < 0 {
		err := fmt.Errorf("오류: 음수 금액은 입금할 수 없습니다.: %d원", amount)
		return 0, err
	} else {
		b.Balance += amount
		fmt.Printf("입금 성공! 현재 잔액: %d원\n", b.Balance)
		return b.Balance, nil
	}
}

func (b *BankAccount) Withdraw(amount int) (int, error) {
	if b.Balance < amount {
		err := fmt.Errorf("오류: 잔액이 부족합니다. 현재 잔액: %d원, 요청한 금액: %d원", b.Balance, amount)
		return -1, err
	} else {
		b.Balance -= amount
		fmt.Printf("출금 성공! 현재 잔액: %d원\n", b.Balance)
		return b.Balance, nil
	}
}
