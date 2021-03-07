package account

const emptyString = ""

type Account struct {
	id      int
	owner   string
	balance float64
}

func GetAccountById(accountId int, events *[]Event) Account {
	a := &Account{}
	for _, event := range *events {
		if event.getId() == accountId {
			a.On(event)
		}
	}
	return *a
}

func (account *Account) On(event Event) {
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
}

func AddEvent(newEvent Event, events *[]Event) error {
	switch e := newEvent.(type) {
	case AccountCreated:
		if err := validateAccountCreatedEvent(*events, e.getId()); err != nil {
			return err
		}
	case OwnerUpdated:
		if err := validateOwnerUpdatedEvent(GetAccountById(e.getId(), events)); err != nil {
			return err
		}
	case WithdrawalPerformed:
		if err := validateWithdrawalPerformed(GetAccountById(e.getId(), events), e.amount); err != nil {
			return err
		}
	case DepositPerformed:
		if err := validateDepositPerformed(GetAccountById(e.getId(), events), e.amount); err != nil {
			return err
		}
	}
	*events = append(*events, newEvent)
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
