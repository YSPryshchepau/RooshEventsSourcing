package account

import "testing"

const testAccountId = 1
const initialOwnerName = "Test"
const updatedOwnerName = "Test1"
const depositAmount = 100
const incorrectAmount = -10
const withdrawalAmount = 75
const tooBigWithdrawal = 110
const balanceAmount = 25

func TestGetAccountById(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	events = append(events, OwnerUpdated{id: testAccountId, owner: updatedOwnerName})
	events = append(events, DepositPerformed{id: testAccountId, amount: depositAmount})
	events = append(events, WithdrawalPerformed{id: testAccountId, amount: withdrawalAmount})
	expected := Account{
		id:      testAccountId,
		owner:   updatedOwnerName,
		balance: balanceAmount,
	}
	actual := GetAccountById(testAccountId, &events)
	if expected != actual {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestOnAccountCreated(t *testing.T) {
	expected := Account{
		id:    testAccountId,
		owner: initialOwnerName,
	}
	actual := Account{}
	actual.On(AccountCreated{id: testAccountId, owner: initialOwnerName})
	if expected != actual {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestOnOwnerUpdated(t *testing.T) {
	expected := Account{
		id:    testAccountId,
		owner: updatedOwnerName,
	}
	actual := Account{}
	actual.On(AccountCreated{id: testAccountId, owner: initialOwnerName})
	actual.On(OwnerUpdated{id: testAccountId, owner: updatedOwnerName})
	if expected != actual {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestOnDepositPerformed(t *testing.T) {
	expected := Account{
		id:      testAccountId,
		owner:   initialOwnerName,
		balance: depositAmount,
	}
	actual := Account{}
	actual.On(AccountCreated{id: testAccountId, owner: initialOwnerName})
	actual.On(DepositPerformed{id: testAccountId, amount: depositAmount})
	if expected != actual {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestOnWithdrawalPerformed(t *testing.T) {
	expected := Account{
		id:      testAccountId,
		owner:   initialOwnerName,
		balance: balanceAmount,
	}
	actual := Account{}
	actual.On(AccountCreated{id: testAccountId, owner: initialOwnerName})
	actual.On(DepositPerformed{id: testAccountId, amount: depositAmount})
	actual.On(WithdrawalPerformed{id: testAccountId, amount: withdrawalAmount})
	if expected != actual {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestAddEventAccountCreatedSuccess(t *testing.T) {
	var events []Event
	err := AddEvent(AccountCreated{id: testAccountId, owner: initialOwnerName}, &events)
	if err != nil {
		t.Error(err)
	}
}

func TestAddEventAccountCreatedFailure(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	err := AddEvent(AccountCreated{id: testAccountId, owner: initialOwnerName}, &events)
	if err == nil {
		t.Error("Expected ", IncorrectIdError{})
	}
}

func TestAddEventOwnerUpdatedSuccess(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	err := AddEvent(OwnerUpdated{id: testAccountId, owner: updatedOwnerName}, &events)
	if err != nil {
		t.Error(err)
	}
}

func TestAddEventOwnerUpdatedFailureWithNotExistingAccount(t *testing.T) {
	var events []Event
	err := AddEvent(OwnerUpdated{id: testAccountId, owner: updatedOwnerName}, &events)
	if err == nil {
		t.Error("Expected ", AccountNotExistsError{})
	}
}

func TestAddEventDepositPerformedSuccess(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	err := AddEvent(DepositPerformed{id: testAccountId, amount: depositAmount}, &events)
	if err != nil {
		t.Error(err)
	}
}

func TestAddEventDepositPerformedFailureWithNotExistingAccount(t *testing.T) {
	var events []Event
	err := AddEvent(DepositPerformed{id: testAccountId, amount: depositAmount}, &events)
	if err == nil {
		t.Error("Expected ", AccountNotExistsError{})
	}
}

func TestAddEventDepositPerformedFailureWithNonPositiveAmount(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	err := AddEvent(DepositPerformed{id: testAccountId, amount: incorrectAmount}, &events)
	if err == nil {
		t.Error("Expected ", IncorrectAmountError{})
	}
}

func TestAddEventWithdrawalPerformedSuccess(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	events = append(events, DepositPerformed{id: testAccountId, amount: depositAmount})
	err := AddEvent(WithdrawalPerformed{id: testAccountId, amount: withdrawalAmount}, &events)
	if err != nil {
		t.Error(err)
	}
}

func TestAddEventWithdrawalPerformedFailureWithNotExistingAccount(t *testing.T) {
	var events []Event
	err := AddEvent(WithdrawalPerformed{id: testAccountId, amount: withdrawalAmount}, &events)
	if err == nil {
		t.Error("Expected ", AccountNotExistsError{})
	}
}

func TestAddEventWithdrawalPerformedFailureWithNonPositiveAmount(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	events = append(events, DepositPerformed{id: testAccountId, amount: depositAmount})
	err := AddEvent(WithdrawalPerformed{id: testAccountId, amount: incorrectAmount}, &events)
	if err == nil {
		t.Error("Expected ", IncorrectAmountError{})
	}
}

func TestAddEventWithdrawalPerformedFailureWithWithdrawalMoreThanBalance(t *testing.T) {
	var events []Event
	events = append(events, AccountCreated{id: testAccountId, owner: initialOwnerName})
	events = append(events, DepositPerformed{id: testAccountId, amount: depositAmount})
	err := AddEvent(WithdrawalPerformed{id: testAccountId, amount: tooBigWithdrawal}, &events)
	if err == nil {
		t.Error("Expected ", IncorrectAmountError{})
	}
}
