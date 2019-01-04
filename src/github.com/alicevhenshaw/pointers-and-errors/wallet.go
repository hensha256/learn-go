package wallet

import "fmt"
import "errors"

type Token int

type Wallet struct {
	balance Token
}

// Errors are just values and so can be stored in variables
var ErrorInsufficientFunds = errors.New("negative balance not permitted")

func (w *Wallet) Balance() Token {
	return w.balance
}

func (w *Wallet) Deposit(toDeposit Token) {
	w.balance += toDeposit
}

// Allows an error to be thrown if the withdrawal amount will take the balance negative
func (w *Wallet) Withdraw(toWithdraw Token) error {
	if w.balance >= toWithdraw {
		w.balance -= toWithdraw
		return nil
	}
	return ErrorInsufficientFunds
}

// Implementing this function on a type makes it compatible with the Stringer interface
// In a print statement %s can be used on the Token type and it will print the following:
func (t Token) String() string {
	return fmt.Sprintf("%d tokens", t)
}