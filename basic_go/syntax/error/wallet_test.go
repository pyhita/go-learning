package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		if err == nil {
			// t.Fatal。如果它被调用，它将停止测试。这是因为我们不希望对返回的错误进行更多断言。
			//如果没有这个，测试将继续进行下一步，并且因为一个空指针而引起 panic。
			t.Fatal("err should not be nil")
		}

		if err != want {
			t.Errorf("got %s, want %s", err.Error(), want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(100)
		want := Bitcoin(100)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(10)
		want := Bitcoin(90)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(20)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertError(t, err, InsufficientFundsError)
	})
}
