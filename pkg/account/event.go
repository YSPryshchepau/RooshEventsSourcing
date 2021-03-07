package account

type Event interface {
	isEvent()
	getId() int
}

func (event AccountCreated) isEvent() {

}

func (event AccountCreated) getId() int {
	return event.id
}

func (event OwnerUpdated) isEvent() {

}

func (event OwnerUpdated) getId() int {
	return event.id
}

func (event WithdrawalPerformed) isEvent() {

}

func (event WithdrawalPerformed) getId() int {
	return event.id
}

func (event DepositPerformed) isEvent() {

}

func (event DepositPerformed) getId() int {
	return event.id
}

type AccountCreated struct {
	id    int
	owner string
}

type OwnerUpdated struct {
	id    int
	owner string
}

type WithdrawalPerformed struct {
	id     int
	amount float64
}

type DepositPerformed struct {
	id     int
	amount float64
}
