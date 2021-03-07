package account

import (
	"../repository"
)

const emptyString = ""

type Account struct {
	id      int
	owner   string
	balance float64
}

func GetAccountById(accountId int) Account {
	a := &Account{}
	for _, event := range repository.Events {
		if event.getId() == accountId {
			if err := a.On(event); err != nil {

			}
		}
	}
	return *a
}

func (account *Account) On(event Event) error {
	switch e := event.(type) {
	case AccountCreated:
		account.id = e.id
		account.owner = e.owner
	case OwnerUpdated:
			account.owner = e.owner
	case WithdrawalPerformed:
			account.balance -= e.amount
	case DepositPerformed:
			account.balance += e.amount
	}
	return nil
}


func AddEvent(newEvent Event) error {
	switch e := newEvent.(type) {
	case AccountCreated:
		if err := validateAccountCreatedEvent(repository.Events, e.getId()); err != nil {
			return err
		}
	case OwnerUpdated:
		if err := validateOwnerUpdatedEvent(GetAccountById(e.getId())); err != nil {
			return err
		}
	case WithdrawalPerformed:
		if err := validateWithdrawalPerformed(GetAccountById(e.getId()), e.amount); err != nil {
			return err
		}
	case DepositPerformed:
		if err := validateDepositPerformed(GetAccountById(e.getId()), e.amount); err != nil {
			return err
		}
	}
	repository.Events = append(repository.Events, newEvent)
	return nil
}

func validateAccountCreatedEvent(events []Event, accountId int) error {
	for _, event := range events {
		if accountId == event.getId() {
			return IncorrectIdError{}
		}
	}
	return nil
}

func validateOwnerUpdatedEvent(account Account) error {
	if account.owner == emptyString {
		return AccountNotExistsError{}
	}
	return nil
}

func validateWithdrawalPerformed(account Account, amount float64) error {
	if account.owner == emptyString {
		return AccountNotExistsError{}
	}
	if amount <= 0 || amount > account.balance {
		return IncorrectAmountError{}
	}
	return nil
}

func validateDepositPerformed(account Account, amount float64) error {
	if account.owner == emptyString {
		return AccountNotExistsError{}
	}
	if amount <= 0 {
		return IncorrectAmountError{}
	}
	return nil
}