package main

import (
	"fmt"
)

func Select(b *BankAccount) {
	var op int

	fmt.Printf("입금(1), 출금(2), 종료(other) : ")
	fmt.Scan(&op)

	for op == 1 || op == 2 {
		var err error
		switch op {
		case 1:
			var amount int
			fmt.Printf("입금할 금액을 입력하세요: ")
			fmt.Scan(&amount)
			_, err = b.Deposit(amount)
		case 2:
			var amount int
			fmt.Printf("출금할 금액을 입력하세요: ")
			fmt.Scan(&amount)
			_, err = b.Withdraw(amount)
		default:
			fmt.Println("프로그램을 종료합니다.")
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("입금(1), 출금(2), 종료(other) : ")
		fmt.Scan(&op)
	}

}
