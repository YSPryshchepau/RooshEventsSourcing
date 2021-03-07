package account

func (error AccountNotExistsError) Error() string {
	return "Account no exists"
}

func (error IncorrectAmountError) Error() string {
	return "Incorrect amount"
}

func (error IncorrectIdError) Error() string {
	return "Incorrect account ID"
}

type AccountNotExistsError struct {

}

type IncorrectAmountError struct {

}

type IncorrectIdError struct {

}
