package wallet

import "testing"

func TestDeposit(t *testing.T) {

	wallet := Wallet{}

	initialBalance := wallet.Balance()

	if initialBalance != 0 {
		t.Errorf("got %d for initial balance, want 0", initialBalance)
	}

	wallet.Deposit(10)
	want := Token(10)
	afterDeposit := wallet.Balance()

	if afterDeposit != want {
		t.Errorf("got %s after deposit, want %s", afterDeposit, want)
	}

}

func TestWithdraw(t *testing.T) {

	t.Run("Allowable withdrawal 1", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(100)
		err := wallet.Withdraw(Token(47))
		
		if err != nil {
			t.Errorf("got an error when none was expected")
		}

		want := Token(53)
		got := wallet.Balance()

		if want != got {
			t.Errorf("got %s after withdrawal, want %s", got, want)
		}
	})

	t.Run("Allowable withdrawal 2", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(100)
		err := wallet.Withdraw(Token(100))
		
		if err != nil {
			t.Errorf("got an error when none was expected")
		}

		want := Token(0)
		got := wallet.Balance()

		if want != got {
			t.Errorf("got %s after withdrawal, want %s", got, want)
		}
	})

	t.Run("Excess withdrawal - should throw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(100)
		err := wallet.Withdraw(Token(101))
		
		// Errors can be compared to the exact error you were expecting and printed as strings
		if err != ErrorInsufficientFunds {
			t.Errorf("An insufficient funds error was expected, got: %s", err)
		}

		// Balance should remain unchanged
		want := Token(100)
		got := wallet.Balance()

		if want != got {
			t.Errorf("got %s after withdrawal, want %s (balance shouldnt have changed)", got, want)
		}
	})
}