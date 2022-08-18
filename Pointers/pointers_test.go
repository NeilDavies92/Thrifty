package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		updateBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		wallet.Withdraw(Bitcoin(20))
		updateBalance(t, wallet, Bitcoin(30))
	})

	t.Run("insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(1)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(2))

		assertError(t, err, "Insufficient Funds")
		updateBalance(t, wallet, startBalance)
	})
}

func updateBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("No Error")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
