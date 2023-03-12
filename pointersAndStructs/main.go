package main

type Account struct {
	balance int
}

func (a *Account) increment(value int) int { // Receiving account memory address instead of an copy
	a.balance += value

	return a.balance
}

func createAccount() *Account {
	return &Account{ // Return account memory address
		balance: 0,
	}
}

func main() {
	account := Account{
		balance: 100,
	}

	account2 := createAccount()
	account3 := createAccount()

	account.increment(200)
	account2.increment(100)
	account3.increment(200)

	println(account.balance)
	println(account2.balance)
	println(account3.balance)
}
