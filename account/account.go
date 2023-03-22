package account

func Withdraw(amount float64, balance float64) float64 {
	if amount < balance {
		return balance - amount
	}

	return balance
}
